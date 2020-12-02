package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

type info struct{
    id int,
    name string,
}

func Dao() (*info, error) {
	return nil, sql.ErrNoRows
}

func Service() error {
	info, err := Dao()
	if err != nil{
	    return errors.Wrap(err,"no query rows")
    }
}

func main() {
	err := Service()
	fmt.Printf(err)
}


