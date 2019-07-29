/*
Cache file server
storing file in memory when they're first requested and serve responses using Serve-Content rather than a file server

Has a data structure to hold the content in memory (time and content are important)
first try to get the file from in memory cache, if not found add to cache, use lock to prevent racce conditions modifying the cache


*/

package main 

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

// data structure to store a file in memory 
type cacheFile struct {
	content io.ReadSeeker
	modTime time.Time 
}

var cache map[string]*cacheFile // map to store files in memory
var mutex = new(sync.RWMutex) // mutext to handle race conditions while handling parallel cache changes

func main() {
	cache = make(map[string]*cacheFile) // makes the map usable
	http.HandleFunc("/", serveFiles)
	http.ListenAndServe(":8080", nil)
}

func serveFiles(w http.ResponseWriter, r *http.Request) {
	mutex.RLock()
	v, found := cache[r.URL.Path] // loads from the cache if already populated
	mutex.RUnlock()

	if !found { // when file isn't in the cache, starts loading process
		mutex.Lock()
		defer mutex.Unlock()
		fileName := "./files" + r.URL.Path
		f, err := os.Open(fileName) // opens the file to cache, making sure to defer the close
		defer f.Close()

		if err != nil {
			http.NotFound(w, r)
			return
		}

		var b bytes.Buffer // copy the file into in-memory buffer
		_, err = io.Copy(&b, f)
		if err != nil {
			http.NotFound(w, r) // handle errors copying from file to memory
			return 
		}

		reader := bytes.NewReader(b.Bytes()) // puts the bytes into a reader for later use
		info, _ := f.Stat()
		v := &cacheFile{
			content: reader,
			modTime: info.ModTime(),
		}
		cache[r.URL.Path] = v // populates the cache objects and stores for later
	}

	http.ServeContent(w, r, r.URL.Path, v.modTime, v.content) // serves the file from cache

}
