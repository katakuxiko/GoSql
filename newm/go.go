package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)
type Artiles struct {
	Id uint16 `json:"id"`
	Title string `json:"title"`
	Anons string `json:"anons"`
	Full_text string `json:"full_text"`

}
var posts = []Artiles{}
var showPost = Artiles{}
func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("newm/templates/index.html", "newm/templates/header.html", "newm/templates/footer.html")
	if err != nil {
		log.Print(err)
	}

	db, err := sql.Open("mysql", "root:Vinter1973@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		panic(err)
	}
	defer db.Close()
//Выборка
	res, err := db.Query("SELECT* FROM `articles`")
	if err != nil{
		log.Println(err)	
		}
	posts = []Artiles{}	
	for res.Next(){
		var post Artiles
		err = res.Scan(&post.Id,&post.Title,&post.Anons,&post.Full_text)
		if err != nil{
			log.Println(err)
		
		}
		posts = append(posts, post)

}

	t.ExecuteTemplate(w, "index", posts	)
}
func create(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("newm/templates/create.html", "newm/templates/header.html", "newm/templates/footer.html")
	if err != nil {
		log.Print(err)
	}
	t.ExecuteTemplate(w, "create", nil)
}
func save_articles(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	anons := r.FormValue("anons")
	full_text := r.FormValue("full_text")
	if title == "" || full_text == "" || anons == "" {
		fmt.Fprintf(w, "Не все данные заполнены")
	}else{
	db, err := sql.Open("mysql", "root:Vinter1973@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	

	insert, err := db.Query(fmt.Sprintf("INSERT INTO `articles` (`title`,`anons`,`full_text`) VALUES('%s', '%s','%s')", title, full_text, anons))
	if err != nil {
		panic(err)
	}
	defer insert.Close()

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
}
func show_post(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	t, err := template.ParseFiles("newm/templates/show.html", "newm/templates/header.html", "newm/templates/footer.html")
	if err != nil {
		log.Print(err)
	}
	db, err := sql.Open("mysql", "root:Vinter1973@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	res, err := db.Query(fmt.Sprintf("SELECT* FROM `articles` WHERE `id` = '%s'",vars["id"]))
	if err != nil{
		log.Println(err)	
		}
	showPost = Artiles{}	
	for res.Next(){
		var post Artiles
		err = res.Scan(&post.Id,&post.Title,&post.Anons,&post.Full_text)
		if err != nil{
			log.Println(err)
		
		}
		showPost = post
	}
	t.ExecuteTemplate(w, "show", showPost)
}

func handleFunc() {
	rtr:=mux.NewRouter()
	rtr.HandleFunc("/", index).Methods("GET")
	rtr.HandleFunc("/create", create).Methods("GET")
	rtr.HandleFunc("/save_articles", save_articles).Methods("POST")
	rtr.HandleFunc("/post/{id:[0-9]+}", show_post).Methods("GET")

	http.Handle("/", rtr)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	
	http.ListenAndServe(":8080", nil)


}

func main() {
	handleFunc()

}
