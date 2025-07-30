package usecase

type CatalogUseCase[T any] struct {
	repo CatalogRepository[T]
}

func NewCatalogUseCase[T any](r CatalogRepository[T]) *CatalogUseCase[T] {
	return &CatalogUseCase[T]{
		repo: r,
	}
}

func (uc *CatalogUseCase[T]) GetTrending() ([]T, error) {
	return uc.repo.GetTrending()
}
