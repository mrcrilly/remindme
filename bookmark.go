
package remindme

import (
  "time"
  "errors"
  "regexp"
)

type Bookmark struct {
  Url string `json:"url"`
  Comment string `json:"comment"`
  Added time.Time `json:"added"`
}

func IsValidUrl(url string) (err error) {
  match, err := regexp.MatchString(`\b(([\w-]+://?|www[.])[^\s()<>]+(?:\([\w\d]+\)|([^[:punct:]\s]|/)))`, url)
  
  if err != nil {
    return err
  }

  if ! match {
    return errors.New("Invalid URL Provided")
  }

  return nil
}

func IsValidComment(comment string) (err error) {
  match, err := regexp.MatchString(`^[a-zA-Z0-9 ]+$`, comment)

  if err != nil {
    return err 
  }

  if ! match {
    return errors.New("Invalid Comment Format")
  }

  return nil 
}

func NewBookmark(url, comment string) (b *Bookmark, err error) {
  var tmpbm *Bookmark 

  err = IsValidUrl(url)

  if err != nil {
    return nil, err 
  }

  err = IsValidComment(comment)

  if err != nil {
    return nil, err 
  }

  tmpbm = &Bookmark{
    Url: url,
    Comment: comment,
    Added: time.Now(),
  }

  return tmpbm, nil 
}

