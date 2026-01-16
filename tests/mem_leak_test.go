package tests

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"testing"
	"time"
)

var (
	pauseAfterClose = 1 * time.Minute
	iterations      = 500
)

func TestMemLeak(t *testing.T) {
	data := make([]byte, 262144)
	_, err = rand.Read(data)
	if err != nil {
		t.Fatalf("rand error: %v", err)
	}

	for i := 0; i < iterations; i++ {
		_, err = mod.SignData("", string(data), "", flags)
		if err != nil {
			t.Fatalf("sign error: %v", err)
		}

		log.Println(memStatString(i))
	}

	log.Println("before close:", memStatString(iterations))

	if err = mod.Close(); err != nil {
		t.Fatalf("mod close error: %v\n", err)
	}

	log.Println("after close:", memStatString(iterations))

	time.Sleep(pauseAfterClose)

	log.Println("after sleep:", memStatString(iterations))

}
func memStatString(iter int) string {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	rss := readVmRSSBytes()

	return fmt.Sprintf("iter=%d alloc=%.2fMB rss=%.2fMB",
		iter,
		float64(m.Alloc)/1024.0/1024.0,
		float64(rss)/1024.0/1024.0,
	)
}

func readVmRSSBytes() uint64 {
	f, err := os.Open("/proc/self/status")
	if err != nil {
		log.Printf("open status error: %v", err)
	}

	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()
		log.Println("line: ", sc.Text())
		if strings.HasPrefix(line, "VmRSS:") {
			fields := strings.Fields(line)
			if len(fields) >= 2 {
				kb, _ := strconv.ParseUint(fields[1], 10, 64)
				return kb * 1024
			}
		}
	}

	return 0
}
