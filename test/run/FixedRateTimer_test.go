package test_run

import (
	"context"
	"sync/atomic"
	"testing"
	"time"

	"github.com/pip-services3-gox/pip-services3-commons-gox/run"
	"github.com/stretchr/testify/assert"
)

func TestTimerWithCallback(t *testing.T) {
	var counter int32

	timer := run.NewFixedRateTimerFromCallback(
		func(ctx context.Context) {
			atomic.AddInt32(&counter, 1)
		},
		100, 0, 5,
	)

	ctx := context.Background()
	timer.Start(ctx)
	time.Sleep(time.Millisecond * 500)
	timer.Stop(ctx)

	assert.True(t, atomic.LoadInt32(&counter) > 3)
}

func TestTimerWithCancelCallback(t *testing.T) {
	var counter, counterCanceled int32

	timer := run.NewFixedRateTimerFromCallback(
		func(ctx context.Context) {
			atomic.AddInt32(&counter, 1)
			select {
			case <-ctx.Done():
				atomic.AddInt32(&counterCanceled, 1)
				break
			}
		},
		100, 0, 5,
	)

	ctx, cancel := context.WithCancel(context.Background())
	timer.Start(ctx)
	time.Sleep(time.Millisecond * 500)
	cancel()
	timer.Stop(ctx)
	time.Sleep(time.Millisecond * 100)

	assert.True(t, atomic.LoadInt32(&counter) > 3)
	assert.True(t, atomic.LoadInt32(&counterCanceled) > 3)
}
