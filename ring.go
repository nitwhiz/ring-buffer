package ring

import (
	"errors"
	"io"
)

var ErrOverflow = errors.New("overflow")
var ErrEOF = io.EOF

type Buffer[T any] struct {
	data []T
	size int

	readPtr  int
	writePtr int
}

func NewBuffer[T any](size int) *Buffer[T] {
	return &Buffer[T]{
		data: make([]T, size),
		size: size,

		readPtr:  0,
		writePtr: 0,
	}
}

func (b *Buffer[T]) read() (element T, err error) {
	if b.readPtr == b.writePtr {
		err = ErrEOF
		return
	}

	element = b.data[b.readPtr]
	b.readPtr = (b.readPtr + 1) % b.size

	return
}

func (b *Buffer[T]) write(element T) (err error) {
	nextPtr := (b.writePtr + 1) % b.size

	if nextPtr == b.readPtr {
		err = ErrOverflow
		return
	}

	b.data[b.writePtr] = element
	b.writePtr = nextPtr

	return
}

func (b *Buffer[T]) ReadOne() (element T, err error) {
	return b.read()
}

func (b *Buffer[T]) WriteOne(element T) (err error) {
	return b.write(element)
}

func (b *Buffer[T]) Write(p []T) (n int, err error) {
	for _, d := range p {
		if writeErr := b.write(d); writeErr != nil {
			err = writeErr
			return
		}

		n++
	}

	return n, nil
}

func (b *Buffer[T]) Read(p []T) (n int, err error) {
	for i := 0; i < len(p); i++ {
		element, readErr := b.read()

		if readErr != nil {
			err = readErr
			return
		}

		p[i] = element

		n++
	}

	return
}
