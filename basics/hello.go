package kubeless

import (
	"time"

	"github.com/kubeless/kubeless/pkg/functions"
)

func Handler(event functions.Event, context functions.Context) (string, error) {
	return event.Data, nil
}

func Foo(event functions.Event, context functions.Context) (string, error) {
	select {
	case <-event.Extensions.Context.Done():
		return "", nil
	case <-time.After(5 * time.Second):
	}
	return "Function returned after 5 seconds", nil
}
