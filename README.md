# Gator CLI Feeds Viewer
Welcome to Gator, our tool to view feeds via CLI

# Installation
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

# Setup
In your CLI navigate to home and create a gator config file:  
`cd ~ && touch .gatorconfig.json`

Copy and paste the following contents into .gatorconfig.json, replacing \<connection string\> with your postgres connection string.
If you don't know your connection string click [here](github.com/luckyBambooBro/gator/connection_string.md) 
