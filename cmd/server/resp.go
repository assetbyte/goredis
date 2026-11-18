// all the related to serializing and deserializing on the buffer.


package main
import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

const (
	STRING  = '+'
	ERROR   = '-'
	INTEGER = ':'
	BULK    = '$'
	ARRAY   = '*'
)


type Value struct {
	typ   string
	str   string
	num   int
	bulk  string
	array []Value
}


type Resp struct {
	reader *bufio.Reader
}

func NewResp(rd io.Reader) *Resp {
	return &Resp{reader: bufio.NewReader(rd)}
}

