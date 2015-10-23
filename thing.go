package main

import (
        "net/http"
        "fmt"
        "net"
)




func main() {
        addrs, err := net.LookupHost("foo.dev")
        if err != nil {
                fmt.Println(err.Error())
        }

        fmt.Printf("addrs: %+v\n", addrs)

        //req, err := http.NewRequest("GET", "https://gopro-dev.com/v1/oauth2/token/info", nil)
        req, err := http.NewRequest("GET", "http://foo.dev", nil)
        if err != nil {
                fmt.Printf("error: %s\n", err.Error())
        }

        client := &http.Client{}
        resp, err := client.Do(req)
        if err != nil {
                fmt.Printf("Error on Do: %s\n", err.Error())
        }

        fmt.Printf("Response: %+v\n", resp)
}
