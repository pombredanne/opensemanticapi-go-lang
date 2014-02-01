package main
import (
    "flag"
    "fmt"
    "log"
    "redis"
    "opensemanticapi/scraper"
)

func main() {
    // Parse command-line flags; needed to let flags used by Go-Redis be parsed.
    flag.Parse()

    spec := redis.DefaultSpec().Db(13).Password("")
    client, e := redis.NewSynchClientWithSpec(spec)
    if e != nil {
        log.Println("failed to create the client", e)
        return
    }

    key := "testkey"

    input := []byte("Testinput")

    client.Set(key, input)

    value, e := client.Get(key)
    if e != nil {
        log.Println("error on Get", e)
        return
    }

    fmt.Printf("Hey, ciao %s!\n", fmt.Sprintf("%s", value))

    // this is using the scraper section

    // catch a suggested list of results for a random keyword
    url := "http://en.wikipedia.org/w/api.php?action=opensearch&search=database&format=json&limit=3"
    val := scraper.FetchUrlContent(url)

    fmt.Printf("Hey, %s!\n", fmt.Sprintf("%s", val))
}
