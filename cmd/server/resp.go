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

func (r *Resp) readLine() (line []byte, n int, err error){
// line - read string as slice of bytes 
// n - amount of read bytes
// err - error type 
	for {
		b, err := r.reader.ReadByte()
		if err != nil {
			return nil, 0, err
		}

		n +=1 
		line = append(line, b)
		if len(line) >= 2 && line[len(line) - 2 ] == '\r'{break}
	}
	return line[:len(line)-2], n, nil
}

