package main
 
import (
    "database/sql" 
    _ "github.com/go-sql-driver/mysql"
    "log"
	"fmt"
)
 
func main() {
    db, err := sql.Open("mysql", "root:Qwpo1209@tcp(localhost:3366)/")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

	var host []uint8
    var user string


    rows, err := db.Query("select host, user from mysql.user;")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()
	
    for rows.Next() {
        err := rows.Scan(&host, &user)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println(string(host), user)
    }
}
