package nsq

import (
	"errors"
	"fmt"
)

// returned when a publish command is made against a Writer that is not connected
var ErrNotConnected = errors.New("not connected")

// returned when a publish command is made against a Writer that has been stopped
var ErrStopped = errors.New("stopped")

// returned from ConnectToNSQ when already connected
var ErrAlreadyConnected = errors.New("already connected")

// returned from updateRdy if over max-in-flight
var ErrOverMaxInFlight = errors.New("over configure max-inflight")

// returned from ConnectToLookupd when given lookupd address exists already
var ErrLookupdAddressExists = errors.New("lookupd address already exists")

type ErrIdentify struct {
	Reason string
}

func (e ErrIdentify) Error() string {
	return fmt.Sprintf("failed to IDENTIFY - %s", e.Reason)
}

type ErrProtocol struct {
	Reason string
}

func (e ErrProtocol) Error() string {
	return e.Reason
}
