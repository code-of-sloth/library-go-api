package config

import (
	"context"
	"fmt"
)

func InitCore() (err error) {
	err = initDb()
	if err != nil {
		err = fmt.Errorf("initcore:error when connecting to db : %w", err)
		return
	}

	return
}

func InitClose() {
	DbConn.Close(context.Background())
}
