package keeper

import (
	"context"

	wasm "github.com/CosmWasm/go-cosmwasm"
)

func (k Keeper) getWasmer(ctx context.Context) (*wasm.Wasmer, error) {
	err := k.wasmerReadSemaphore.Acquire(ctx, 1)
	if err != nil {
		return nil, err
	}

	k.wasmerReadMutex.Lock()
	wasmVM := k.wasmerReadPool[0]
	k.wasmerReadPool = k.wasmerReadPool[1:]
	k.wasmerReadMutex.Unlock()

	return wasmVM, nil
}

func (k Keeper) putWasmer(wasmVM *wasm.Wasmer) {
	k.wasmerReadMutex.Lock()
	k.wasmerReadPool = append(k.wasmerReadPool, wasmVM)
	k.wasmerReadMutex.Unlock()

	k.wasmerReadSemaphore.Release(1)
}
