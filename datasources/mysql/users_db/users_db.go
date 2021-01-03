package users_db

import(
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

const(
	mysql_uers_username = "mysql_uers_username"
	mysql_uers_password = "mysql_uers_password"
	mysql_uers_host = "mysql_uers_host"
	mysql_uers_schema = "mysql_uers_schema"


)

var(
	Client *sql.DB
	
	username = os.Getenv(mysql_uers_username)
	password = os.Getenv(mysql_uers_password)
	host = os.Getenv(mysql_uers_host)
	schema = os.Getenv(mysql_uers_schema)
)

func init(){
	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset-utf8",
	username, password, host, schema,
)
	var err error
	Client, err =  sql.Open("mysql", datasourceName)
	if err != nil{
		panic(err)
	}
	if err = Client.Ping(); err != nil{
		panic(err)
	}
	log.Println("database successfully configured")
}