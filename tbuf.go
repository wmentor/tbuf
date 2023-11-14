package tbuf

import (
	"errors"
	"strings"
	"sync"
)

var (
	ErrInvalidBufferSize error = errors.New("invalid buffer size")
	ErrInvalidIndexRange error = errors.New("invalid index")
)

var (
	pool = sync.Pool{
		New: func() any {
			return &strings.Builder{}
		},
	}
)

type Buffer struct {
	data  []string
	start int
	used  int
	size  int
}

func New(maxSize int) (*Buffer, error) {
	if maxSize >= 1 {
		return &Buffer{
			data:  make([]string, maxSize),
			start: 0,
			used:  0,
			size:  maxSize,
		}, nil
	}

	return nil, ErrInvalidBufferSize
}

func (b *Buffer) Len() int {
	return b.used
}

func (b *Buffer) IsEmpty() bool {
	return b.used == 0
}

func (b *Buffer) IsFull() bool {
	return b.used == b.size
}

func (b *Buffer) Push(val string) {
	if b.used < b.size {
		b.data[(b.start+b.used)%b.size] = val
		b.used++
	} else {
		b.data[b.start] = val
		b.start = (b.start + 1) % b.size
	}
}

func (b *Buffer) Get(idx int) (string, error) {
	if idx >= 0 && idx < b.used {
		return b.data[(b.start+idx)%b.size], nil
	}

	return "", ErrInvalidIndexRange
}

func (b *Buffer) ShiftN(num int) {
	if num > 0 {
		if num >= b.used {
			b.start = 0
			b.used = 0
		} else {
			b.start = (b.start + num) % b.size
			b.used -= num
		}
	}
}

func (b *Buffer) Shift() {
	b.ShiftN(1)
}

func (b *Buffer) PopN(num int) {
	if num > 0 {
		if num >= b.used {
			b.start = 0
			b.used = 0
		} else {
			b.used -= num
		}
	}
}

func (b *Buffer) Reset() {
	b.used = 0
	b.start = 0
}

func (b *Buffer) Pop() {
	b.PopN(1)
}

func (b *Buffer) String() string {
	buf := pool.Get().(*strings.Builder)
	defer pool.Put(buf)

	buf.Reset()

	for i := 0; i < b.Len(); i++ {
		if i != 0 {
			buf.WriteRune(' ')
		}
		buf.WriteString(b.data[(b.start+i)%b.size])
	}

	return buf.String()
}
