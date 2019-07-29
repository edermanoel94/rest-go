package request

import (
	"net/url"
	"strconv"
)

type PageRequest struct {
	Max    int
	Offset int
	Query  string
}

func NewPaginate(values url.Values) (*PageRequest, error) {
	return paginate(values)
}

func paginate(values url.Values) (*PageRequest, error) {

	pageRequest := &PageRequest{}

	query := values.Get("query")

	pageRequest.Query = query

	if offsetParam := values.Get("offset"); offsetParam != "" {
		offset, err := strconv.Atoi(offsetParam)
		if err != nil {
			return nil, err
		} else {
			pageRequest.Offset = offset
		}
	}

	if maxParam := values.Get("max"); maxParam != "" {
		max, err := strconv.Atoi(maxParam)
		if err != nil {
			return nil, err
		} else {
			pageRequest.Max = max
		}
	}

	return pageRequest, nil
}

func (p PageRequest) Sort(fields []string) {

}
