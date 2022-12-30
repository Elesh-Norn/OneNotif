package main

import (
  "github.com/twilio/twilio-go"
  api "github.com/twilio/twilio-go/rest/api/v2010"
)


func initTwilioClient(accountSid string, authToken string) *twilio.RestClient { 
        client := twilio.NewRestClientWithParams(twilio.ClientParams{
                Username: accountSid,
                Password: authToken})
        return client
}

func sendMessage(client *twilio.RestClient, toNumber string, fromNumber string, message string) (*api.ApiV2010Message, error) {
        params := &api.CreateMessageParams{}
        params.SetBody(message)
        params.SetTo(toNumber)
        params.SetFrom(fromNumber)

        resp, err := client.Api.CreateMessage(params)
        return resp, err
}

