package domain

type Status int

const (
	StateIdle Status = iota
	StateFinished
	StateError
)

var stateName = map[Status]string{
	StateIdle:     "idle",
	StateFinished: "finished",
	StateError:    "error",
}

func (s Status) String() string {
	return stateName[s]
}
