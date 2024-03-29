package commands

import "github.com/pip-services3-gox/pip-services3-commons-gox/run"

// IEvent an interface for Events, which are part of the Command design pattern.
// Events allows sending asynchronous notifications to multiple subscribed listeners.
//	see IEventListener
type IEvent interface {
	run.INotifiable

	// Name gets the event name.
	//	Returns: string the name of the event.
	Name() string

	// Listeners gets all subscribed listeners.
	//	Returns: []IEventListener a list of listeners.
	Listeners() []IEventListener

	// AddListener adds a listener to receive notifications for this event.
	//	Parameters: listener: IEventListener the listener reference to add.
	AddListener(listener IEventListener)

	// RemoveListener removes a listener, so that it no longer receives notifications for this event.
	//	Parameters: listener: IEventListener the listener reference to remove.
	RemoveListener(listener IEventListener)
}
