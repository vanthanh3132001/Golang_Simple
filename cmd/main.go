package main

import "github.com/example/db"

func main() {
	var sql = new(db.SQL)
	sql.Connect()
	defer sql.Close()

}
