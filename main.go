package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
    "os"
    "strconv"
    "github.com/joho/godotenv"
    "github.com/tidwall/gjson"
)

type Laptop struct {
    Id string
    Title string
    Location string
}

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

func get_json(url string) string {
    resp, err := http.Get(url)
    if err != nil {
        fmt.Println("Couldn't fetch results")
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Couldn't read results")
    }
    return string(body)
}

func getLaptops(json string) []Laptop {
    count, err := strconv.Atoi(gjson.Get(json, "findItemsAdvancedResponse.0.searchResult.0.@count").String())
    if err != nil {
        fmt.Println("Couldn't parse results")
    }

    laptops := []Laptop{}
    for i := 0; i <= count; i++ {
        root := fmt.Sprintf("findItemsAdvancedResponse.0.searchResult.0.item.%d.", i)
        title := gjson.Get(json, root + "title.0")
        itemId := gjson.Get(json, root + "itemId.0")
        location := gjson.Get(json, root + "location.0")
        laptop := Laptop{
            Id: itemId.String(),
            Title: title.String(),
            Location: location.String(),
        }
        laptops = append(laptops, laptop)
    }
    return laptops
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

    data := get_json(url)
    laptops := getLaptops(data)
    fmt.Println(laptops[1].Title)
}
