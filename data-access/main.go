package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Album struct {
	ID int64
	Title string
	Artist string
	Price float64
}

func main() {
    // Capture connection properties.
    cfg := mysql.NewConfig()
    cfg.User = "root"
    cfg.Passwd = "123"
    cfg.Net = "tcp"
    cfg.Addr = "127.0.0.1:3306"
    cfg.DBName = "recordings"

    // Get a database handle.
    var err error
    db, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal(err)
    }

    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    fmt.Println("Connected!")

	albums, err := albumsByArtist("John Coltrane")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", albums)

	album, err := AlbumById(2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Album found %v\n", album)

	lastId, err := addAlbum(Album{
		Title: "Ahmed Adel",
		Artist: "Mohsen",
		Price: 420,
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("new id: %d", lastId)

}

// albumsByArtist queries for albums that have the specified artist name.
func albumsByArtist(name string) ([]Album, error) {
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

	if err:= rows.Err();err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}

	return albums, nil

}

func AlbumById(id int64) (Album, error) {
	var album Album

	row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)
	if err := row.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
		if err == sql.ErrNoRows {
			return album, fmt.Errorf("Album by id %d no such column", id)
		}

		return album, fmt.Errorf("Album by id %d, has error: %v", id, err)
	}

	return album, nil
}

func addAlbum(album Album) (int64, error) {
	result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", album.Title, album.Artist, album.Price)
	
	if err != nil {
		return  0, fmt.Errorf("addAlbum err: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum result err: %v", err)
	}

	return id, nil
}