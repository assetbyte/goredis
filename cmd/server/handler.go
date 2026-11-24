package main

import (
	"fmt"
	"sync"
)

var Handlers = map[string]func([]Value) Value{
	"PING": ping,
	"SET": set,
	"GET": get,
}

func ping(args []Value) Value {
	if len(args) == 0 {
		return Value{typ: "string", str: "PONG"}
	}
	return Value{typ: "string", str: args[0].bulk}
}

var SETs = map[string]string{}
var SETsMU = sync.RWMutex{}

func set(args []Value) Value {
	if len(args) != 2 {
		return Value{typ: "error", str: "ERR wrong number of arguments for 'set' command"} // SET command takes exactly 2 argumants 
	}

	key := args[0].bulk
	value := args[1].bulk

	SETsMU.Lock()
	SETs[key] = value
	SETsMU.Unlock()
	return Value{typ: "string", str: "OK"}
}

func get (args []Value) Value {
	if len(args) != 1 {
		return Value{typ: "error", str: "ERR wrong number of arguments for 'get' command"}
	}
	key := args[0].bulk
	SETsMU.RLock()
	value, ok := SETs[key]
	SETsMU.RUnlock()

	if !ok {
		return Value{typ: "null"}
	}
	return Value{typ: "bulk", bulk: value}

}