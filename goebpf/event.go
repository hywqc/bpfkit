package goebpf

const (
	EVENTTYPE_TEST = iota
	EVENTTYPE_PROCESS_EXIT
	EVENTTYPE_TCP_TX
	EVENTTYPE_TCP_RX
	EVENTTYPE_TCP_RETRANS
)

type IEvent interface {
	Type() int
}

type EventHandler func(event IEvent)

type EventProcessExit struct {
	Pid      uint32
	Comm     string
	ExitCode int32
}

func (e *EventProcessExit) Type() int {
	return EVENTTYPE_PROCESS_EXIT
}
