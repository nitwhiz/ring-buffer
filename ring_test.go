package ring

import (
	"errors"
	"testing"
)

func TestBuffer_Write(t *testing.T) {
	buf := NewBuffer[int](10)

	n, err := buf.Write([]int{1, 2, 3, 4, 5})

	if n != 5 {
		t.Error("n != 5")
	}

	if err != nil {
		t.Error("err != nil")
	}

	err = buf.WriteOne(6)

	if err != nil {
		t.Error("err != nil")
	}
}

func TestBuffer_Read(t *testing.T) {
	buf := NewBuffer[int](10)

	_, _ = buf.Write([]int{1, 2, 3, 4, 5, 6, 7, 8})

	one, err := buf.ReadOne()

	if one != 1 {
		t.Fatal("one != 1")
	}

	if err != nil {
		t.Fatal("err != nil")
	}

	two, err := buf.ReadOne()

	if two != 2 {
		t.Fatal("two != 2")
	}

	if err != nil {
		t.Fatal("err != nil")
	}

	three, err := buf.ReadOne()

	if three != 3 {
		t.Fatal("three != 3")
	}

	if err != nil {
		t.Fatal("err != nil")
	}

	data := make([]int, 2)

	n, err := buf.Read(data)

	if n != 2 {
		t.Fatal("n != 2")
	}

	if err != nil {
		t.Fatal("err != nil")
	}

	if data[0] != 4 {
		t.Fatal("data[0] != 4")
	}

	if data[1] != 5 {
		t.Fatal("data[0] != 4")
	}

	data = make([]int, 5)

	n, err = buf.Read(data)

	if n != 3 {
		t.Fatal("n != 3")
	}

	if !errors.Is(err, ErrEOF) {
		t.Fatal("err != ErrEOF")
	}

	if data[0] != 6 {
		t.Fatal("data[0] != 6")
	}

	if data[1] != 7 {
		t.Fatal("data[1] != 7")
	}

	if data[2] != 8 {
		t.Fatal("data[2] != 8")
	}
}

func TestBuffer_Overflow(t *testing.T) {
	buf := NewBuffer[int](3)

	data := []int{1, 2, 3, 4, 5, 6}

	n, err := buf.Write(data)

	if n != 2 {
		t.Fatal("n != 2")
	}

	if !errors.Is(err, ErrOverflow) {
		t.Fatal("err != ErrOverflow")
	}
}
