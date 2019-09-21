package main

import (
    "fmt"
    "github.com/joho/godotenv"
    "io/ioutil"
    "net/http"
    "net/url"
    "os"
)

func generate_search_url(global_id string, app_name string, keywords string, currency string) (string, error) {
    var u *url.URL
    u, err := url.Parse("https://svcs.ebay.com/services/search/FindingService/v1")
    if err != nil {
        return "", err
    }
    params := url.Values{}
    params.Add("OPERATION-NAME", "findItemsAdvanced")
    params.Add("SERVICE-VERSION", "1.0.0")
    params.Add("SECURITY-APPNAME", app_name)
    params.Add("GLOBAL-ID", global_id)
    params.Add("RESPONSE-DATA-FORMAT", "JSON")
    params.Add("REST-PAYLOAD", "")
    params.Add("keywords", keywords)
    params.Add("categoryId", "177")
    params.Add("categoryId", "175672")
    params.Add("paginationInput.entriesPerPage", "30")
	params.Add("itemFilter.name", "Condition")
	params.Add("itemFilter.value", "Used")
    params.Add("itemFilter.name", "MinPrice")
    params.Add("itemFilter.value", "0")
    params.Add("itemFilter.paramName", "Currency")
    params.Add("itemFilter.paramValue", currency)
    params.Add("itemFilter.name", "MaxPrice")
    params.Add("itemFilter.value", "120")
    params.Add("itemFilter.paramName", "Currency")
    params.Add("itemFilter.paramValue", currency)

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
    currency := os.Getenv("CURRENCY")
    keyword := "thinkpad t440p"
    url, err := generate_search_url(global_id, app_name, keyword, currency)
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
