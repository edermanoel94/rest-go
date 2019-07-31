package rest

type PageRequest struct {
	Max       int
	Offset    int
	Query     string
	Direction rune
}

func NewPageRequest(offset, max int, direction string, properties ...string) (*PageRequest, error) {
	return paginate(offset, max, direction, properties...)
}

func paginate(offset, max int, direction string, properties ...string) (*PageRequest, error) {
	return nil, nil
}
