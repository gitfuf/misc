package api

type House struct {
	Name      string
	Rooms     int
	HasGarden bool
	Toilets   int
}

type HouseOption func(*House)

//returns new house with specified options
func NewHouse(name string, opts ...HouseOption) (*House, error) {
	house := &House{
		Name:      name,
		HasGarden: false,
		Rooms:     2,
		Toilets:   1,
	}

	for _, opt := range opts {
		opt(house)
	}

	return house, nil
}

func AddGarden() HouseOption {
	return func(h *House) {
		h.HasGarden = true
	}
}

func Rooms(num int) HouseOption {
	return func(h *House) {
		h.Rooms = num
	}
}

func Toilets(num int) HouseOption {
	return func(h *House) {
		h.Toilets = num
	}
}
