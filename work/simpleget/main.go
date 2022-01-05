package main

import (
    "io/util"
    "log"
    "net/http"
)

func main() {
    resp, err := http.Get("http://localhost:18888")
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
    bodu, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }
    log.Println(string(body))
}
