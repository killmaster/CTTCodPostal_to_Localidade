package ctt

import (
	"database/sql"
	"fmt"
	"log"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3" // SQL drivers
)

const (
	dbPath = "files/cploc.db"
)

var (
	db, err = sql.Open("sqlite3", filepath.FromSlash(dbPath))
)

func main() {
	fmt.Println(filepath.FromSlash(dbPath))
}

// CPToLocalidade recebe um código postal como string e devolve uma localidade
func CPToLocalidade(cp string) string {
	var query string
	switch len(cp) {
	case 4:
		query = fmt.Sprintf("SELECT localidade from cploc where cp4 like '%s'", cp)
	case 7:
		query = fmt.Sprintf("SELECT localidade from cploc where cp4 like '%s' and cp3 like '%s'", cp[0:4], cp[4:7])
	case 8:
		query = fmt.Sprintf("SELECT localidade from cploc where cp4 like '%s' and cp3 like '%s'", cp[0:4], cp[5:8])
	default:
		query = fmt.Sprintf("SELECT localidade from cploc where cp4 like '%s'", cp)
	}
	var localidade string
	rows, err := db.Query(query)
	checkError(err)
	defer rows.Close()
	rows.Next() 
	rows.Scan(&localidade)
	return localidade
}

func checkError(err error) {
	log.Println(err)
}
