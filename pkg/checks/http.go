package checks

import (
     "fmt"
     "log"
     "time"
     "github.com/gocolly/colly"
     "encoding/base64"
     "net/http"
     "net/url"
     "github.com/QuodOrbis/SiteSignal/pkg/config"

)

func MakeHTTPRequest(configx config.Config, urlx string, ch chan<-string) {

  start := time.Now()
  //resp, _ := http.Get(url)
  // Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		//colly.AllowedDomains("hackerspaces.org", "wiki.hackerspaces.org"),
	)

  // if we should use a proxy..
  if configx.ProxyUse && configx.ProxyURL != "" {

    proxyURL, err := url.Parse(configx.ProxyURL)
  	if err != nil {
  		log.Println(err)
  	}
    log.Println("using proxy: ",configx.ProxyURL)
    c.WithTransport(&http.Transport{
        Proxy: http.ProxyURL(proxyURL),
    })
  }

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		log.Printf("Link found: %q -> %s\n", e.Text, link)
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		c.Visit(e.Request.AbsoluteURL(link))
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting ", r.URL.String())

    if configx.ProxyUse && configx.ProxyUser != "" && configx.ProxyPass != "" {
      log.Println("Using Proxy auth, with username: ",configx.ProxyUser)
      // add proxy auth
      auth := configx.ProxyUser+":"+configx.ProxyPass
    	basicAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
      r.Headers.Set("Proxy-Authorization", basicAuth)
    }
	})

	// Start scraping..
	c.Visit(urlx)

  secs := time.Since(start).Seconds()
  ch <- fmt.Sprintf("%.2f elapsed against URL: %s", secs, urlx)
}
