package filters

type Request interface{}

type Response interface{}

type Filter interface {
	Process(data Request) (Response, error)
}

type PipeLine struct {
	Name    string
	filters *[]Filter
}

func (p *PipeLine) Process(data Request) (Response, error) {
	var ret interface{}
	var err error
	for _, filter := range *p.filters {
		ret, err = filter.Process(data)
		if err != nil {
			return nil, err
		}
		data = ret
	}
	return ret, nil

}

// NewStraightPipe used to pipeline filters
func NewStraightPipe(name string, filters ...Filter) *PipeLine {
	return &PipeLine{
		Name:    name,
		filters: &filters,
	}
}
