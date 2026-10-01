package kalkan

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

var (
	ErrPoolClosed     = errors.New("kalkan pool is closed")
	ErrAcquireTimeout = errors.New("timeout acquiring kalkan instance from pool")
)

// KalkanProvider определяет контракт получения экземпляра Kalkan для выполнения операции.
type KalkanProvider interface {
	Do(ctx context.Context, fn func(k Kalkan) error) error
	Close() error
}

// Pool управляет пулом экземпляров Kalkan для обеспечения параллелизма CGO.
type Pool struct {
	instances chan Kalkan
	mu        sync.RWMutex
	closed    bool
}

// NewPool создает и инициализирует пул экземпляров Kalkan.
func NewPool(size int, libPath string, options ...Option) (*Pool, error) {
	if size <= 0 {
		size = 4 // Значение по умолчанию
	}

	p := &Pool{
		instances: make(chan Kalkan, size),
	}

	for i := 0; i < size; i++ {
		mod, err := Open(libPath, options...)
		if err != nil {
			p.Close()
			return nil, fmt.Errorf("failed to initialize kalkan instance %d: %w", i, err)
		}
		if err := mod.Init(); err != nil {
			p.Close()
			return nil, fmt.Errorf("failed to init kalkan module %d: %w", i, err)
		}
		p.instances <- mod
	}

	return p, nil
}

// Do выполняет функцию fn с выделенным из пула экземпляром Kalkan.
func (p *Pool) Do(ctx context.Context, fn func(k Kalkan) error) error {
	p.mu.RLock()
	if p.closed {
		p.mu.RUnlock()
		return ErrPoolClosed
	}
	p.mu.RUnlock()

	select {
	case instance := <-p.instances:
		defer func() {
			p.instances <- instance
		}()
		return fn(instance)
	case <-ctx.Done():
		return ctx.Err()
	}
}

// Close грациозно закрывает все экземпляры Kalkan и освобождает ресурсы C-библиотеки.
func (p *Pool) Close() error {
	p.mu.Lock()
	if p.closed {
		p.mu.Unlock()
		return nil
	}
	p.closed = true
	close(p.instances)
	p.mu.Unlock()

	var errs []error
	for inst := range p.instances {
		inst.XMLFinalize()
		inst.Finalize()
		if err := inst.Close(); err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("errors closing kalkan pool: %v", errs)
	}
	return nil
}
