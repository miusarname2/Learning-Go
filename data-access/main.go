package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  string
}

func albumByArtist(name string) ([]Album, error) {
	var albums []Album

	rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)

	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}

	defer rows.Close()

	for rows.Next() {
		var alb Album

		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}

		albums = append(albums, alb)

	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}

	return albums, nil

}

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

	albums, err := albumByArtist("katiusha")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Albums found: %v\n", albums)

}
