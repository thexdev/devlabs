package usecases

import "gohexarch/internal/app/ports/secondary"

type RetrieveHistoryUseCase struct {
	repo secondary.Repository
}

func NewRetrieveHistoryUseCase(
	repo secondary.Repository,
) *RetrieveHistoryUseCase {
	return &RetrieveHistoryUseCase{
		repo: repo,
	}
}

func (uc *RetrieveHistoryUseCase) Execute() ([]secondary.Record, error) {
	histories := uc.repo.All()

	if len(histories) == 0 {
		return []secondary.Record{}, nil
	}

	return histories, nil
}
