package seeder

import (
	"bufio"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"uk-hacktiv8-prakerja/config"
	"uk-hacktiv8-prakerja/database"
)

func main() {
	config.Init()
	database.StartDatabase()
	defer database.CloseConn()

	db := database.GetDatabase()
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	//Add seeders files
	path := []string{
		basepath + "/files/users.sql",
	}

	for _, val := range path {
		tx := db.Begin()
		file, err := os.Open(val)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			err := tx.Exec(scanner.Text())
			if err.Error != nil {
				tx.Rollback()
				break
			}
		}

		if err := scanner.Err(); err != nil {
			if err != io.EOF {
				tx.Rollback()
				break
			}

		}
		tx.Commit()
	}
}
