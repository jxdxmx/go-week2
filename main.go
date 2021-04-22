package main

import (
	"database/sql"
	"github.com/pkg/errors"
)

func main() {
	var err error
	if err == sql.ErrNoRows {
		err = errors.WithStack(err)
	}
}
