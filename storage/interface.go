package storage

type Interface interface {
	Get(string) (string, error)
	Set(string, string) error
}
