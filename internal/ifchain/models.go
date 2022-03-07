package ifchain

type Object struct {
	Name    string
	Version string
	state   string
}

func (o *Object) UpdateState(state string) {
	o.state = state
}
