package main
import "fmt"

type Counter struct {
	value int // name <-> DT
}

func (c *Counter) increment(val int){
	c.value += val
}

func main() {
	counter := &Counter{value: 0}
	counter.increment(10)
    counter.increment(5)
	fmt.Println("Final Counter Value:", counter.value)
}