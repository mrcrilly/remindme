
package main 

import "github.com/mrcrilly/remindme"

type GetResponse struct {
  Unread []*remindme.Bookmark `json:"unread"`
  Read []*remindme.Bookmark `json:"read"`
}

type PostResponse struct {
  Status string `json:"status"`
}

type DeleteResponse struct {
  Status string `json:"status"`
}
