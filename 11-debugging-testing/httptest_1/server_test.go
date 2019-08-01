/* Test Http request and response
Check the httptest_2 for more comprehensive testing 
*/



package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandleGet(t *testing.T) {
	// create a multiplexer to run tests on 
	mux := http.NewServeMux()
	mux.HandleFunc("/post/", handleRequest)

	// capture returned HTTP response 
	writer := httptest.NewRecorder()

	// create request to handler you want to test 
	request, _ := http.NewRequest("GET", "/post/1", nil)

	// send request to tested handler
	mux.ServeHTTP(writer, request)

	// checks response recorder for results 
	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
	var post Post
	json.Unmarshal(writer.Body.Bytes(), &post)
	if post.Id != 1 {
		t.Errorf("Cannot retrieve JSON post")
	}
}

func TestHandlePut(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/post/", handleRequest)

	writer := httptest.NewRecorder()
	json := strings.NewReader(`{"content":"Updated post","author":"Sau Sheong"}`)
	request, _ := http.NewRequest("PUT", "/post/1", json)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
}
