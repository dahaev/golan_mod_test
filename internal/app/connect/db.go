package connect

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db_count = 0

func db(pro *Projects) {
	db, err := sql.Open("mysql", "root:@/ANORSI")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()
	db_count += len(pro.Values)
	for _, res := range pro.Values {
		_, err := db.Exec("insert into tasks (`projectID`, `projectName`, `taskName`, `startDate`, `endDate`) VALUES(?, ?, ?,?,?)",
			res.ProjectID, res.ProjectName, res.TaskName, res.StartDate, res.EndDate)
		if err != nil {
			fmt.Println("Значения при ошибке", res.ProjectID, res.ProjectName, res.TaskName)
			fmt.Println("Ошибка во вводе")
			log.Fatal()
		}
	}
	fmt.Println("Инъектировано: ", db_count)

}
