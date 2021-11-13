package hepa

func WithFunc(p *Purifier, fn FilterFn) *Purifier {
	return With(p, fn)
}

func With(p *Purifier, f Filter) *Purifier {
	c := p.clone()

	c.filter = f

	return c
}
