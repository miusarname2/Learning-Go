package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	// Configuración de la conexión a la base de datos
	ctf := mysql.Config{
		User:                 "root",
		Passwd:               "", // Reemplaza "tu_contraseña" con la contraseña real
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "recordings",
		AllowNativePasswords: true,
	}

	// Intentar abrir la conexión a la base de datos
	var err error
	db, err = sql.Open("mysql", ctf.FormatDSN())
	if err != nil {
		log.Fatal("Error abriendo la conexión:", err)
	}
	defer db.Close() // Cerrar la conexión al final

	// Probar la conexión con Ping
	if err = db.Ping(); err != nil {
		log.Fatal("Error haciendo ping a la base de datos:", err)
	}

	fmt.Println("Conexión establecida")
}
