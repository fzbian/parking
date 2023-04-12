package main

import (
	"fmt"
	"github.com/fzbian/parking/config"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	err := config.TestDBConnection()
	if err != nil {
		fmt.Println("Error de conexión a la base de datos:", err)
		return
	}

	fmt.Println("Conexión exitosa a la base de datos")
}
