package apps

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/dig"
)

// Start simple start function for an app using DI container
type Start func(ctx context.Context, dic *dig.Container)

// Container App container including DI, a list of app Start
// functions that will be executed on Run.
type Container struct {
	DIC       *dig.Container
	StartList []Start
	Ctx       context.Context
	sig       chan os.Signal
}

// NewContainer constructor for the app container
func NewContainer(ctx context.Context, dic *dig.Container, startList []Start) *Container {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	signal.Notify(sig, syscall.SIGTERM)

	return &Container{
		DIC:       dic,
		StartList: startList,
		Ctx:       ctx,
		sig:       sig,
	}
}

// Run simple func to start all the applications on the Start list and hold the threads
func (ref *Container) Run() {
	if len(ref.StartList) == 0 {
		return
	}
	for _, start := range ref.StartList {
		go start(ref.Ctx, ref.DIC)
	}

	<-ref.sig
	fmt.Println("Graceful shutdown, sigterm received.")
	_ = ref.Ctx.Done()
}

// Shutdown gracefully stops the Container context
func (ref *Container) Shutdown() {
	ref.sig <- os.Interrupt
}
