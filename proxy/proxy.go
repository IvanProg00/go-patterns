package proxy

import (
	"fmt"

	"github.com/IvanProg00/go-patterns/proxy/db"
)

const (
	Administrator = "administrator"
	User          = "user"
)

func Run() {
	users := map[string]bool{
		Administrator: true,
		User:          false,
	}

	proxy := &db.ProxyDatabase{
		Users: users,
		DB:    &db.Database{},
	}

	admin, err := proxy.GetData(Administrator)
	fmt.Printf("from %q; data: %v; err: %v\n", Administrator, admin, err)

	user, err := proxy.GetData(User)
	fmt.Printf("from %q; data: %v; err: %v\n", User, user, err)
}
