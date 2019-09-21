package main

import (
    "fmt"
    "github.com/joho/godotenv"
    "io/ioutil"
    "net/http"
    "net/url"
    "os"
)

func generate_search_url(global_id string, app_name string, keywords string) (string, error) {
    var u *url.URL
    u, err := url.Parse("https://svcs.ebay.com/services/search/FindingService/v1")
    if err != nil {
        return "", err
    }
    params := url.Values{}
    params.Add("OPERATION-NAME", "findItemsByKeywords")
    params.Add("SERVICE-VERSION", "1.0.0")
    params.Add("SECURITY-APPNAME", app_name)
    params.Add("GLOBAL-ID", global_id)
    params.Add("RESPONSE-DATA-FORMAT", "JSON")
    params.Add("REST-PAYLOAD", "")
    params.Add("keywords", keywords)
    params.Add("paginationInput.entriesPerPage", "30")
    u.RawQuery = params.Encode()
    return u.String(), err
}

func main() {
    err := godotenv.Load()
    if err != nil {
        fmt.Println("Error loading .env file")
    }
    global_id := os.Getenv("GLOBAL_ID")
    app_name := os.Getenv("APP_NAME")
    keyword := "thinkpad t440p"
    url, err := generate_search_url(global_id, app_name, keyword)
    if err != nil {
        return
    }

    resp, err := http.Get(url)
    if err != nil {
        fmt.Printf("%s", err)
    } else {
        data, _ := ioutil.ReadAll(resp.Body)
        fmt.Println(string(data))
    }
}
