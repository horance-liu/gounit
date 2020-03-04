package quantity

type Mile uint

func (m Mile) Equal(r Mile) bool {
	return m == r
}

func (m Mile) NotEqual(r Mile) bool {
	return !(m.Equal(r))
}