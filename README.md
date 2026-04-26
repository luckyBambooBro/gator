# Gator CLI Feeds Viewer
Welcome to Gator, our tool to view feeds via CLI

## Installation
You will need Postgres and Go to get started

***Postgres***

Linux: install via apt:  
`sudo apt install postresql postgresql-contrib`  
macOS: install with brew  
`brew install postgresql@15`

***Go***  
Linux: download with curl  
`curl -OL https://golang.org/dl/go1.26.2.linux-amd64.tar.gz`  
`sudo tar -C /usr/local -xzf go1.26.2.linux-amd64.tar.gz`  
macOS: download with brew  
`brew install go`

## Install Gator
In your command line install with go:  
`go install github.com/luckyBambooBro/gator@latest`

## Setup
In your CLI navigate to home and create a gator config file:  
`cd ~ && touch .gatorconfig.json`

Copy and paste the following contents into .gatorconfig.json, replacing \<connection string\> with your postgres connection string.

`{"db_url":"<connection string>","current_user_name":""}`

If you don't know your connection string click [here](https://github.com/luckyBambooBro/gator/blob/main/connection_string.md) 

## Usage
In your CLI run "gator" with any of the following commands:
- login (user): logs into a registered gator user
- register (new user): registers a new user
- reset: resets all data in gator 
- users: lists all users
- agg (time interval): fetches feeds according to time interval 
- addfeed (feedname, feedURL): adds feed into database and follows it for logged in user
- feeds: lists all feeds followed by user
- follow (feedURL): follows feedURL for current user 
- following: lists all feeds followed by user
- unfollow: unfollows feed for user
- browse (limit): prints information on feed(s) specified by limit