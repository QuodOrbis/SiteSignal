package config

import (
     "fmt"
     //"log"
     "os"
     "encoding/json"

)

type Config struct {
    ProxyUse bool `json:"ProxyUse"`
    ProxyURL string `json:"ProxyURL"`
    ProxyUser string `json:"ProxyUser"`
    ProxyPass string `json:"ProxyPass"`
}

func LoadConfiguration(file string) Config {
    var config Config
    configFile, err := os.Open(file)
    defer configFile.Close()
    if err != nil {
        fmt.Println("config file env.json error: ",err.Error())
        os.Exit(1)
    }
    jsonParser := json.NewDecoder(configFile)
    err = jsonParser.Decode(&config)
    if err != nil {
        fmt.Println("config file env.json error: ",err.Error())
        os.Exit(1)
    }
    return config
}
