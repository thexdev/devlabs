package primary

type Add interface {
	Execute(a float32, b float32) (float32, error)
}

type Subtract interface {
	Execute(a float32, b float32) (float32, error)
}

type Devide interface {
	Execute(a float32, b float32) (float32, error)
}

type Multiply interface {
	Execute(a float32, b float32) (float32, error)
}
