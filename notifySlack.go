
package main

import (
    "encoding/json"
    "flag"
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
)

var (
    IncomingUrl string = "https://hooks.slack.com/services/**********************************"
)

type Slack struct {
    Text       string `json:"text"`
    Username   string `json:"username"`
    Icon_emoji string `json:"icon_emoji"`
    Icon_url   string `json:"icon_url"`
    Channel    string `json:"channel"`
}

func main() {
    var (
        text     string
        username string
        emoji    string
        iconurl  string
        channel  string
        filename string
    )

    flag.StringVar(&text, "text", "", "text")
    flag.StringVar(&username, "username", "testbot", "username")
    flag.StringVar(&emoji, "icon_emoji", ":chicken:", "icon_emoji")
    flag.StringVar(&iconurl, "icon_url", "", "icon_url")
    flag.StringVar(&channel, "channel", "#general", "channel")
    flag.StringVar(&filename, "filename", "", "filename")

    flag.Parse()

    if len(text) == 0 && len(filename) > 0 {
        contents, err := ioutil.ReadFile(filename)
        if err != nil {
            fmt.Println(contents, err)
            return
        }
        text = string(contents)
    }

    params, _ := json.Marshal(Slack{
        string(text),
        string(username),
        string(emoji),
        string(iconurl),
        string(channel)})

    resp, _ := http.PostForm(
        IncomingUrl,
        url.Values{"payload": {string(params)}},
    )

    defer resp.Body.Close()

}
