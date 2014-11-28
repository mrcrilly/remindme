
package main 

import (
  "fmt"
  "flag"

  "github.com/mrcrilly/remindme"
)

func check_error(err error) {
  if err != nil {
    fmt.Println(err.Error())
  }
}

func main() {
  var given_url string 
  var given_comment string 
  var given_database string 

  var bookmark *remindme.Bookmark
  var err error

  flag.StringVar(&given_url, "url", "", "The URL to add to the database")
  flag.StringVar(&given_comment, "comment", "", "The comment to help identify the bookmark")
  flag.StringVar(&given_database, "database", "/tmp/bookmarks.sqlite3.db", "The sqlite3 database to write to")
  flag.Parse()

  err = remindme.BringUpDatabase("/tmp/crilly.bookmarks")
  check_error(err)

  bookmark, err = remindme.NewBookmark(given_url, given_comment)
  check_error(err)

  err = remindme.AddBookmark(bookmark)
  check_error(err)
}