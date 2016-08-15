# transfer_tweet_to_slack

This is slack bot that transfer tweets to slack channel, writtin in golang.

## build

after ` go get github.com/kosen14s/tanasfer_tweet_to_slack` 
```
go build -o ttts
```

The "ttts" is shorted form of Transfer_Tweet_To_Slack. I don't like this name :|

## usage

```
./ttts <Your Twitter Client Consumer Key> <Your Twitter Client Consumer Secret> <Your Twitter Access Token> <Your Twitter Access Token Secret> <Your Slack Api Token> <The Words Search from Twitter (Comma Separated), like "#kosen14s, #muscle"> <Channel Name of Transfering to, like "random">
```

## Dependency

- [nlopes/slack](https://github.com/nlopes/slack)
- [ChimeraCoder/anaconda](https://github.com/ChimeraCoder/anaconda)

