package main

import (
	"fmt"

	mydb "github.com/Hiddan13/bookshop/pkg/config"
)

func main() {
	fmt.Println("good")

	db := mydb.InitDB("./employees.db") // инициализируем базу данных
}
