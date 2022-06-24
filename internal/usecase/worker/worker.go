package worker

import (
	"context"
	"sync"

	"github.com/dimk00z/go-musthave-diploma/config"
	"github.com/dimk00z/go-musthave-diploma/internal/usecase"
	"github.com/dimk00z/go-musthave-diploma/pkg/logger"
	"golang.org/x/sync/errgroup"
)

type WorkersPool struct {
	workersNumber int
	inputCh       chan func(ctx context.Context) error
	done          chan struct{}
	l             *logger.Logger
}

var (
	wp   usecase.IWorkerPool
	once sync.Once
)

func NewWorkersPool(workersNumber int, poolLength int, l *logger.Logger) usecase.IWorkerPool {
	return &WorkersPool{
		workersNumber: workersNumber,
		inputCh:       make(chan func(ctx context.Context) error, poolLength),
		done:          make(chan struct{}),
		l:             l,
	}
}

func (wp *WorkersPool) Push(task func(ctx context.Context) error) {
	wp.inputCh <- task
}

func doTasksByWorkers(ctx context.Context,
	workerIndex int,
	taskCh chan func(ctx context.Context) error,
	l *logger.Logger) error {
	l.Debug("worker_%v started", workerIndex)
workerLoop:
	for {
		select {
		case <-ctx.Done():
			l.Debug("worker_%v got context.Done", workerIndex)
			break workerLoop
		case workerTask := <-taskCh:
			// l.Debug("worker_%v is busy", workerIndex)
			if err := workerTask(ctx); err != nil {
				l.Error("worker_%v got error:%s", workerIndex, err.Error())
				return err
			} else {
				// l.Debug("worker %v finished task correctly", workerIndex)
			}
		}
	}
	return nil
}

func (wp *WorkersPool) Run(ctx context.Context) {
	g, ctx := errgroup.WithContext(ctx)
	for workerIndex := 0; workerIndex < wp.workersNumber; workerIndex++ {
		workerIndex := workerIndex
		g.Go(func() error {
			return doTasksByWorkers(ctx, workerIndex, wp.inputCh, wp.l)
		})
	}
	if err := g.Wait(); err != nil {
		wp.l.Info(err.Error())
	}
	close(wp.inputCh)
}

func (wp *WorkersPool) Close() {
	close(wp.done)
}

func GetWorkersPool(wpConfig config.Workers, l *logger.Logger) usecase.IWorkerPool {
	once.Do(func() {
		wp = NewWorkersPool(wpConfig.WorkersNumber, wpConfig.PoolLength, l)
	})
	return wp
}
