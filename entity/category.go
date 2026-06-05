package entity

type Category string

const (
	TechCat    = "tech"
	HistoryCat = "history"
	SportsCat  = "sports"
)

func (c Category) IsValid() bool {

	switch c {
	case TechCat:

		return true
	case HistoryCat:

		return true
	case SportsCat:

		return true
	}

	return false
}
