package views

type Model struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
}

func InitialModel() Model {
	return Model{
		choices:  []string{"Buy carrots", "Buy celery", "Buty kohlrabi"},
		selected: make(map[int]struct{}),
		cursor:   0,
	}
}
