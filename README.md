# transfer_tweet_to_slack

A simple slack bot which picks up tweets posted on twitter.    

This bot picks up every tweet which has specific words, and post them on a slack channel. This bot is purely written in Golang.

## Build

After running `go get github.com/kosen14s/tanasfer_tweet_to_slack`,
you need to run the following command.
```
go build -o ttts
```
The "ttts" means "Transfer Tweet To Slack." (I don't like this one :|)

## Usage

Just run the following command.
```
./ttts [consumer key] [consumer secret] [access token] [access token secret] [slack api token] [words to be picked up] [slack channel]
```

#### Arguments

###### consumer key
The consumer key of twitter client you have.

###### consumer secret
The consumer secret of twitter client you have.

###### access token
The access token you got on twitter.

###### access token secret
The access token secret you got on twitter.

###### slack api token
The slack API token you have.

###### words to be picked up
The words you want to pick up from twitter.
You can specify several words by dividing words with commas.  
Example: #kosen14s, #muscle

###### slack channel
The slack channel which this bot will post on.  
Example: #random

## Dependency

- [nlopes/slack](https://github.com/nlopes/slack)
- [ChimeraCoder/anaconda](https://github.com/ChimeraCoder/anaconda)

