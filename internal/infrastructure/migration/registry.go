package migration

import (
	"gin-scaffold/internal/domain"
)

func Models() []interface{} {

	return []interface{}{

		&domain.User{},
	}

}
