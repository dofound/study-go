package main
import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	//链接数据库
	db,err := sql.Open("mysql","xxx:xxxx@tcp(xx.95.136.xxxx:3306)/mygo?charset=utf8")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	//插入数据
	var result sql.Result
	result,err = db.Exec("insert into my (age,name) values (?,?)",27,"肖建和")
	if err !=nil {
		fmt.Println(err)
		return
	}
	lastId,_:= result.LastInsertId()
	fmt.Println("新插入记录id为",lastId)

	//显示单条数据
	fmt.Println("获取单条记录")
	var row *sql.Row
	row = db.QueryRow("select * from my")
	var id,age uint8
	var name string
	err = row.Scan(&id,&name,&age)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(id,"\t",name,"\t",age)

	//显示多条数据
	fmt.Println("获取多条记录")
	var rows *sql.Rows
	rows,err = db.Query("select * from my")
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var name string
		var id,age int
		rows.Scan(&id,&name,&age)
		fmt.Println(id,"\t",name,"\t",age)
	}

	defer rows.Close()

	fmt.Printf("这是一个Mysql例子\n")
}
func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

/*
mygo 下需要的表
CREATE TABLE `my` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `age` int(10) unsigned NOT NULL DEFAULT '0',
  `name` varchar(45) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8 |
 */