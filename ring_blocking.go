package ring

import "sync"

type BlockingBuffer[T any] struct {
	*Buffer[T]
	mu *sync.Mutex
}

func NewBlockingBuffer[T any](size int) *BlockingBuffer[T] {
	return &BlockingBuffer[T]{
		Buffer: NewBuffer[T](size),
		mu:     &sync.Mutex{},
	}
}

func (b *BlockingBuffer[T]) Write(p []T) (n int, err error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	return b.Buffer.Write(p)
}

func (b *BlockingBuffer[T]) Read(p []T) (n int, err error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	return b.Buffer.Read(p)
}
