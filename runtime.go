package manifold

import (
	"encoding/json"

	v8 "github.com/ry/v8worker2"
)

// New creates a new runtime with the given Invoker.
func New(i Invoker) (*Runtime, error) {
	v8.SetFlags([]string{"--harmony", "--harmony-dynamic-import"})
	worker := v8.New(func(args []byte) []byte {
		var req Request
		if err := json.Unmarshal(args, &req); err != nil {
			return Error(err)
		}
		return i.Invoke(req.Function, req.Args)
	})

	store := MemStore()
	resolver, err := newResolver(worker, store)
	if err != nil {
		worker.Dispose()
		return nil, err
	}

	return &Runtime{worker, store, resolver}, nil
}

// Runtime is the execution environment for v8 code. It resolves imports and manages communications between Go and v8.
type Runtime struct {
	worker   *v8.Worker
	store    Store
	resolver *moduleResolver
}

func (r *Runtime) LoadAsset(scriptName string) error {
	code, err := r.store.ReadModule(scriptName)
	if err != nil {
		return err
	}
	return r.worker.Load(scriptName, code)
}

func (r *Runtime) Load(scriptName string, code string) error {
	if err := r.store.WriteModule(scriptName, code); err != nil {
		return err
	}
	return r.worker.LoadModule(scriptName, code, r.resolver.Resolve)
}

func (r *Runtime) Terminate() {
	r.worker.TerminateExecution()
}

func (r *Runtime) Dispose() {
	r.worker.Dispose()
}

func (r *Runtime) Execute(main, code string) error {
	return r.worker.LoadModule(main, code, r.resolver.Resolve)
}
