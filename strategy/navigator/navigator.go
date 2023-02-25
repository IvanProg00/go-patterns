package navigator

type Navigator struct {
	Strategy
}

func (n *Navigator) SetStrategy(s Strategy) {
	n.Strategy = s
}
