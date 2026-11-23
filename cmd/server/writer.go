package main 

func (v Value) Marshal() []byte {
	switch v.typ {
	case "array": 
		return v.marshalArray()
	case "bulk":
		return v.marshalBulk()
	case "string":
		return v.marshalString()
	case "null":
		return v.marshalNull()
	case "error":
		return v.marshalError()
	default:
		return []byte{}
	}
}

// STRING - const 
func (v Value) marshalString() []byte {
	var bytes []byte
	bytes = append(bytes, STRING) //bytes: ['+']
	bytes = append(bytes, v.str...) // ellipsis
	bytes = append(bytes, '\r', '\n')
	return bytes
}

