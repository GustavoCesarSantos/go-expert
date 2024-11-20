package main

import (
	"flag"
	"fmt"
	"log"
)

const endpoint = "https://dummyjson.com"

func main() {
    action := flag.String("action", "", "which action to perform")
    id := flag.Int("id", 0, "user id")
    username := flag.String("username", "", "user name")
    password := flag.String("password", "", "user password")
    flag.Parse()
    switch *action {
    case "get-user":
        u, err := getUserByID(*id)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println(u)
    case "login":
        u, err := login(*username, *password)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println(u)
    case "get-current-user":
        u, err := login(*username, *password)
        if err != nil {
            log.Fatal(err)
        }
        me, err := getCurrentUser(u.Token)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println(me)
    }
}
