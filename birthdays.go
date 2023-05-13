package main
import (
        "fmt"
        "time"
        "bytes"
        "text/template"
)
func birthdaysReminders(bdays map[string][]string) (string, bool, error){
        currentTime := time.Now()
        today := fmt.Sprintf("%d-%d", currentTime.Month(), currentTime.Day())
        names, present := bdays[today]
        
        message := ""
        
        if present {
                bdayTemplate := template.New("bday")
                bdayTemplate, err := bdayTemplate.Parse("Birthday(s) today: {{range .}}{{.}} {{end}}\n")
                if err != nil {
                    return message, present, err
                }
                var bdayMessage bytes.Buffer
                bdayTemplate.Execute(&bdayMessage, names)
                message = bdayMessage.String()
        }
        return message, present, nil
}

