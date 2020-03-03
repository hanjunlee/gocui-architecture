package service

// UseCase is the use cases in UI.
type UseCase interface {
	Search(pattern string) []string
}
