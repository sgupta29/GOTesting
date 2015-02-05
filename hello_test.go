package hello

import (
    "fmt"
    "net/http"
    "testing"
)

func init() {
    http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, world12!")
    testmodle();
}

func testmodle(t *testing.T) {
        fmt.Fprint("TestSuccess")
	}