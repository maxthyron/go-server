package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func router(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.Error(w, "404 Not Found", http.StatusNotFound)
        return
    }

    switch r.Method {
    case "GET":
        http.ServeFile(w, r, "static/html/index.html")
    case "POST":
        if err := r.ParseForm(); err != nil {
            fmt.Fprintf(w, "ParseForm() err %v", err)
            return
        }
        fmt.Fprintf(w, "Post from: %v", r.PostForm)
    case "OPTIONS":
        fmt.Fprintf(w, "GET and POST methods available")

    default:
        http.Error(w, "418 I'm teapot", http.StatusTeapot)
    }
}

func getPort() string {
    p := os.Getenv("PORT")
    if p != "" {
        return ":" + p
    }

    return ":8080"
}

func server() {
    http.HandleFunc("/", router)
    fs := http.FileServer(http.Dir("static/"))
    fmt.Printf("Start server @8080\n")
    http.Handle("/static/", http.StripPrefix("/static/", fs))
    if err := http.ListenAndServe(getPort(), nil); err != nil {
        log.Fatal(err)
    }
}

func main() {
    server()
}
