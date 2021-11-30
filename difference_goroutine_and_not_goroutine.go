// package main

// import (
//   "fmt"
//   "net/http"
//   "time"
// )
// func main() {
//   start := time.Now()
//   sitelist := []string{
//     "https://www.google.com//",
//     "https://www.duckduckgo.com/",
//     "https://www.developer.com/",
//     "https://www.codeguru.com/",
//     "https://www.nasa.gov/",
//     "https://golang.com/",
//   }
//   for _, site := range sitelist {
//     GetSiteStatus(site)
//   }
//   fmt.Printf("\n\nTime elapsed since %v\n\n", time.Since(start))

// }
// func GetSiteStatus(site string) {
//   if _, err := http.Get(site); err != nil {
//     fmt.Printf("%s is down\n", site)
//   } else {
//     fmt.Printf("%s is up\n", site)
//   }
// }



package main
import (
  "fmt"
  "net/http"
  "sync"
  "time"
)
func main() {
  var wg sync.WaitGroup
  start := time.Now()
  sitelist := []string{
    "https://www.google.com//",
    "https://www.duckduckgo.com/",
    "https://www.developer.com/",
    "https://www.codeguru.com/",
    "https://www.nasa.gov/",
    "https://golang.com/",
  }
  for _, site := range sitelist {
    go GetSiteStatus(site, &wg)
    wg.Add(1)
  }
  wg.Wait()
  fmt.Printf("\n\nTime elapsed since %v\n\n", time.Since(start))
}
func GetSiteStatus(site string, wg *sync.WaitGroup) {
  defer wg.Done()
  if _, err := http.Get(site); err != nil {
    fmt.Println("%s is down\n", site)

  } else {
    fmt.Printf("%s is up\n", site)
  }
}