package database

import (
	"database/sql"
	"fmt"
	"url_shortener/internals/models"

	_ "github.com/mattn/go-sqlite3"
	"github.com/teris-io/shortid"
)

const (
	dbLocation = "database.sqlite3"
	urlTable   = "urls"
)

type Database struct{}

func (db Database) CreateDatabase() error {
	dbInstance, err := sql.Open("sqlite3", dbLocation)
	if err != nil {
		return err
	}

	query := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		url STRING NOT NULL,
		short_url STRING NOT NULL
	);`, urlTable)

	_, err = dbInstance.Exec(query)
	if err != nil {
		return err
	}

	fmt.Println("Database and table urls created successfully")
	return nil
}

func (db Database) CreateShortUrl(url string) (models.UrlRecord, error) {
	var urlRecord models.UrlRecord

	dbInstance, err := sql.Open("sqlite3", dbLocation)
	if err != nil {
		return urlRecord, err
	}

	shortUrl, err := shortid.Generate()
	if err != nil {
		return urlRecord, err
	}

	result, err := dbInstance.Exec("INSERT INTO urls(url, short_url) VALUES (?, ?)", url, shortUrl)
	if err != nil {
		return urlRecord, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return urlRecord, err
	}

	urlRecord.Id = int(lastId)
	urlRecord.Url = url
	urlRecord.ShortUrl = shortUrl

	return urlRecord, nil
}

func (db Database) QueryShortUrl(shortUrl string) (models.UrlRecord, error) {
	var urlRecord models.UrlRecord

	dbInstance, err := sql.Open("sqlite3", dbLocation)
	if err != nil {
		return urlRecord, err
	}

	content := dbInstance.QueryRow("SELECT id, url, short_url FROM urls WHERE short_url = ?", shortUrl)

	if content.Err() != nil {
		return urlRecord, content.Err()
	}

	content.Scan(urlRecord.Id, urlRecord.Url, urlRecord.ShortUrl)

	return urlRecord, nil
}

var Db = Database{}
