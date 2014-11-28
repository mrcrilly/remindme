
package remindme

import (
  "errors"

  "github.com/jinzhu/gorm"
  _ "github.com/mattn/go-sqlite3"
)

var Bookmarks gorm.DB

func BringUpDatabase(conns string) (err error) {
  Bookmarks, err = gorm.Open("sqlite3", conns)

  Bookmarks.AutoMigrate(&Bookmark{})

  return err 
}

func CloseDatabase() (err error) {
  err = Bookmarks.DB().Close()
  return err 
}

func AddBookmark(b *Bookmark) (err error) {
  if b == nil {
    return errors.New("Nil Bookmark given")
  }

  Bookmarks.Create(b)

  return nil
}

func RemoveBookmark(b *Bookmark) (err error) {
  if b == nil {
    return errors.New("Nil Bookmark given")
  }

  Bookmarks.Delete(b)

  return nil
}

func RetrieveBookmark(url string) (b *Bookmark, err error) {
  if len(url) <= 0 {
    return nil, errors.New("Zero length URL provided")
  }

  b = &Bookmark{}
  query := Bookmarks.Where(&Bookmark{Url: url}).First(b)

  if query.Error != nil {
    return nil, query.Error
  }

  return b, nil
}

func PurgeDatabase() (err error) {
  return errors.New("Not implemented")
}