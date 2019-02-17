package manifold

import "fmt"

//go:generate simpleassets -p builtin -o builtin/gen.go -t "builtin/" builtin/*js
import builtin "github.com/eliquious/manifold/builtin"

// Store stores module code in a virtual file system.
type Store interface {
	ReadModule(scriptName string) (string, error)
	WriteModule(scriptName string, code string) error
}

// MemStore creates a new in-memory module store
func MemStore() Store {
	return &memStore{}
}

type memStore struct {
}

func (ms *memStore) ReadModule(scriptName string) (string, error) {
	b, err := builtin.ReadAsset(scriptName)
	if err != nil {
		return "", fmt.Errorf("unable to load module: %s", scriptName)
	}
	return string(b), nil
}

func (ms *memStore) WriteModule(scriptName string, code string) error {
	builtin.WriteAsset(scriptName, []byte(code))
	return nil
}
