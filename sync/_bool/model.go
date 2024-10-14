package _bool

import (
	"context"
	"sync"

	"github.com/kyaxcorp/go-helper/sync/_map_string_interface"
	"github.com/kyaxcorp/go-helper/sync/waiter"
)

type Bool struct {
	lock  sync.RWMutex
	value bool
	// TODO: we can add here also OnSet, OnGet events

	// Parent Context...
	ctx context.Context

	// it's needed for WaitFor True & False
	eventTrigger chan Bool

	// Here we store the events
	onChange      *_map_string_interface.MapStringInterface
	onChangeAsync *_map_string_interface.MapStringInterface

	onTrue      *_map_string_interface.MapStringInterface
	onTrueAsync *_map_string_interface.MapStringInterface

	onFalse      *_map_string_interface.MapStringInterface
	onFalseAsync *_map_string_interface.MapStringInterface

	onTrueWaiter  *waiter.Waiter
	onFalseWaiter *waiter.Waiter
}
