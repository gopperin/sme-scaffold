package main

import (
	_ "github.com/go-sql-driver/mysql"

	"sme-stage/cmd"
)

func main() {
	cmd.Execute()
}
