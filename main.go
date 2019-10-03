package main

import (
     "fmt"
     "os"
     "time"
     "github.com/QuodOrbis/SiteSignal/pkg/config"
     "github.com/QuodOrbis/SiteSignal/pkg/checks"

)

var Version = "0.0.1"

func main() {

  fmt.Println(`
    _________.__  __           _________.__                     .__
   /   _____/|__|/  |_  ____  /   _____/|__| ____   ____ _____  |  |
   \_____  \ |  \   __\/ __ \ \_____  \ |  |/ ___\ /    \\__  \ |  |
   /        \|  ||  | \  ___/ /        \|  / /_/  >   |  \/ __ \|  |__
  /_______  /|__||__|  \___  >_______  /|__\___  /|___|  (____  /____/
          \/               \/        \/   /_____/      \/     \/      `)

  fmt.Println("> Version: "+Version)

  configx := config.LoadConfiguration("env.json")

  start := time.Now()
  ch := make(chan string)
  for _,url := range os.Args[1:]{
      go checks.MakeHTTPRequest(configx, url, ch)
  }

  for range os.Args[1:]{
    fmt.Println(<-ch)
  }
  fmt.Printf("Total run time: %.2fs elapsed\n", time.Since(start).Seconds())
}
