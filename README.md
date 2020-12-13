# tbuf

Simple text buffer with fixed max size optimized for deletion from begin/end and random elements access

## Sumary

* Go version >= 1.10
* No external dependencies

## Install

```plaintext
go get github.com/wmentor/tbuf
```

## Usage

```golang
package main

import (
  "fmt"
  "strconv"

  "github.com/wmentor/tbuf"
)

func main() {
  buf, e := tbuf.New(-1)
  fmt.Println(buf ,e) // nil and error

  fmt.Println("create buffer:")

  buf, e = tbuf.New(5) // *tbuf.Buffer and nil
  for i := 0 ; i < 5 ; i++ {
    buf.Push(strconv.Itoa(i))
  }

  fmt.Println(buf.IsFull()) // true

  for i := 0 ; i < buf.Len() ; i++ {
    val, err := buf.Get(i)
    fmt.Println(val, err)
  }

  buf.Push("11")
  buf.Push("12")

  fmt.Println("buffer after adding two items:")

  for i := 0 ; i < buf.Len() ; i++ {
    val, err := buf.Get(i)
    fmt.Println(val, err)
  }

  buf.ShiftN(1)
  buf.PopN(2)

  fmt.Println("remove 1 from start and 2 from end")

  for i := 0 ; i < buf.Len() ; i++ {
    val, err := buf.Get(i)
    fmt.Println(val, err)
  }

  buf.Reset() // remove all items
  fmt.Println(buf.IsEmpty()) // true
}
```

Output:

```plaintext
<nil> invalid buffer size
create buffer:
true
0 <nil>
1 <nil>
2 <nil>
3 <nil>
4 <nil>
buffer after adding two items:
2 <nil>
3 <nil>
4 <nil>
11 <nil>
12 <nil>
remove 1 from start and 2 from end
3 <nil>
4 <nil>
true
```
