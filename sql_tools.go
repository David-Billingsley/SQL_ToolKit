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
func sql_send(query string, server string, database string) {
	ConnStringQA := fmt.Sprintf("server=%s;user id=;database=%s;", server, database)

	conn2, err := sql.Open("sqlserver", ConnStringQA)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}

	rows2, err := conn2.Query(query)
	if err != nil {
		fmt.Println("Error reading records: ", err.Error())
	}
	defer rows2.Close()

}

// #region: SQL File Find
// This function takes and a file path and searches that path for files to read and pass into the sql_send function.
func (t *Data) SQL_File_Import(fpath string, server string, database string) bool {
	entries, err := os.ReadDir(fpath)
	count := 0
	if err != nil {
		log.Fatal(err)
	}
	for _, e := range entries {
		if strings.HasSuffix(e.Name(), ".sql") {
			query, err := os.ReadFile(fmt.Sprintf("%s%s", fpath, e.Name()))

			if err != nil {
				log.Fatal("File not read")
			}
			count = count + 1

			sql_send(string(query), server, database)
		}
	}

	if count != 0 {
		return true
	} else {
		return false
	}
}
