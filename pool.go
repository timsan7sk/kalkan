package kalkan

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

var (
	ErrPoolClosed     = errors.New("kalkan pool is closed")
	ErrAcquireTimeout = errors.New("timeout acquiring instance from pool")
)

// Pool — потокобезопасный обобщенный пул ресурсов на основе каналов.
// Использование дженериков позволяет переиспользовать пул для любых CGO сущностей.
type Pool[T any] struct {
	instances chan T
	mu        sync.RWMutex
	closed    bool
	closeFunc func(T) error // Вариативный замыкатель для финализации
}

// NewPool инициализирует пул ресурсов.
// Принимает фабричную функцию и опциональные вариативные конфигураторы (...).
func NewPool[T any](size int, factory func() (T, error), closeFunc func(T) error, opts ...func(*Pool[T])) (*Pool[T], error) {
	if size <= 0 {
		size = 1
	}

	p := &Pool[T]{
		instances: make(chan T, size),
		closeFunc: closeFunc,
	}

	for _, opt := range opts {
		opt(p)
	}

	for i := 0; i < size; i++ {
		inst, err := factory()
		if err != nil {
			p.Close()
			return nil, fmt.Errorf("pool initialization failed at instance %d: %w", i, err)
		}
		p.instances <- inst
	}

	return p, nil
}

// Do исполняет функцию-замыкание, конкурентно извлекая экземпляр из пула.
func (p *Pool[T]) Do(ctx context.Context, fn func(inst T) error) error {
	p.mu.RLock()
	if p.closed {
		p.mu.RUnlock()
		return ErrPoolClosed
	}
	p.mu.RUnlock()

	select {
	case inst := <-p.instances:
		// Гарантированный возврат инстанса обратно в канал (Lease Release)
		defer func() { p.instances <- inst }()
		return fn(inst)
	case <-ctx.Done():
		return ctx.Err()
	}
}

// Close безопасно завершает работу пула, опустошая канал и вызывая closeFunc.
func (p *Pool[T]) Close() error {
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
		if p.closeFunc != nil {
			if err := p.closeFunc(inst); err != nil {
				errs = append(errs, err)
			}
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("errors closing pool: %v", errs)
	}
	return nil
}
