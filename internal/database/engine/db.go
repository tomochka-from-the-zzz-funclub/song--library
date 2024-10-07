package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/tomochka-from-the-zzz-funclub/song-library/internal/config"
	database "github.com/tomochka-from-the-zzz-funclub/song-library/internal/database"
	myErrors "github.com/tomochka-from-the-zzz-funclub/song-library/internal/errors"
	myLog "github.com/tomochka-from-the-zzz-funclub/song-library/internal/logger"
	"github.com/tomochka-from-the-zzz-funclub/song-library/internal/models"
)

type Postgres struct {
	Connection *sql.DB
}

func NewPostgres(cfg config.Config) *Postgres {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s", cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBHost, cfg.DBPort, cfg.SslMode)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		myLog.Log.Fatalf("Failed to connect to PostgreSQL: %v", err)
		return nil
	}
	if err := db.Ping(); err != nil {
		myLog.Log.Fatalf("Failed to ping PostgreSQL: %v", err)
		return nil
	} else {
		myLog.Log.Debugf("ping success")
	}
	query := `
		CREATE TABLE songs (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			name VARCHAR(255) NOT NULL,
			author VARCHAR(255) NOT NULL,
			release_date DATE NOT NULL,
			text TEXT NOT NULL,
			link VARCHAR(255) NOT NULL
		);`
	_, err = db.Exec(query)
	if err != nil {
		myLog.Log.Errorf(err.Error())
	}

	return &Postgres{
		Connection: db,
	}
}

// Create new song
// func (db *Postgres) CreateSong(song database.ISong) (int, error) {
// 	query := `WITH insert_return AS (INSERT INTO songs (name, author, release_date, text, link) VALUES ($1, $2, $3, $4, $5) RETURNING id) SELECT id FROM insert_return`
// 	id := -1
// 	err := db.Connection.QueryRow(query, song.GetName(), song.GetGroup(), song.GetReleaseDateS(), song.GetText(), song.GetLink()).Scan(&id)
// 	if err != nil {
// 		myLog.Log.Errorf("Error CreateSong: %v", err.Error())
// 		return id, myErrors.ErrCreateSongDB
// 	}
// 	return id, nil
// }

func (db *Postgres) CreateSong(song database.ISong) (string, error) {
	query := `
        WITH insert_return AS (
            INSERT INTO songs (name, author, release_date, text, link)
            VALUES ($1, $2, $3, $4, $5)
            RETURNING id
        )
        SELECT id FROM insert_return
    `
	var id string
	err := db.Connection.QueryRow(query, song.GetName(), song.GetGroup(), song.GetReleaseDateS(), song.GetText(), song.GetLink()).Scan(&id)
	if err != nil {
		myLog.Log.Errorf("Error CreateSong: %v", err.Error())
		return "", myErrors.ErrCreateSongDB
	}
	return id, nil
}

func (db *Postgres) DeleteSong(id string) error {
	query := "DELETE FROM songs WHERE id = $1"
	result, err := db.Connection.Exec(query, id)
	if err != nil {
		myLog.Log.Errorf("Error DeleteSong: %v", err.Error())
		return myErrors.ErrNotDeleteDB
	}
	col, err := result.RowsAffected()
	if err != nil {
		myLog.Log.Errorf("Error RowsAffected:%v", err.Error())
		col = 0
	}
	if col == 0 {
		return myErrors.NotFoundDB
	}
	return nil
}

func (db *Postgres) FindIDByNameAndGroup(name string, group string) (bool, error) {
	var id string
	query := `SELECT id FROM songs WHERE name = $1 AND author = $2`
	err := db.Connection.QueryRow(query, name, group).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			// Если записи нет, возвращаем false и nil
			return false, nil
		}
		myLog.Log.Errorf("Error FindIDByNameAndGroup: %v", err.Error())
		// В любом другом случае возвращаем true и ошибку
		return true, myErrors.ErrFindSongDB
	}
	// Если запись найдена, возвращаем true и nil
	return true, nil
}

func (db *Postgres) GetFiltreSong(name, group, release, text, link string, number_records, page int) ([]models.Song, error) {
	query := `SELECT id, name, author, release_date, text, link FROM songs WHERE`
	i := 1
	flag_and := false
	vals := []interface{}{}
	if name != "" {
		if flag_and {
			query += fmt.Sprintf(" AND name = $%d", i)
			i++
			vals = append(vals, name)
		} else {
			query += fmt.Sprintf(" name = $%d", i)
			i++
			vals = append(vals, name)
			flag_and = true
		}
	}
	if group != "" {
		if flag_and {
			query += fmt.Sprintf(` AND author= $%d`, i)
			i++
			vals = append(vals, group)
			flag_and = true
		} else {
			query += fmt.Sprintf(` author= $%d`, i)
			i++
			vals = append(vals, group)
			flag_and = true
		}
	}
	if release != "" {
		releaseT, err := time.Parse("2006/01/02", release)
		if err != nil {
			if flag_and {
				query += fmt.Sprintf(" AND release_date = $%d", i)
				i++
				vals = append(vals, releaseT)
			} else {
				query += fmt.Sprintf(" release_date = $%d", i)
				i++
				vals = append(vals, releaseT)
				flag_and = true
			}
		}
	}
	if text != "" {
		if flag_and {
			query += fmt.Sprintf(" AND text LIKE CONCAT('%%',$%d::text,'%%')", i)
			i++
			vals = append(vals, text)
		} else {
			flag_and = true
			query += fmt.Sprintf(" text LIKE CONCAT('%%',$%d::text,'%%')", i)
			i++
			vals = append(vals, text)
		}
	}
	if link != "" {
		if flag_and {
			query += fmt.Sprintf(" AND link LIKE CONCAT('%%',$%d::text,'%%')", i)
			i++
			vals = append(vals, link)
		} else {
			flag_and = true
			query += fmt.Sprintf(" text LIKE CONCAT('%%',$%d::text,'%%')", i)
			i++
			vals = append(vals, link)
		}
	}
	skip := (page - 1) * number_records
	query += fmt.Sprintf(" LIMIT %d OFFSET %d", number_records, skip)
	songs := make([]models.Song, 0)
	rows, err := db.Connection.Query(query, vals...)
	defer rows.Close()
	if err != nil {
		myLog.Log.Errorf("Error GetFiltreSong in connect: %v", err.Error())
		return songs, myErrors.ErrGetSongWithFiltreDB
	}
	for rows.Next() {
		var song models.Song
		var grp, nm, rlsD, txt, lnk, id string
		err := rows.Scan(&id, &nm, &grp, &rlsD, &txt, &lnk)
		if err == nil {
			song.SetGroup(grp)
			song.SetName(nm)
			layout := time.RFC3339
			releaseDate, err := time.Parse(layout, rlsD)
			if err == nil {
				song.SetReleaseDate(releaseDate)
			}
			song.SetText(txt)
			song.SetLink(lnk)
			song.ID = id
			songs = append(songs, song)
		} else {
			myLog.Log.Errorf("Error GetFiltreSong in scan: %v", err.Error())
		}
	}
	return songs, nil
}

func (db *Postgres) UpdateSong(song models.Song) error {
	query := `UPDATE songs SET name = $1, author = $2, release_date = $3, text = $4, link = $5 WHERE id = $6`
	//i := strconv.Itoa(song.ID)
	result, err := db.Connection.Exec(query, song.Name, song.Group, song.GetReleaseDateS(), song.Text, song.Link, song.ID)
	if err != nil {
		myLog.Log.Errorf("Error UpdateSong: %v", err.Error())
		return myErrors.ErrUpdateSongDB
	}
	col, err := result.RowsAffected()
	if err != nil {
		myLog.Log.Errorf("Error RowsAffected:%v", err.Error())
		col = 0
	}
	if col == 0 {
		return myErrors.NotFoundDB
	}
	return nil
}

func (db *Postgres) GetText(id string) (string, error) {
	var text string
	query := `SELECT text FROM songs WHERE id = $1`
	err := db.Connection.QueryRow(query, id).Scan(&text)
	if err != nil {
		myLog.Log.Errorf("Error GetText: %v", err.Error())
		return "", myErrors.ErrGetTextDB
	}
	return text, nil
}
