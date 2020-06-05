package dp05_PrototypePattern

func (p *Prototype) Copy() *Prototype {
	proto := p
	return proto
}

func (p *Prototype) DeepCopy() *Prototype {
	proto := new(Prototype)
	proto.Data = p.Data
	return proto
}
