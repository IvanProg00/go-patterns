package db

type Database struct{}

func (db *Database) GetData(user string) ([]string, error) {
	return []string{
		"line 1",
		"end line",
	}, nil
}
