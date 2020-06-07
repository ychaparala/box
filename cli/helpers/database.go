package helpers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	// sqlite3 drivers
	_ "github.com/mattn/go-sqlite3"
	homedir "github.com/mitchellh/go-homedir"
)

// PutUserData returns
func PutUserData(payload []byte) {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	database, _ := sql.Open("sqlite3", filepath.Join(home+"/.box/box.db"))
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS user (id INTEGER PRIMARY KEY, kind TEXT, idToken TEXT, email TEXT, refreshToken TEXT, expiresIn TEXT, localID TEXT)")
	statement.Exec()
	statement, _ = database.Prepare("INSERT INTO user (kind, idToken, email, refreshToken, expiresIn, localID) VALUES (?, ?, ?, ?, ?, ?)")
	var userJSON map[string]interface{}
	if err := json.Unmarshal(payload, &userJSON); err != nil {
		panic(err)
	}
	statement.Exec(userJSON["kind"], userJSON["idToken"], userJSON["email"], userJSON["refreshToken"], userJSON["expiresIn"], userJSON["localId"])
}

// GetUserData returns
func GetUserData() map[string]string {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	row := make(map[string]string)
	database, _ := sql.Open("sqlite3", filepath.Join(home+"/.box/box.db"))
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS user (id INTEGER PRIMARY KEY, kind TEXT, idToken TEXT, email TEXT, refreshToken TEXT, expiresIn TEXT, localID TEXT)")
	statement.Exec()
	rows, _ := database.Query("SELECT id, kind, idToken, email, refreshToken, expiresIn, localID from user order by id desc limit 1")

	var id int
	var kind, idToken, email, refreshToken, expiresIn, localID string

	for rows.Next() {
		rows.Scan(&id, &kind, &idToken, &email, &refreshToken, &expiresIn, &localID)
		row["id"] = strconv.Itoa(id)
		row["kind"] = kind
		row["idToken"] = idToken
		row["email"] = email
		row["refreshToken"] = refreshToken
		row["expiresIn"] = expiresIn
		row["localId"] = localID
	}
	return row
}

// PutUsageToBoxApp returns void
func PutUsageToBoxApp(payload []byte) {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	database, _ := sql.Open("sqlite3", filepath.Join(home+"/.box/box.db"))
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS boxappusage (id INTEGER PRIMARY KEY, email TEXT, limit INTEGER, usage INTEGER, usageInDrive INTEGER, usageInDriveTrash INTEGER)")
	statement.Exec()
	statement, _ = database.Prepare("INSERT INTO boxappusage (email, limit, usage, usageInDrive, usageInDriveTrash) VALUES (?, ?, ?, ?, ?)")
	var usageJSON map[string]interface{}
	if err := json.Unmarshal(payload, &usageJSON); err != nil {
		panic(err)
	}
	statement.Exec(usageJSON["email"], usageJSON["limit"], usageJSON["usage"], usageJSON["usageInDrive"], usageJSON["usageInDriveTrash"])
}

// GetUsageToBoxApp returns void
func GetUsageToBoxApp() map[string]int64 {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	row := make(map[string]int64)
	database, _ := sql.Open("sqlite3", filepath.Join(home+"/.box/box.db"))
	statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS boxappusage (id INTEGER PRIMARY KEY, email TEXT, limit INTEGER, usage INTEGER, usageInDrive INTEGER, usageInDriveTrash INTEGER)")
	statement.Exec()
	if err != nil {
		panic(err)
	}
	rows, _ := database.Query("SELECT SUM(limit), SUM(usage),SUM(usageInDrive),SUM(usageInDriveTrash) FROM boxappusage")
	var limit, usage, usageInDrive, usageInDriveTrash int64

	for rows.Next() {
		rows.Scan(&limit, &usage, &usageInDrive, &usageInDriveTrash)
		row["limit"] = limit
		row["usage"] = usage
		row["usageInDrive"] = usageInDrive
		row["usageInDriveTrash"] = usageInDriveTrash
	}
	return row
}
