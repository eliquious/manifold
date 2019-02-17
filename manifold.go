package manifold

import (
	"encoding/json"
	"fmt"
)

// Request is a simple struct for passing data from v8 into Manifold.
type Request struct {
	Function string   `json:"fn"`
	Args     []string `json:"args"`
}

// Response is a simple struct for returning data to v8 after execution.
type Response struct {
	Data  interface{} `json:"data"`
	Error string      `json:"err"`
}

// Invoker handles the invocation of a Golang function by v8.
type Invoker interface {
	Invoke(fn string, args []string) []byte
}

// EncodeResponse marshals a response into JSON.
func EncodeResponse(data interface{}) []byte {
	b, err := json.Marshal(Response{Data: data})
	if err != nil {
		return Error(err)
	}
	return b
}

// Error marshals an error into JSON.
func Error(err error) []byte {
	return []byte(fmt.Sprintf(`{"err": "%s"}`, err.Error()))
}
