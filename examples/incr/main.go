//go:generate simpleassets -p lib -o lib/gen.go -t "./" ./*js
package main

import (
	"fmt"

	"github.com/eliquious/manifold"
	"github.com/eliquious/manifold/examples/incr/lib"
)

func main() {
	// New manifold
	m, err := manifold.New(&Invoker{})
	if err != nil {
		panic(err)
	}

	//
	for _, asset := range lib.ListAssetNames() {
		if asset == "index.js" {
			continue
		}

		src, err := lib.ReadAsset(asset)
		if err != nil {
			panic(err)
		}

		if err := m.Load(asset, string(src)); err != nil {
			panic(err)
		}
	}

	// Read asset
	src, err := lib.ReadAsset("index.js")
	if err != nil {
		panic(err)
	}

	// Execute
	if err := m.Execute("index.js", string(src)); err != nil {
		panic(err)
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
