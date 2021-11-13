package hepa

func WithFunc(p *Purifier, fn FilterFn) *Purifier {
	return With(p, fn)
}

func With(p *Purifier, f Filter) *Purifier {
	c := p.clone()
	og := c.filter

	fn := FilterFn(func(b []byte) ([]byte, error) {
		var err error
		if og != nil {
			b, err = og.Filter(b)
			if err != nil {
				return nil, err
			}
		}
		return f.Filter(b)
	})

	c.filter = fn

	return c
}
