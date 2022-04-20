package run

import (
	"sync"
	"time"
)

// FixedRateTimer timer that is triggered in equal time intervals.
// It has summetric cross-language implementation and is often used by Pip.Services toolkit
// to perform periodic processing and cleanup in microservices.
//	see INotifiable
//	Example:
//		type MyComponent {
//			timer FixedRateTimer
//		}
//		...
//		func (mc* MyComponent) open(correlationId string) {
//			...
//			mc.timer = NewFixedRateTimerFromCallback(() => { this.cleanup }, 60000, 0);
//			mc.timer.start();
//			...
//		}
//		func (mc* MyComponent) open(correlationId: string){
//			...
//			mc.timer.stop();
//			...
//		}
type FixedRateTimer struct {
	task     INotifiable
	callback func()
	delay    int
	interval int
	ticker   *time.Ticker
	mtx      sync.Mutex
}

// NewFixedRateTimer creates new instance of the timer and sets its values.
//	Returns: *FixedRateTimer
func NewFixedRateTimer() *FixedRateTimer {
	return &FixedRateTimer{}
}

// NewFixedRateTimerFromCallback creates new instance of the timer and sets its values.
//	Parameters:
//		- callback func() callback function to call when timer is triggered.
//		- interval int an interval to trigger timer in milliseconds.
//		- delay int a delay before the first triggering in milliseconds.
//	Returns: *FixedRateTimer
func NewFixedRateTimerFromCallback(callback func(), interval int, delay int) *FixedRateTimer {
	return &FixedRateTimer{
		callback: callback,
		interval: interval,
		delay:    delay,
	}
}

// NewFixedRateTimerFromTask creates new instance of the timer and sets its values.
//	Parameters:
//		- callback INotifiable Notifiable object to call when timer is triggered.
//		- interval int an interval to trigger timer in milliseconds.
//		- delay int a delay before the first triggering in milliseconds.
//	Returns: *FixedRateTimer
func NewFixedRateTimerFromTask(task INotifiable, interval int, delay int) *FixedRateTimer {
	c := &FixedRateTimer{
		interval: interval,
		delay:    delay,
	}
	c.SetTask(task)
	return c
}

// Task gets the INotifiable object that receives notifications from this timer.
//	Returns: INotifiable the INotifiable object or null if it is not set.
func (c *FixedRateTimer) Task() INotifiable {
	return c.task
}

// SetTask sets a new INotifiable object to receive notifications from this timer.
//	Parameters: value INotifiable a INotifiable object to be triggered.
func (c *FixedRateTimer) SetTask(value INotifiable) {
	c.task = value
	c.callback = func() {
		c.task.Notify("timer", NewEmptyParameters())
	}
}

// Callback gets the callback function that is called when this timer is triggered.
//	Returns: function the callback function or null if it is not set.
func (c *FixedRateTimer) Callback() func() {
	return c.callback
}

// SetCallback sets the callback function that is called when this timer is triggered.
//	Parameters: value func() the callback function to be called.
func (c *FixedRateTimer) SetCallback(value func()) {
	c.callback = value
	c.task = nil
}

// Delay gets initial delay before the timer is triggered for the first time.
//	Returns: number the delay in milliseconds.
func (c *FixedRateTimer) Delay() int {
	return c.delay
}

// SetDelay sets initial delay before the timer is triggered for the first time.
//	Parameters: value int a delay in milliseconds.
func (c *FixedRateTimer) SetDelay(value int) {
	c.delay = value
}

// Interval gets periodic timer triggering interval.
//	Returns: number the interval in milliseconds
func (c *FixedRateTimer) Interval() int {
	return c.interval
}

// SetInterval sets periodic timer triggering interval.
//	Parameters: value int an interval in milliseconds.
func (c *FixedRateTimer) SetInterval(value int) {
	c.interval = value
}

// IsStarted checks if the timer is started.
//	Returns: bool true if the timer is started and false if it is stopped.
func (c *FixedRateTimer) IsStarted() bool {
	return c.ticker != nil
}

// Start starts the timer. Initially the timer is triggered after delay.
// After that it is triggered after interval until it is stopped.
func (c *FixedRateTimer) Start() {
	c.mtx.Lock()
	// Stop previously set timer
	c.stop()

	// Exit if interval is not defined
	if c.interval <= 0 {
		return
	}

	// Introducing delay
	delay := c.delay - c.interval
	ticker := time.NewTicker(time.Millisecond * time.Duration(c.interval))
	c.ticker = ticker

	c.mtx.Unlock()

	go func() {
		if delay > 0 {
			time.Sleep(time.Millisecond * time.Duration(delay))
		}

		for range ticker.C {
			callback := c.callback
			if callback != nil {
				callback()
			}
		}
	}()
}

// Stop the timer.
func (c *FixedRateTimer) Stop() {
	c.mtx.Lock()
	c.stop()
	c.mtx.Unlock()
}

// stop is a private function to implement thread save
func (c *FixedRateTimer) stop() {
	ticker := c.ticker
	if ticker != nil {
		ticker.Stop()
		c.ticker = nil
	}
}

// Close closes the timer.
// This is required by ICloseable interface, but besides that it is identical to stop().
//	Parameters: correlationId: string transaction id to trace execution through call chain.
//	Returns: error
func (c *FixedRateTimer) Close(correlationId string) error {
	c.mtx.Lock()
	c.stop()
	c.mtx.Unlock()
	return nil
}
