package main

import (
  "fmt"
  "log"
  "net/http"
  "os"
  "io/ioutil"
  "encoding/json"

  "github.com/slack-go/slack"
  "github.com/slack-go/slack/slackevents"
  "github.com/kr/pretty"
)

var api = slack.New(getEnv("SLACK_BOT_TOKEN", ""))

func getEnv(envVar string, defaultTo string) string {
  var currentEnvVal string

  if (envVar == "") {
    fmt.Println("Please feed me an environment variable to look up.")
    os.Exit(1)
  }

  val, ok := os.LookupEnv(envVar)
  if (!ok) {
    if (defaultTo != "") {
      currentEnvVal = defaultTo
    } else {
      fmt.Println("Could not read environment variable", envVar)
      os.Exit(1)
    }
  } else {
    currentEnvVal = val
  }

  return currentEnvVal
}

func main() {
  var port = getEnv("PORT", "3000")

  http.HandleFunc("/", MessageHandler)

  fmt.Println("Abraxas is listening on port", port)
  log.Fatal(http.ListenAndServe(":" + port, nil))
}

func MessageHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Println("inbound slack message")

  var signingSecret = getEnv("SLACK_SIGNING_SECRET", "")

  fmt.Fprintf(w, "yes this is the event endpoint\n")

  body, err := ioutil.ReadAll(r.Body)
  if (err != nil) {
    w.WriteHeader(http.StatusBadRequest)
    return
  }

  sv, err := slack.NewSecretsVerifier(r.Header, signingSecret)
  if (err != nil) {
    w.WriteHeader(http.StatusBadRequest)
    return
  }

  if _, err := sv.Write(body); err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    return
  }

  if err := sv.Ensure(); err != nil {
    w.WriteHeader(http.StatusUnauthorized)
    return
  }

  eventsAPIEvent, err := slackevents.ParseEvent(json.RawMessage(body), slackevents.OptionNoVerifyToken())
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    return
  }

  if eventsAPIEvent.Type == slackevents.URLVerification {
    var r *slackevents.ChallengeResponse
    err := json.Unmarshal([]byte(body), &r)
    if err != nil {
      w.WriteHeader(http.StatusInternalServerError)
      return
    }
    w.Header().Set("Content-Type", "text")
    w.Write([]byte(r.Challenge))
  }

  if eventsAPIEvent.Type == slackevents.CallbackEvent {
    fmt.Println("got callback event")

    innerEvent := eventsAPIEvent.InnerEvent
    fmt.Printf("%# v \n", pretty.Formatter(innerEvent))


    switch ev := innerEvent.Data.(type) {
      case *slackevents.AppMentionEvent:
        fmt.Println("thank you for mentioning me")

        eventData := innerEvent.Data.(*slackevents.AppMentionEvent)
        fmt.Printf("%# v \n", pretty.Formatter(eventData.Channel))

        timestamp := fmt.Sprintf("%# v", eventData.EventTimeStamp)
        msgRef := slack.NewRefToMessage(eventData.Channel, timestamp)

        if err = api.AddReaction("cry", msgRef); err != nil {
          fmt.Println("Dag, there was an error reacting to the message.")
          fmt.Println(err)
        }
      default:
        fmt.Println("Hm, no handler found for this message")
    }
  }
}
