package cli

import (
  "github.com/codegangsta/cli"

    "net/http"
    // "io/ioutil"
    "fmt"
)

func login (c *cli.Context) {
    println("echo login,", c.Args().Get(0), c.Args().Get(1))


    // client := &http.Client{}
    reqest, _ := http.NewRequest("GET", "https://passport.baidu.com", nil)
     
    reqest.Header.Set("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
    reqest.Header.Set("Accept-Encoding","gzip, deflate, sdch")
    reqest.Header.Set("Accept-Language","en-US,en;q=0.8,zh-CN;q=0.6,zh-TW;q=0.4")
    reqest.Header.Set("Connection","Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.93 Safari/537.36")
     


    transport := http.Transport{}
    response, err := transport.RoundTrip(reqest)
    if err != nil {
        // t.Fatal(err)
    }
    // // Check if you received the status codes you expect. There may
    // // status codes other than 200 which are acceptable.
    // if resp.StatusCode != 200 && resp.StatusCode != 302 {
    //     t.Fatal("Failed with status", resp.Status)
    // }


    // response,_ := client.Do(reqest)
    if response.StatusCode == 302 {
        // body, _ := ioutil.ReadAll(response.Body)
        // bodystr := string(body);
        fmt.Println("cookie: ", response.Cookies())
    } else {
        fmt.Println("error, code=", response.StatusCode)
    }

    println("login done")
}