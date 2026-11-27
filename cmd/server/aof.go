type Aof struct { 
	file *os.file
	rd *bufio.Reader
	mu   sync.Mutex
}