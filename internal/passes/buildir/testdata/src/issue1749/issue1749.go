package issue1749

type T struct{}

func (T) M[P any]() P {
	var zero P
	return zero
}

func G[P any]() P {
	var zero P
	return zero
}

func F(t T) string {
	f := t.M[string]
	return f() + G[string]()
}
