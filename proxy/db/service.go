package db

type Service interface {
	GetData(user string) ([]string, error)
}
