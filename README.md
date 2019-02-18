# manifold

`manifold` is a toolkit for building v8/go APIs. It uses a fork of [v8worker2](https://github.com/ry/v8worker2) as the communication layer. This is a simple abstraction for message passing between Go and JS running in v8.

This is not meant to be used as complete bindings for v8, but rather a barebones library for communication and extensibility between the languages.

# Usage

Manifold consists of very few interfaces which allow for the loading of importable modules, executing scripts and handling message passing. THe message passing is handled by an `Invoker` which is invoked every time the JS sends a message. The message handler uses the simplistic signature of `Invoke(fn string, args []string) []byte`. The message is then handled and the raw bytes are sent back to the JS as an ArrayBuffer which is parsed as JSON upon receipt. 

## Execution Environment

There are only a few methods exposed to the JS runtime. `v8worker2` provides for global `print`, `send` and `recv` methods. However, `send` is wrapped to provide for the builtin formatting of the messages. Additionally, due to the ability to load modules base64 support is also provided. More builtin modules can easily be added. Builtin modules are loaded via the `simpleassets` command line tool.



# Complete example

```go
//go:generate simpleassets -p lib -o lib/gen.go -t "./" ./*js
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/eliquious/manifold"
	"github.com/eliquious/manifold/examples/incr/lib"
)

func main() {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds)

	// New manifold
	m, err := manifold.New(&Invoker{})
	if err != nil {
		panic(err)
	}
	defer m.Dispose()

	LoadAssets(logger, m)
	ExecuteScript(logger, m)
}

// ExecuteScript runs the script.
func ExecuteScript(logger *log.Logger, r *manifold.Runtime) {
	src, err := lib.ReadAsset("index.js")
	if err != nil {
		log.Fatal(err)
	}
	err = r.Execute("index.js", string(src))
	if err != nil {
		log.Fatal(err)
	}
}

// LoadAssets loads requried assets.
func LoadAssets(logger *log.Logger, r *manifold.Runtime) {
	for _, asset := range lib.ListAssetNames() {
		if asset != "index.js" {
			src, err := lib.ReadAsset(asset)
			if err != nil {
				logger.Fatal(err)
			}
			if err := r.Load(asset, string(src)); err != nil {
				logger.Fatal(err)
			}
		}
	}
}

// Invoker managed the API and state
type Invoker struct {
	num int
}

// Invoke calls a function on the API
func (i *Invoker) Invoke(fn string, args []string) []byte {
	if fn == "incr" {
		return i.Incr()
	} else if fn == "decr" {
		return i.Decr()
	} else if fn == "get" {
		return manifold.EncodeResponse(i.num)
	}
	return manifold.Error(fmt.Errorf("undefined function: %s", fn))
}

// Incr increments the stored value
func (i *Invoker) Incr() []byte {
	i.num++
	return manifold.EncodeResponse(i.num)
}

// Decr decrements the stored value
func (i *Invoker) Decr() []byte {
	if i.num > 0 {
		i.num--
		return manifold.EncodeResponse(i.num)
	}
	return manifold.Error(fmt.Errorf("cannot decrement below 0"))
}
```

## Future Work

The API needs to be updated to provide the ability to limit the size of the v8 runtime. This is currently possible via v8 flags and a heap listener it just needs to be supported on the API.

More builtin libraries could be added or optionally included in the runtime.
