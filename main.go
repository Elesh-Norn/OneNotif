package main

import (
  "os"
  "bytes"
  "log"
  "time"
  "gopkg.in/yaml.v2"
  "path/filepath"
  "net/http"
  "encoding/json"
)

type Config struct {
  DiscordEnabled bool `yaml:"DiscordEnabled"`
  DiscordWebhook string `yaml:"DiscordWebhook"`
  TwilioEnabled bool `yaml:"TwilioEnabled"`
  TwilioAccountSid string `yaml:"TwilioAccountSid"`
  TwilioAuthToken string `yaml:"TwilioAuthToken"`
  TwilioToNumber string `yaml:"TwilioToNumber"`
  TwilioFromNumber string `yaml:"TwilioFromNumber"`
  Birthdays map[string][]string `yaml:"Birthdays"`
  City string `yaml:"City"`
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
        message := ""
        
        bdays, bdays_present, err := birthdaysReminders(config.Birthdays)
        if err != nil {log.Fatal(err)}
        weather, err := getWeatherData(config.City)
        if err != nil {log.Fatal(err)}
        
        // TODO Ok for now, but template next time
        if bdays_present {message += bdays}
        if weather != "" {message += weather}
        if message == "" {message = "Nothing today"}

        if config.TwilioEnabled {
          client := initTwilioClient(config.TwilioAccountSid, config.TwilioAuthToken)
          _, err = sendMessage(client, config.TwilioToNumber, config.TwilioFromNumber, message)
          if err != nil {
           log.Fatal(err)
           return
         }
       }
        if config.DiscordEnabled {
          json_body, _ := json.Marshal(map[string]string{
            "username": "OneNotif",
            "content": message,
          })
          body := bytes.NewBuffer(json_body)
          _, err := http.Post(config.DiscordWebhook, "application/json", body)
          if err != nil {
            log.Fatal(err)
            return
          }
        }
       log.Println("Sent message on %d. \n %d", time.Now(), message)
}
