package main

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

type Course struct {
	primary_id int64   `db:"primary_id"`
	owner_id   int64   `db:"owner_id"`
	info       string  `db:"info"`
	section_id []int64 `db:"section_id"`
}

type Section struct {
	primary_id int64   `db:"primary_id"`
	course_id  int64   `db:"course_id"`
	themes_id  []int64 `db:"themes_id"`
}

type Theme struct {
	primary_id int64  `db:"primary_id"`
	section_id int64  `db:"section_id"`
	content    string `db:"content"`
}

const (
	createRhemeStmt        = "INSERT INTO theme VALUES($1, $2, $3);"
	updateThemeStmt        = "UPDATE theme SET section_id = $1, content = $2 WHERE primary_id = $3;"
	deleteThemeStmt        = "DELETE FROM theme WHERE section_id = $1;"
	readThemeStmt          = "SELECT * from theme WHERE primary_id = $1;"
	readThemeSectionStmt   = "SELECT (section_id) FROM theme WHERE primary_id = $1;"
	updateThemeSectionStmt = "UPDATE theme SET section_id = $1 WHERE primary_id = $2;"
	readThemeContentStmt   = "SELECT(content) FROM theme WHERE primary_id = $1;"
	updateThemeContentStmt = "UPDATE theme SET content = $1 WHERE primary_id = $2;"

	createThemeTableStmt = `CREATE TABLE theme(
  	  primary_id 	integer primary key,
  	  owner_id 	integer,
  	  info 		text,
  	  section_id 	integer[]               
	);`

	createSectionTableStmt = `CREATE TABLE section(
  	  	primary_id 	integer primary key,
  	  	courser_id 	integer,
  	 	themes_id 	integer[]               
	);`

	createCourseTableStmt = `CREATE TABLE section(
  	  	primary_id 	integer primary key,
  	  	owner_id 	integer,
  	 	info		text,
  	 	section_id  integer[]          
	);`
)

/*func CreateOrUseExistedSchema(db *sqlx.DB, ctx context.Context) error {
	err := db.CreateSchema(ctx)
	if err != nil && strings.Contains(err.Error(), "already exists") {
		log.Println("Existed " + table_name + " schema will be used")
		return nil
	} else {
		return err
	}
}

func CreateSchema(db *sqlx.DB, ctx context.Context) error {
	_, err := db.ExecContext(ctx, schema)
	if err != nil {
		log.Println("CreateSchema: " + err.Error())
	} else {
		log.Println("Schema ", table_name, " created")
	}
	return err
}*/

func createTheme(ctx context.Context, db *sqlx.DB, primary_id int64, section_id int64, content string) {
	db.ExecContext(ctx, createRhemeStmt, primary_id, section_id, content)
}

func updateTheme(ctx context.Context, db *sqlx.DB, section_id int64, primary_id int64, content string) {
	db.ExecContext(ctx, updateThemeContentStmt, section_id, content, primary_id)
}

func deleteTheme(ctx context.Context, db *sqlx.DB, primary_id int64) {
	db.ExecContext(ctx, deleteThemeStmt, primary_id)
}

func readTheme(ctx context.Context, db *sqlx.DB, primary_id int64) Theme {
	var (
		primaryId, sectionId int64
		content              string
	)
	query, _ := db.QueryContext(ctx, readThemeStmt, primary_id)
	query.Next()
	query.Scan(&primaryId, &sectionId, &content)
	return Theme{primary_id: primaryId, section_id: primaryId, content: content}
}

func readThemeSection(ctx context.Context, db *sqlx.DB, primary_id int64) int64 {
	query, _ := db.QueryContext(ctx, readThemeContentStmt, primary_id)
	query.Next()
	var res int64
	query.Scan(&res)
	return res
}

func updateThemeSection(ctx context.Context, db *sqlx.DB, primary_id int64, section_id int64) {
	db.ExecContext(ctx, updateThemeStmt, section_id, primary_id)
}

func readThemeContent(ctx context.Context, db *sqlx.DB, primary_id int64) string {
	query, _ := db.QueryContext(ctx, readThemeContentStmt, primary_id)
	query.Next()
	var res string
	query.Scan(&res)
	return res
}

func updateThemeContent(ctx context.Context, db sqlx.DB, primary_id int64, content string) {
	db.ExecContext(ctx, updateThemeContentStmt, content, primary_id)
}

func main() {
	db, err := sqlx.Open("postgres", "host=localhost dbname=habrdb sslmode=disable user=habrpguser password=pgpwd4habr")
	if err != nil {
		log.Fatalln(err)
	}
	ctx := context.Background()
	//db.MustExecContext(ctx, schema)
	db.ExecContext(ctx, createSectionTableStmt)
	db.ExecContext(ctx, createCourseTableStmt)
	db.ExecContext(ctx, createThemeTableStmt)
	createTheme(ctx, db, 5, 8, "attack_heli")
	newTheme := readTheme(ctx, db, 5)
	fmt.Print(newTheme.content)
}
