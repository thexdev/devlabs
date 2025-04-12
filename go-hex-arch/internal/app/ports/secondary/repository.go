package secondary

type Expression struct {
	A        float32
	B        float32
	Operator string
}

type Record struct {
	ID         string
	Expression Expression
	Result     float32
}

type Repository interface {
	All() []Record

	Save(record Record) (bool, error)
}
