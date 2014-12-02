
package main 

import (
  "net/http"
  "database/sql"

  "github.com/jinzhu/gorm"
  _ "github.com/mattn/go-sqlite3"
  "github.com/mrcrilly/remindme"
)

func checkError(err error) {
  if err != nil {
    panic(err)
  }
}

func main() {
  var regex_handler *RegexpHandlers
  regex_handler = &RegexpHandlers{}

  remindme.BringUpDatabase("/tmp/api_database.sqlite")

  regex_handler.HandleFunc("^/$", http.HandlerFunc(GetHandler))
  regex_handler.HandleFunc(`^/add/\b(([\w-]+://?|www[.])[^\s()<>]+(?:\([\w\d]+\)|([^[:punct:]\s]|/)))$`, http.HandlerFunc(PostHandler))
  regex_handler.HandleFunc("^/list$", http.HandlerFunc(GetHandler))

  http.ListenAndServe(":8080", regex_handler)
}
