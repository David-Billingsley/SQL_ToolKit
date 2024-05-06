package DataBaseTools

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
)

type Data struct {
	MaxFileSize      int
	AllowedFileTypes []string
}

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

func (t *Data) SQL_File_Import(fpath string, server string, database string) {
	entries, err := os.ReadDir(fpath)
	if err != nil {
		log.Fatal(err)
	}
	for _, e := range entries {
		fmt.Println(e.Name())
		query, err := os.ReadFile(fmt.Sprintf("%s%s", fpath, e.Name()))

		if err != nil {
			log.Fatal("File not read")
		}

		sql_send(string(query), server, database)
	}
}
