package hepa

func WithFunc(p *Purifier, fn FilterFn) *Purifier {
	return With(p, fn)
}

func With(p *Purifier, f Filter) *Purifier {
	wp := &Purifier{
		parent: p,
		filter: f,
	}

	return wp
}
