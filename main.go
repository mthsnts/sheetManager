package main

import (
	"database/sql"
	"fmt"
	"github.com/xuri/excelize/v2"
	"strconv"
)

func ConectaComBancoDeDados() *sql.DB {
	conexao := "user=postgres dbname=recursos_didaticos. password=postgres host=localhost:5432 sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		fmt.Println(err)
	}

	return db
}
func main() {
	db := ConectaComBancoDeDados()
	f, err := excelize.OpenFile("./docs/tags.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	rows, _ := f.GetRows("Cursos RD v2 (Modificar)")
	rowLength := len(rows)

	for i := 2; i <= rowLength; i++ {
		columnE, _ := f.GetCellValue("Cursos RD v2 (Modificar)", fmt.Sprintf("E%s", strconv.Itoa(i)))
		columnF, _ := f.GetCellValue("Cursos RD v2 (Modificar)", fmt.Sprintf("F%s", strconv.Itoa(i)))
		_, _ = fmt.Printf("E%s", strconv.Itoa(i))
		sql := "UPDATE recursos_didaticos.tag SET nome = \"%s\" WHERE nome ILIKE(\"%s\")"

		fmt.Println(
			fmt.Sprintf(sql, columnF, columnE),
		)
	}

}
