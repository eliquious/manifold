package manifold

import v8 "github.com/eliquious/v8worker2"
import builtin "github.com/eliquious/manifold/builtin"

// newResolver creates a new resolver for v8 imports.
func newResolver(w *v8.Worker, s Store) (*moduleResolver, error) {
	r := &moduleResolver{w, s}

	for _, assetName := range builtin.ListAssetNames() {
		if assetName == "index.js" {
			continue
		}

		code, err := s.ReadModule(assetName)
		if err != nil {
			return nil, err
		}

		// Include required built in functions
		if err := w.Load(assetName, code); err != nil {
			return nil, err
		}
	}

	// // Read core.js from VFS
	// b64, err := s.ReadModule("base64.js")
	// if err != nil {
	// 	return nil, err
	// }

	// // Include required built in functions
	// if err := w.Load("base64.js", b64); err != nil {
	// 	return nil, err
	// }

	// // Read core.js from VFS
	// core, err := s.ReadModule("core.js")
	// if err != nil {
	// 	return nil, err
	// }

	// // Include required built in functions
	// if err := w.Load("core.js", core); err != nil {
	// 	return nil, err
	// }
	return r, nil
}

type moduleResolver struct {
	worker *v8.Worker
	store  Store
}

func (r *moduleResolver) Resolve(specifier, referrer string) int {

	// Read from store
	code, err := r.store.ReadModule(specifier)
	if err != nil {
		return 1
	}

	// Include in runtime
	// if err := r.worker.Load(specifier, code); err != nil {
	// 	return 1
	// }
	if err := r.worker.LoadModule(specifier, code, r.Resolve); err != nil {
		return 1
	}
	return 0
}
