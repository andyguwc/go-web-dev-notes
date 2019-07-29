/*
Large file uploaded and timeout occurs

Resume the download from a file. Retry the download


*/



package main

import (
	// "io"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	// local file to store the download
	file, err := os.Create("ff.dmg")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// download remote file to local file 
	// retry up to 100 times 
	location := "https://download-installer.cdn.mozilla.net/pub/firefox/releases/40.0.3/mac/en-US/Firefox%2040.0.3.dmg"
	err = download(location, file, 100)
	if err != nil {
		fmt.Println(err)
		return
	}

	// displaying the size of the file 
	fi, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Got it with %v bytes downloaded", fi.Size())
}

// download with retries 
func download(location string, file *os.File, retries int64) error {
	// create new get request for the file being downloaded
	req, err := http.NewRequest("GET", location, nil)
	if err != nil {
		return err
	}
	fi, err := file.Stat()
	if err != nil {
		return err
	}

	// retrieves size of the current file 
	current := fi.Size()
	if current > 0 {
		start := strconv.FormatInt(current, 10)
		// when local file already has content, sets a header requesting where the local file left off
		req.Header.Set("Range", "bytes="+start+"-")
	}

	cc := &http.Client{Timeout: 5 * time.Minute}
	res, err := cc.Do(req)

	// when checking for an error, tries the request again if the error was caused by timeout
	if err != nil && hasTimedOut(err) {
		if retries > 0 {
			return download(location, file, retries-1)
		}
		return err
	} else if err != nil {
		return err
	}

	// handles nonsuccessful HTTP statu codes 
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		errFmt := "Unsuccess HTTP request. Status: %s"
		return fmt.Errorf(errFmt, res.Status)
	}

	// If the server doesn’t support serving partial files, sets retries to 0
	if res.Header.Get("Accept-Ranges") != "bytes" {
		retries = 0
	}

	// if a timeout error occcurs while copying , tries retrieving the remaining content 
	_, err = io.Copy(file, res.Body)
	if err != nil && hasTimedOut(err) {
		if retries > 0 {
			return download(location, file, retries-1)
		}
		return err
	} else if err != nil {
		return err
	}

	return nil
}

func hasTimedOut(err error) bool {
	switch err := err.(type) {
	case *url.Error:
		if err, ok := err.Err.(net.Error); ok && err.Timeout() {
			return true
		}
	case net.Error:
		if err.Timeout() {
			return true
		}
	case *net.OpError:
		if err.Timeout() {
			return true
		}
	}

	errTxt := "use of closed network connection"
	if err != nil && strings.Contains(err.Error(), errTxt) {
		return true
	}

	return false
}