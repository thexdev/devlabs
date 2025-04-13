package primary

import "gohexarch/internal/app/ports/secondary"

type History interface {
	Execute() ([]secondary.Record, error)
}
