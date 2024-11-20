package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type user struct {
    ID int `json:"id"`
    Username string `json:"username"`
    Password string `json:"password"`
    Email string `json:"email"`
    Token string `json:"token"`
}

func getUserByID(id int) (user, error) {
    addr := fmt.Sprintf("%s/users/%d", endpoint, id)
    res, err := http.Get(addr)
    if err != nil {
        return user{}, err
    }
    defer res.Body.Close()
    body, err := io.ReadAll(res.Body)
    if err != nil {
        return user{}, err
    }
    var u user
    err = json.Unmarshal(body, &u)
    if err != nil {
        return user{}, err
    }
    return u, nil
}

func login(username, password string) (user, error) {
    addr := fmt.Sprintf("%s/auth/login", endpoint)
    payload, err := json.Marshal(user{
        Username: username,        
        Password: password,
    })
    if err != nil {
        return user{}, err
    }
    res, err := http.Post(addr, "application/json", bytes.NewBuffer(payload))
    if err != nil {
        return user{}, err
    }
    defer res.Body.Close()
    body, err := io.ReadAll(res.Body)
    if err != nil {
        return user{}, err
    }
    var u user
    err = json.Unmarshal(body, &u)
    if err != nil {
        return user{}, err
    }
    return u, nil
}

func getCurrentUser(token string) (user, error) {
    client := http.Client{
        Timeout: time.Second,
    }
    addr := fmt.Sprintf("%s/auth/me", endpoint)
    req, err := http.NewRequest(http.MethodGet, addr, nil)
    if err != nil {
        return user{}, err
    }
    req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
    res, err := client.Do(req)
    if err != nil {
        return user{}, err
    }
    defer res.Body.Close()
    body, err := io.ReadAll(res.Body)
    if err != nil {
        return user{}, err
    }
    var u user
    err = json.Unmarshal(body, &u)
    if err != nil {
        return user{}, err
    }
    return u, nil
}
