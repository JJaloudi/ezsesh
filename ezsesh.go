package ezsesh

type EzSesh struct {
	Store EzStoreMethods
}

func New() *EzSesh {
	return &EzSesh{}
}
