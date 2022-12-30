# One Notification

This project is a service that send you reminders or information only once a day. 

In my quest of trying to reduce my time spent on screen and my reliance on third party apps that become bloated/hacked or hostile to the user.
This will send me a sms once a day with stuff I want or need to check regularly such as; birthdays, appointements, meteo, stocks.

## Configuring it
This app requires a Twilio account. So much to not rely on third party but it's the fastest I could get something setup. 

Everything is in the yaml.config file. Rename or copy the example to config.yaml.
- birthdays: must be the birthdays list. The keys are mm/dd and the values a list of strings.
- AuthToken : the auth token that you will use. Don't commit this.
- AccountSid: Account Sid of the Twilio account you will send sms from
- FromNumber: "Twilio Phone Number message are sent from"
- ToNumber: "Phone Number this app will send message to"

### Installation and usage 

Clone this repository then `go build`.
Create a `config.yaml` at the root of the project and fill the values as indicated above. `config_example.yaml` will give you an idea of what to do.
Just execute main once you have built it.
