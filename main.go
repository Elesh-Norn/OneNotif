package main

import (
  "os"
  "fmt"
  "log"
  "time"
  "gopkg.in/yaml.v2"
  "path/filepath"
)

type Config struct {
        AccountSid string `yaml:"accountsid"`
        AuthToken string `yaml:"authtoken"`
        ToNumber string `yaml:"tonumber"`
        FromNumber string `yaml:"fromnumber"`
        Birthdays map[string][]string
}

func load_config() Config{
  configPath := filepath.Join(getExecutableDirPath(), "config.yaml")
  file, err := os.ReadFile(configPath)
  if err != nil {
    panic(err)
  }
  var c Config
  if err := yaml.Unmarshal(file, &c);
  err != nil {
    panic(err)
  }
  return c
}

func getExecutableDirPath() (string) {
  ex, err := os.Executable()
  if err != nil {
    panic(err)
  }
  return filepath.Dir(ex)
}

func getLogFile() (*os.File) {
  logPath := filepath.Join(getExecutableDirPath(), "OneNotif.log")
  file, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
  if err !=nil {
    log.Fatal(err)
  }
  return file
}

func main() {
        logFile := getLogFile()
        log.SetOutput(logFile)

        config := load_config()
        client := initTwilioClient(config.AccountSid, config.AuthToken)
        message := ""
        
        bdays, bdays_present, err := birthdays_reminder(config.Birthdays)
        if err != nil {
                log.Fatal(err)
        }
        
        if bdays_present {message += bdays}
        if message == "" {message = "Nothing today"}
 
        _, err = sendMessage(client, config.ToNumber, config.FromNumber, message)
        if err != nil {
                log.Fatal(err)
                return
        }
        log.Println("Sent message on %d. \n %d", time.Now(), message)
}
