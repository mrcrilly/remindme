
package main 

import (
  "testing"
  "net/http"
  "net/http/httptest"
  "io/ioutil"
  "encoding/json"
  "reflect"
)

func checkErrorT(err error, t *testing.T) {
  if err != nil {
    t.Fatal(err.Error())
  }
}

func buildAndRunGet(handler func(http.ResponseWriter, *http.Request)) (data []byte, err error) {
  var testsrv *httptest.Server
  var httprsp *http.Response
  var httpdata []byte

  testsrv = httptest.NewServer(http.HandlerFunc(handler))

  httprsp, err = http.Get(testsrv.URL)
  checkError(err)

  httpdata, err  = ioutil.ReadAll(httprsp.Body)
  checkError(err)

  httprsp.Body.Close()
  testsrv.Close()

  return httpdata, err 
}

func buildAndRunPost(handler func(http.ResponseWriter, *http.Request), topost interface{}) (data []byte, err error) {
  return nil, nil 
}

func TestGetHandler(t *testing.T) {
  var err error 
  var foundit bool

  var getrsp GetResponse
  var httpdata []byte 
  var rsptype reflect.Type

  httpdata, err = buildAndRunGet(GetHandler)
  checkErrorT(err, t)

  err = json.Unmarshal(httpdata, &getrsp)
  checkErrorT(err, t)

  rsptype = reflect.TypeOf(getrsp)

  _, foundit = rsptype.FieldByName("Unread")
  if ! foundit {
    t.Fatal("No Unread field in response")
  }

  _, foundit = rsptype.FieldByName("Read")
  if ! foundit {
    t.Fatal("No Read field in response")
  }

}

func TestPostHandler(t *testing.T) {
  var err error 
  var foundit bool

  var postrsp PostResponse
  var httpdata []byte 
  var rsptype reflect.Type

  httpdata, err = buildAndRunGet(PostHandler)
  checkErrorT(err, t)

  err = json.Unmarshal(httpdata, &postrsp)
  checkErrorT(err, t)

  rsptype = reflect.TypeOf(postrsp)

  _, foundit = rsptype.FieldByName("Status")
  if ! foundit {
    t.Fatal("No Status field in response")
  }
}

func TestDeleteHandler(t *testing.T) {
  var err error 
  var foundit bool

  var deletersp DeleteResponse
  var httpdata []byte 
  var rsptype reflect.Type

  httpdata, err = buildAndRunGet(DeleteHandler)
  checkErrorT(err, t)

  err = json.Unmarshal(httpdata, &deletersp)
  checkErrorT(err, t)

  rsptype = reflect.TypeOf(deletersp)

  _, foundit = rsptype.FieldByName("Status")
  if ! foundit {
    t.Fatal("No Status field in response")
  }
}


