package money

type pair struct {
	from string
	to   string
}

func (p *pair) Equals(obj interface{}) bool {
	objP, ok := obj.(pair)
	if !ok {
		return false
	}

	return p.from == objP.from && p.to == objP.to
}
