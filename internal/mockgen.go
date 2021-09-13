package internal

//go:generate mockgen -destination=./mocks/flusher_mock.go -package=mocks github.com/ozonva/ova-person-api/internal/flusher Flusher

//go:generate mockgen -destination=./mocks/repo_mock.go -package=mocks github.com/ozonva/ova-person-api/internal/repo PersonRepo
