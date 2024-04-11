package main;

import (
    "flag"
    "fmt"
    "net/http"
)


func main() {
    directory := flag.String("dir", "./", "the directory that the static data is in")
    port := flag.Int("port", 3000, "the port the server will listen on")

    flag.Parse()

    hostString := fmt.Sprintf(":%d", *port);

    http.Handle("/", http.FileServer(http.Dir(*directory)))
    http.ListenAndServe(hostString, nil)
}
