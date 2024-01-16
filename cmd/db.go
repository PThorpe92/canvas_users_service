package cmd

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type DbHandler struct {
	Db *sql.DB
}

func NewDbHandler(db *sql.DB) *DbHandler {
	return &DbHandler{Db: db}
}

func Init() *sql.DB {
	db, err := sql.Open("sqlite3", "./canvasuserservice.db")
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	sqlStmt := `
		CREATE TABLE IF NOT EXISTS canvas_users (
			id INTEGER NOT NULL PRIMARY KEY,
			canvas_id INTEGER NOT NULL UNIQUE,
			name TEXT NOT NULL,
			created_at TEXT NOT NULL,
			sortable_name TEXT NOT NULL,
			short_name TEXT NOT NULL,
			sis_user_id TEXT NOT NULL UNIQUE,
			integration_id TEXT NOT NULL UNIQUE,
			login_id TEXT NOT NULL UNIQUE
		);

		CREATE TABLE IF NOT EXISTS unlocked_users (
			id INTEGER NOT NULL PRIMARY KEY,
			unlocked_id INTEGER NOT NULL UNIQUE,
			name_first TEXT NOT NULL,
			name_last TEXT NOT NULL,
			email TEXT NOT NULL UNIQUE,
			username TEXT NOT NULL UNIQUE
		);

		CREATE TABLE IF NOT EXISTS user_mapping (
			id INTEGER NOT NULL PRIMARY KEY,
			canvas_id INTEGER NOT NULL,
			unlocked_id INTEGER NOT NULL,
			FOREIGN KEY (canvas_id) REFERENCES canvas_users (canvas_id),
			FOREIGN KEY (unlocked_id) REFERENCES unlocked_users (unlocked_id)
		);
	`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatalf("Error creating tables: %v", err)
	}

	fmt.Println("Success, tables created")
	return db
}

func (db *DbHandler) InsertCanvasUsers(users []CanvasUser) {
	sqlStmt := `INSERT INTO canvasusers(canvas_id, name, created_at, sortable_name, short_name, sis_user_id, integration_id, login_id) VALUES(?,?,?,?,?,?,?,?)`
	for _, user := range users {
		_, err := db.Db.Exec(sqlStmt, user.Id, user.Name, user.CreatedAt, user.SortableName, user.ShortName, user.SisUserId, user.IntegrationId, user.LoginId)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("Success, canvas user inserted")
}

func (db *DbHandler) InsertUnlockedUsers(users []UnlockEdUser) {
	sqlStmt := `INSERT INTO unlockedusers(unlocked_id, name_first, name_last, email, username) VALUES(?,?,?,?,?)`
	for _, user := range users {
		_, err := db.Db.Exec(sqlStmt, user.Id, user.NameFirst, user.NameLast, user.Email, user.Username)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("Success, unlocked user inserted")
}

func (db *DbHandler) InsertUserMapping(user *UserMapping) {
	sqlStmt := `INSERT INTO usermapping(canvas_id, unlocked_id) VALUES(?,?)`
	_, err := db.Db.Exec(sqlStmt, user.ExternId, user.UnlockedId)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success, user mapping inserted")
}

func (db *DbHandler) QuerySimilarUsers(user *UnlockEdUser) []CanvasUser {
	fullName := user.NameFirst + " " + user.NameLast
	likeName := "%" + fullName + "%"
	sqlStmt := `SELECT * FROM canvasusers 
                WHERE name LIKE ? 
                OR sortable_name LIKE ? 
                OR short_name LIKE ? 
                OR sis_user_id LIKE ? 
                OR integration_id LIKE ? 
                OR login_id LIKE ?;`
	rows, err := db.Db.Query(sqlStmt, likeName, likeName, likeName, likeName, likeName, likeName)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var users []CanvasUser
	for rows.Next() {
		var user CanvasUser
		err = rows.Scan(&user.Id, &user.Name, &user.CreatedAt, &user.SortableName, &user.ShortName, &user.SisUserId, &user.IntegrationId, &user.LoginId)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}
	return users
}

func (db *DbHandler) QueryUserMapping(user *UnlockEdUser) []UserMapping {
	sqlStmt := `SELECT * FROM usermapping WHERE unlocked_id = ?`
	rows, err := db.Db.Query(sqlStmt, user.Id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var users []UserMapping
	for i := 0; rows.Next(); i++ {
		var user UserMapping
		err = rows.Scan(&user.Id, &user.ExternId, &user.UnlockedId)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}
	return users
}
