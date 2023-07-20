package main

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

var schema = `
CREATE TABLE course (
    primary_id 	integer,
    owner_id 	integer,
    info 		text,
    section_id 	integer[]               
);

CREATE TABLE section (
    primary_id 	integer,
    cource_id 	integer,
    themes_id 	integer[]
);

CREATE TABLE theme (
    primary_id integer,
    section_id integer,
    content text
)`



type Course struct {
	primary_id	int64		`db:"primary_id"`
	owner_id	int64		`db:"owner_id"`
	info		string		`db:"info"`
	section_id	[]int64		`db:"section_id"`
}

type Section struct {
	primary_id	int64		`db:"primary_id"`
	course_id	int64		`db:"course_id"`
	themes_id	[]int64		`db:"themes_id"`
}

type Theme struct{
	primary_id	int64		`db:"primary_id"`
	section_id	int64		`db:"section_id"`
	content		string		`db:"content"`
}

func createTh(ctx context.Context, newTheme Theme){
	tx := ctx.Value("tx").(*sqlx.Tx)
	tx.Query("INSERT INTO themes VALUES ($1, $2, $3);", newTheme.primary_id, newTheme.section_id, newTheme.content)
}

func createTheme(ctx context.Context, newTheme Theme){
	tx := ctx.Value("tx").(*sqlx.Tx)
	tx.Query("INSERT INTO themes VALUES ($1, $2, $3);", newTheme.primary_id, newTheme.section_id, newTheme.content)
}

func readTheme(ctx context.Context, themeId int64) Theme{
	tx := ctx.Value("tx").(*sqlx.Tx)
	var primaryId, sectionId int64
	var content string
	query,_ := tx.Query("SELECT * from theme WHERE primary_id = $1;", themeId)
	query.Next()
	query.Scan(&primaryId, &sectionId, &content)
	return Theme{primary_id: primaryId, section_id: primaryId, content: content}
}



func main() {
	db, err := sqlx.Open("postgres", "host=localhost dbname=habrdb sslmode=disable user=habrpguser password=pgpwd4habr")
	if err != nil {
		log.Fatalln(err)
	}
	myctx := context.Background()
	tx := db.MustBegin()
	ctx := context.WithValue(myctx, "tx", tx)
	theme := readTheme(ctx, 137)

	fmt.Println(theme.content)
}




