package main

import (

	"fmt"
	"log"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)
type User struct {
	Name string `json:"name"`
	Age uint16 `json:"age"`

}
func main(){
	db, err := sql.Open("mysql", "root:Vinter1973@tcp(127.0.0.1:3306)/golang")
	if err != nil{
		panic(err)
	}
	defer db.Close()

	insert, err := db.Query("INSERT INTO `users` (`name`,`age`) VALUES('A', 425)")
	if err != nil{
		panic(err)
	}
	defer insert.Close()
	
	//Выборка
	res, err := db.Query("SELECT `name`,`age` FROM `users` ")
	if err != nil{
		log.Println(err)
			
		}
	for res.Next(){
		var user User
		err = res.Scan(&user.Name,&user.Age)
		if err != nil{
			log.Println(err)
		
		}
		fmt.Printf("User: %s with age %d \n", user.Name,user.Age)
	   
	}
	
}	