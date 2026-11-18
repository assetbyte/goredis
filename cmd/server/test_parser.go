package main 
import (
    "bufio"
    "fmt"
    "strings"
    "os"
    "strconv"
)


func main() {
    input := "$5\r\nAsset\r\n"
	reader := bufio.NewReader(strings.NewReader(input)) // IO reader

	b, _ := reader.ReadByte()
	if b != '$' {
      fmt.Println("Invalid type, expecting bulk strings only")
      os.Exit(1)
    }

	size, _ :=  reader.ReadByte()

	strSize, _ := strconv.ParseInt(string(size), 10, 64) // -> int64

	reader.ReadByte()
    reader.ReadByte()

	name := make([]byte, strSize)
	reader.Read(name)

	fmt.Println(string(name))

}