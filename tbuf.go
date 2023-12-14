package tbuf

import (
	"errors"
	"strings"
)

var (
	ErrInvalidBufferSize error = errors.New("invalid buffer size")
	ErrInvalidIndexRange error = errors.New("invalid index")
)

type Buffer struct {
	data  []string
	start int
	used  int
	size  int
	buf   *strings.Builder
}

func New(maxSize int) (*Buffer, error) {
	if maxSize >= 1 {
		return &Buffer{
			data:  make([]string, maxSize),
			start: 0,
			used:  0,
			size:  maxSize,
			buf:   &strings.Builder{},
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
	b.buf.Reset()

	for i := 0; i < b.used; i++ {
		if i != 0 {
			b.buf.WriteRune(' ')
		}
		b.buf.WriteString(b.data[(b.start+i)%b.size])
	}

	return b.buf.String()
}

func (b *Buffer) Join(sep string) string {
	b.buf.Reset()

	for i := 0; i < b.used; i++ {
		if i != 0 {
			b.buf.WriteString(sep)
		}
		b.buf.WriteString(b.data[(b.start+i)%b.size])
	}

	return b.buf.String()
}

func (b *Buffer) JoinFirst(limit int, sep string) string {
	if limit < 1 || b.used < 1 {
		return ""
	}

	if b.used < limit {
		limit = b.used
	}

	b.buf.Reset()

	for i := 0; i < limit; i++ {
		if i > 0 {
			b.buf.WriteString(sep)
		}
		b.buf.WriteString(b.data[(b.start+i)%b.size])
	}

	return b.buf.String()
}
