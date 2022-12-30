package main
import (
        "fmt"
        "time"
        "bytes"
        "text/template"
)
func birthdays_reminder(bdays map[string][]string) (string, bool, error){
        currentTime := time.Now()
        today := fmt.Sprintf("%d-%d", currentTime.Month(), currentTime.Day())
        names, present := bdays[today]
        
        message := ""
        
        if present {
                bday_template := template.New("bday")
                bday_template, err := bday_template.Parse("Birthday(s) today: {{range .}}{{.}} {{end}}\n")
                if err != nil {
                    return message, present, err
                }
                var bday_message bytes.Buffer
                bday_template.Execute(&bday_message, names)
                message = bday_message.String()
        }
        return message, present, nil
}

