package commands

import (
	"context"
	"github.com/pip-services3-gox/pip-services3-commons-gox/run"
)

// IEventListener an interface for listener objects that receive notifications on fired events.
//	see IEvent
//	see Event
//	Example:
//		type MyListener {
//			msg string;
//		}
//		func (l* MyListener) onEvent(correlationId string, event IEvent, args Parameters) {
//			fmt.Println("Fired event " + event.Name());
//		}
//
//		var event = NewEvent("myevent");
//		_listener := MyListener{};
//		event.addListener(_listener);
//		event.notify("123", Parameters.FromTuples("param1", "ABC"));
//
//		// Console output: Fired event myevent
type IEventListener interface {
	// OnEvent a method called when events this listener is subscrubed to are fired.
	//	Parameters:
	//		- ctx context.Context
	//		- correlationId: string (optional) transaction id to trace execution through call chain.
	//		- e: IEvent a fired evemt
	//		- value: *run.Parameters event arguments.
	OnEvent(ctx context.Context, correlationId string, e IEvent, value *run.Parameters)
}
