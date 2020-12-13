package tbuf

import (
	"testing"
)

func TestBuf(t *testing.T) {

	tNewF := func(size int) {
		b, e := New(size)
		if b != nil || e != ErrInvalidBufferSize {
			t.Fatalf("Success New with size=%d", size)
		}
	}

	tNewF(-1)
	tNewF(0)
	tNewF(-412342)

	b, e := New(5)
	if b == nil || e != nil || !b.IsEmpty() || b.IsFull() {
		t.Fatalf("New failed with size=%d", 5)
	}

	tPush := func(val string, wait []string) {
		b.Push(val)
		if b.Len() != len(wait) || b.IsEmpty() || (b.Len() == 5 && !b.IsFull()) {
			t.Fatalf("Push failed for value=%s", val)
		}

		for i, v := range wait {
			cur, err := b.Get(i)
			if err != nil || cur != v {
				t.Fatal("Invalid buffer value")
			}
		}
	}

	tPush("1", []string{"1"})
	tPush("2", []string{"1", "2"})
	tPush("3", []string{"1", "2", "3"})
	tPush("4", []string{"1", "2", "3", "4"})
	tPush("5", []string{"1", "2", "3", "4", "5"})
	tPush("6", []string{"2", "3", "4", "5", "6"})
	tPush("7", []string{"3", "4", "5", "6", "7"})
	tPush("8", []string{"4", "5", "6", "7", "8"})

	v, e := b.Get(7)
	if v != "" || e != ErrInvalidIndexRange {
		t.Fatal("Get out of range failed")
	}

	tShift := func(size int, wait []string) {
		b.ShiftN(size)
		if b.Len() != len(wait) {
			t.Fatal("Shift failed")
		}

		for i, v := range wait {
			cur, err := b.Get(i)
			if err != nil || cur != v {
				t.Fatal("Invalid buffer value")
			}
		}
	}

	tPop := func(size int, wait []string) {
		b.PopN(size)
		if b.Len() != len(wait) {
			t.Fatal("Pop failed")
		}

		for i, v := range wait {
			cur, err := b.Get(i)
			if err != nil || cur != v {
				t.Fatal("Invalid buffer value")
			}
		}
	}

	tShift(0, []string{"4", "5", "6", "7", "8"})
	b.Shift()
	tShift(1, []string{"6", "7", "8"})
	tShift(2, []string{"8"})
	tShift(10, nil)
	tPush("1", []string{"1"})
	tPush("2", []string{"1", "2"})
	tPush("3", []string{"1", "2", "3"})
	tPush("4", []string{"1", "2", "3", "4"})
	tPush("5", []string{"1", "2", "3", "4", "5"})
	tPush("6", []string{"2", "3", "4", "5", "6"})
	b.Pop()
	tPop(3, []string{"2"})
	tPop(2, nil)

}
