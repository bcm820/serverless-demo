package kubeless

import (
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/kubeless/kubeless/pkg/functions"
	"github.com/sirupsen/logrus"
)

func Handler(event functions.Event, context functions.Context) (string, error) {
	// return event.Data, nil
	return "hello decipher", nil
}

//this can be used for timeout events
func Foo(event functions.Event, context functions.Context) (string, error) {
	select {
	case <-event.Extensions.Context.Done():
		return "", nil
	case <-time.After(5 * time.Second):
	}
	return "Function returned after 5 seconds", nil
}

func IsPrime(event functions.Event, context functions.Context) (string, error) {
	num, err := strconv.Atoi(event.Data)
	if err != nil {
		return "", fmt.Errorf("Failed to parse %s as int! %v", event.Data, err)
	}
	logrus.Infof("Checking if %s is prime", event.Data)
	if num <= 1 {
		return fmt.Sprintf("%d is not prime", num), nil
	}
	for i := 2; i <= int(math.Floor(float64(num)/2)); i++ {
		if num%i == 0 {
			return fmt.Sprintf("%d is not prime", num), nil
		}
	}
	return fmt.Sprintf("%d is prime", num), nil
}
