package DataBaseTools

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/denisenkom/go-mssqldb"
)

type Data struct {
	MaxFileSize      int
	AllowedFileTypes []string
}

// #region: SQL Insert
// This function takes the server string and database string from the SQL_File_Import and then also takes the file from the
// File path file reads the whole file as a string and send to the server
func sql_send(query string, server string, database string, user string, pass string) {
	var ConnString string
	// Checks for user to be passed, if no user is passed then generate the windows auth string
	if len(user) > 0 {
		ConnString = fmt.Sprintf("server=%s;user id=;database=%s;", server, database)
	} else {
		ConnString = fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s;", server, user, pass, database)
	}
	conn, err := sql.Open("sqlserver", ConnString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}

	defer conn.Close()

	rows, err := conn.Query(query)
	if err != nil {
		fmt.Println("Error reading records: ", err.Error())
	}
	defer rows.Close()

}

// #region: SQL File Find
// This function takes a file path and searches that path for files to read and pass into the sql_send function.
func (t *Data) SQL_File_Import(fpath string, server string, database string, user string, pass string) (map[string]string, bool) {
	// Creates a map to hold the fileanems found below.
	var filefound = make(map[string]string)
	entries, err := os.ReadDir(fpath)
	count := 0
	if err != nil {
		log.Fatal(err)
	}
	// Loops over the files in the directory search above
	for _, e := range entries {
		// If the file name as .sql as a suffix run the import else fail
		if strings.HasSuffix(e.Name(), ".sql") {
			query, err := os.ReadFile(fmt.Sprintf("%s%s", fpath, e.Name()))

			if err != nil {
				log.Fatal("File not read")
			}
			count = count + 1

			// Insert filename with an index as the key
			filefound[fmt.Sprintf("%d", count)] = e.Name()

			sql_send(string(query), server, database, user, pass)
		}
	}

	// If the number of files imported is not equal to zero then return filenames and true to let the user know it has been completed.
	if count != 0 {
		return filefound, true
	} else {
		// If equal to zero return with filenames and false to let the user know of its failure.
		return filefound, false
	}
}
