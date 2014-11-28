
package remindme

import "testing"

var fodder = []struct {
  Url string 
  Comment string
  ShouldFail bool 
  Error string 
} {

  // Should pass:
  {"www.mcrilly.me/example/url", "Michael Crilly", false, ""},
  {"http://yahoo.co.me", "Yahoo!", false, ""},
  {"https://www.google.com/?isvalue=true", "Google", false, ""},

  // Should fail:
  {"invalid-address", "Fake Title", true, "Nil Bookmark given"},
  {"https://invalid/address", "Invalid Title", true, "Nil Bookmark given"},
}

func TestBringUpDatabase(t *testing.T) {
  var err error 

  err = BringUpDatabase(":memory:")

  if err != nil {
    t.Fatal(err.Error())
  }

  err = Bookmarks.DB().Ping()

  if err != nil {
    t.Fatal(err.Error())
  }

  CloseDatabase()
}

func TestAddAndRetrieveRecord(t *testing.T) {
  var newentry *Bookmark 
  var err error 

  err = BringUpDatabase(":memory:")

  if err != nil {
    t.Fatal(err.Error())
  }

  for _, f := range fodder {
    newentry, _ = NewBookmark(f.Url, f.Comment)

    err = AddBookmark(newentry)

    if err == nil && f.ShouldFail {
      t.Fatalf("This test case should have failed: %s", string(f.Url))
    }

    if err != nil && f.ShouldFail {
      if f.Error != err.Error() {
        t.Fatalf("Wrong error message returned: %s", err.Error())
      }

      t.Skip()
    }

    entry, err := RetrieveBookmark(newentry.Url)

    if err != nil {
      t.Fatal(err.Error())
    }

    if entry == nil {
      t.Fatalf("Got Nil entry when I should have had: %s", string(newentry.Url))
    }

    if entry.Url != newentry.Url {
      t.Fatalf("Added entry and retrieved entry mismatch: %s versus %s", string(entry.Url), string(newentry.Url))
    }
  }

  // PurgeDatabase()
  CloseDatabase()
}

func TestAddAndRemoveRecord(t *testing.T) {
  var newentry *Bookmark 
  var err error 

  err = BringUpDatabase(":memory:")

  if err != nil {
    t.Fatal(err.Error())
  }

  for _, f := range fodder {
    newentry, _ = NewBookmark(f.Url, f.Comment)

    err = AddBookmark(newentry)

    if err == nil && f.ShouldFail {
      t.Fatalf("This test case should have failed: %s", string(f.Url))
    }

    if err != nil && f.ShouldFail {
      if f.Error != err.Error() {
        t.Fatalf("Wrong error message returned: %s versus %s", f.Error, err.Error())
      }

      t.Skip()
    }

    err = RemoveBookmark(newentry)

    if err != nil {
      t.Fatal(err.Error())
    }

    entry, _ := RetrieveBookmark(newentry.Url)

    if entry != nil {
      t.Fatalf("Able to retrieve deleted item: %s", newentry.Url)
    }
  }

  // PurgeDatabase()
  CloseDatabase()
}


