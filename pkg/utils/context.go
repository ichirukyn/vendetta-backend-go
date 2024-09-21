package utils

import (
	"context"
	"sync"
	"time"
)

type CancelFunc context.CancelFunc
type DoneFunc func()
type TimeoutFunc func()

// State - состояние жизненного цикла контекста.
type State int

const (
	Created State = iota
	Running
	Canceled
	Deadlined
	Finished
)

// Context - наша обёртка для контекста.
type Context struct {
	ctx    context.Context
	cancel context.CancelFunc
	mu     sync.Mutex

	onDone    []func()
	onCancel  []func()
	onTimeout []func()

	state State
}

// startMonitoring переводит состояние в "В работе" и следит за завершением контекста.
func (cw *Context) startMonitoring() {
	cw.setState(Running)
	<-cw.ctx.Done()

	if err := cw.ctx.Err(); err != nil {
		switch err {
		case context.Canceled:
			cw.setState(Canceled)
			for _, f := range cw.onCancel {
				f()
			}
		case context.DeadlineExceeded:
			cw.setState(Deadlined)
			for _, f := range cw.onTimeout {
				f()
			}
		default:
			cw.setState(Finished)
		}
	}
	for _, f := range cw.onDone {
		f()
	}
}

// OnDone добавляет коллбек, вызываемый при завершении контекста.
func (cw *Context) OnDone(f func()) {
	cw.onDone = append(cw.onDone, f)
}

// OnCancel добавляет коллбек, вызываемый при отмене контекста.
func (cw *Context) OnCancel(f func()) {
	cw.onCancel = append(cw.onCancel, f)
}

// OnTimeout добавляет коллбек, вызываемый при истечении времени контекста.
func (cw *Context) OnTimeout(f func()) {
	cw.onTimeout = append(cw.onTimeout, f)
}

// Done возвращает канал завершения контекста.
func (cw *Context) Done() <-chan struct{} {
	return cw.ctx.Done()
}

// IsDone возвращает статус, завершен ли контекст.
func (cw *Context) IsDone() bool {
	return cw.state != Created && cw.state != Running
}

// Cancel вызывает отмену контекста.
func (cw *Context) Cancel() {
	cw.cancel()
}

// GetCancelFunc возвращает функцию отмены.
func (cw *Context) GetCancelFunc() context.CancelFunc {
	return cw.cancel
}

// Err возвращает ошибку контекста.
func (cw *Context) Err() error {
	return cw.ctx.Err()
}

// Value возвращает значение, сохранённое в контексте.
func (cw *Context) Value(key interface{}) interface{} {
	return cw.ctx.Value(key)
}

// State возвращает текущее состояние жизненного цикла контекста.
func (cw *Context) State() State {
	cw.mu.Lock()
	defer cw.mu.Unlock()
	return cw.state
}

// setState устанавливает новое состояние жизненного цикла.
func (cw *Context) setState(state State) {
	cw.mu.Lock()
	defer cw.mu.Unlock()
	cw.state = state
}

// GetContext возвращает оригинальный контекст.
func (cw *Context) GetContext() context.Context {
	return cw.ctx
}

// NewContextWithCancel создает новый Context с возможностью отмены.
func NewContextWithCancel(parent context.Context) *Context {
	ctx, cancel := context.WithCancel(parent)
	cw := &Context{
		ctx:    ctx,
		cancel: cancel,
		state:  Created,
	}
	go cw.startMonitoring()
	return cw
}

// NewContextWithTimeout создает новый Context с таймаутом.
func NewContextWithTimeout(parent context.Context, timeout time.Duration) *Context {
	ctx, cancel := context.WithTimeout(parent, timeout)
	cw := &Context{
		ctx:    ctx,
		cancel: cancel,
		state:  Created,
	}
	go cw.startMonitoring()
	return cw
}
