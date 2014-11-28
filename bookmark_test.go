
package remindme

import "testing"

func checkerror(err error, t *testing.T) {
  if err != nil {
    t.Fatal(err)
  }
}

var test_bookmarks = []struct {
  Url string
  Comment string 
  ShouldFail bool
  Error string
} {
  {"www.mcrilly.me/example/url", "Michael Crilly", false, ""},
  {"http://yahoo.co.me", "Yahoo!", false, ""},
  {"https://www.google.com/?isvalue=true", "Google", false, ""},
  {"invalid-address", "Fake Title", true, "Invalid URL Provided"},
}

var test_urls = []struct {
  Url string
  ShouldFail bool 
  Error string 
} {
  {"not-valid-at-all", true, "Invalid URL Provided"},
  {"http://very-valid.example", false, ""},
}

func TestNewBookmark(t *testing.T) {
  var tmpbm *Bookmark 
  var err error

  for _, i := range test_bookmarks {
    tmpbm, err = NewBookmark(i.Url, i.Comment)

    if err == nil && i.ShouldFail {
      t.Log(string(i.Url))
      t.Fatal("This test should have failed")
    }

    if err != nil && i.ShouldFail {
      if err.Error() != i.Error {
        t.Fatalf("Incorrect error message returned: %s", err.Error())
      }

      t.Skip()
    }

    if tmpbm.Url != i.Url {
      t.Fatal("URL mismatch: %s versus %s", tmpbm.Url, i.Url)
    }

    if tmpbm.Comment != i.Comment {
      t.Fatal("Comment mismatch: %s versus %s", tmpbm.Comment, i.Comment)
    }
  }
}

func TestIsValidUrl(t *testing.T) {
  var err error 

  for _, u := range test_urls {
    err = IsValidUrl(u.Url)

    if err == nil && u.ShouldFail {
      t.Log(string(u.Url))
      t.Fatal("This test should have failed")
    }

    if err != nil && u.ShouldFail {
      if err.Error() != u.Error {
        t.Log(string(u.Url))
        t.Fatalf("Incorrect error message received: %s", err.Error())
      }

      t.Skip()
    }
  }
}


