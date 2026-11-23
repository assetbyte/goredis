package main

import "strconv"

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

func (v Value) marshalString() []byte {
	var bytes []byte
	// STRING - const
	bytes = append(bytes, STRING)   //bytes: ['+']
	bytes = append(bytes, v.str...) // ellipsis
	bytes = append(bytes, '\r', '\n')
	return bytes
}

// bulk string format:
// $<length(in bytes)>\r\n<content>\r\n
func (v Value) marshalBulk() []byte {
	var bytes []byte
	// BULK - const
	bytes = append(bytes, BULK)
	bytes = append(bytes, strconv.Itoa(len(v.bulk))...)
	bytes = append(bytes, '\r', '\n')
	bytes = append(bytes, v.bulk...)
	bytes = append(bytes, '\r', '\n')
}

// array format:
// *<elements_count>\r\n<element_1_marshaled><element_2_marshaled>...
func (v Value) marshalArray() []byte {
	var bytes []byte
	// ARRAY - const
	_len := len(v.array)

	bytes = append(bytes, ARRAY)
	bytes = append(bytes, []byte(strconv.Itoa(_len))...)
	bytes = append(bytes, '\r', '\n')

	for i := 0; i < _len; i++ {
		bytes = append(bytes, v.array[i].Marshal()...)
	}
	return bytes
}