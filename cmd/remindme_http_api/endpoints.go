
package main 

import (
  "net/http"
  "encoding/json"

  "github.com/mrcrilly/remindme"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
  var err error 
  var getrsp GetResponse
  var tmpbk *remindme.Bookmark
  var jsondata []byte

  tmpbk, err = remindme.NewBookmark("http://mcrilly.me/", "Michael Crilly")
  checkError(err)

  getrsp.Unread = append(getrsp.Unread, tmpbk)
  getrsp.Read = nil 

  jsondata, err = json.Marshal(getrsp)
  checkError(err)

  _, err = w.Write(jsondata)
  checkError(err)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
  
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {

}
