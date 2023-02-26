package db

import "errors"

type ProxyDatabase struct {
	Users map[string]bool
	DB    *Database
}

var ErrInsufficientAccessRights = errors.New("insufficient access rights")

func (pDB *ProxyDatabase) GetData(user string) ([]string, error) {
	if !pDB.Users[user] {
		return nil, ErrInsufficientAccessRights
	}

	return pDB.DB.GetData(user)
}
