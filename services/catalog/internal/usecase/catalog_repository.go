package usecase

type CatalogRepository[T any] interface {
	GetTrending() ([]T, error)
}
