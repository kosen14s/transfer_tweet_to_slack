package main

import (
	"fmt"
	"os"
    "net/url"
	"github.com/ChimeraCoder/anaconda"
	"github.com/nlopes/slack"
)

func  PostMessageToSlack (slackApi *slack.Client, channel slack.Channel, tweetText string) error {
    /// post message
    params := slack.NewPostMessageParameters()
    params.Username = "Twitter#Kosen14s"
    _, _, err := slackApi.PostMessage(channel.ID, tweetText, params) /// PostMessage function is changed now on github, to rewrite here when package updated
    if err != nil {
        return err
    }
    return nil
}

func main() {
    if len(os.Args) < 8 {
        fmt.Fprintln(os.Stderr, "There are 7 arguments required. There is ConsumerKey, ConsumerSecret, UserAccessToken, UserAccessTokenSecret, SlackApiTOken, twitter search word (, spearated), SlackChannelToPostMessage")
        return
    }


    /// TODO: Handling Login Error
    anaconda.SetConsumerKey(os.Args[1])
    anaconda.SetConsumerSecret(os.Args[2])

    api := anaconda.NewTwitterApi(os.Args[3], os.Args[4])
    // api.SetLogger(anaconda.BasicLogger)


    /// Slack Api
    channelName := os.Args[7]
    slackApi := slack.New(os.Args[5])

    /// get channel list
    channels, err := slackApi.GetChannels(false)
    if err != nil {
        fmt.Fprintln(os.Stderr, "The Slack Api is invalid.")
        return
    }

    /// get channel to post message
    var channel slack.Channel
    for _, c := range channels {
        if c.Name == channelName {
            channel = c
            break
        }
    }
    if channel.ID == "" {
        fmt.Fprintln(os.Stderr, "Channel " + channelName + " is not exist.")
        return
    }

    /// start streaming
    v := url.Values{}
    v.Set("track", os.Args[6])
    stream := api.PublicStreamFilter(v)

    for {
        t := <-stream.C
        switch tweet := t.(type) {
        case anaconda.Tweet:
            if tweet.RetweetedStatus == nil {
                go PostMessageToSlack(slackApi, channel, "https://twitter.com/" + tweet.User.IdStr + "/status/" + tweet.IdStr);
            }
        }
    }
}
