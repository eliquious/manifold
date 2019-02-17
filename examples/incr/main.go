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
	// defer m.Dispose()

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
