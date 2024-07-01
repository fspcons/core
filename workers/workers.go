package workers

import (
	"context"
	"fmt"

	"go.uber.org/zap"
)

// Result godoc
type Result struct {
	Val any
	Err error
}

// Job godoc
type Job struct {
	ID   string
	Task func(args any) Result
	Args any
}

// Pool godoc
type Pool interface {
	Start()
	Stop()
	AddJob(jobs ...Job)
	Results() <-chan Result
}

type pool struct {
	workerCount int
	jobsCh      chan Job
	resultsCh   chan Result
	ctx         context.Context
	cancel      context.CancelFunc
	logger      *zap.Logger
}

// NewPool godoc
func NewPool(ctx context.Context, workerCount int, logger *zap.Logger) Pool {
	ctx, cancel := context.WithCancel(ctx)

	return &pool{
		workerCount: workerCount,
		jobsCh:      make(chan Job, workerCount),
		resultsCh:   make(chan Result, workerCount),
		ctx:         ctx,
		cancel:      cancel,
		logger:      logger,
	}
}

// Start godoc
func (ref pool) Start() {
	for i := 0; i < ref.workerCount; i++ {
		ref.logger.Info(fmt.Sprintf("Starting worker %d", i+1))
		go ref.run()
	}
}

// Stop godoc
func (ref pool) Stop() {
	ref.cancel()
}

// Results godoc
func (ref pool) Results() <-chan Result {
	return ref.resultsCh
}

// AddJob godoc
func (ref pool) AddJob(jobs ...Job) {
	for _, j := range jobs {
		ref.jobsCh <- j
	}
}

func (ref pool) run() {
	for {
		select {
		case job := <-ref.jobsCh:
			ref.logger.With(zap.String("jobId", job.ID)).Info("running job")

			ref.resultsCh <- job.Task(job.Args) //multiplexing results into the results channel

		case <-ref.ctx.Done():
			ref.logger.Info("Worker graceful shutdown")
			ref.resultsCh <- Result{
				Err: ref.ctx.Err(),
			}
			return
		}
	}
}
