package main

import (
	_ "github.com/go-sql-driver/mysql"

	"sme-scaffold/cmd"
)

func main() {
	cmd.Execute()
}
