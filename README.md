# Gator
Gator is a Go-based application for managing and browsing RSS feeds. I built this application as part of a boot.dev tutorial. It allows users to subscribe to feeds as well as fetch and read posts.

## Features
* Add and Manage RSS feeds
* Fetch posts
* User management: Seperate users with seperate feed subscriptions
* Post browsing: View posts with optional limits on the number of posts displayed

## Installation
To properly use Gator, you need to have Postgres and Go installed.
You can install Postgres on Mac using Brew like this:
```bash
brew install postgresql@16
```
And on Linux/WSL (Debian) like this:
```bash
sudo apt update
sudo apt install postgresql postgresql-contrib
```
To confirm that your installation worked, you can check that you're on version 16+ of Postgres:
```bash
psql --version
```

You can install Go [here](https://go.dev/doc/install)

### Installing Gator

```bash
go install github.com/samersawan/gator@latest
```

## Set Up
Start by setting up a config file in your home directory called ```gatorconfig.json```. You can do this using
```bash
touch ~/.gatorconfig.json
```
Once that's done, make sure you set up your config file using this format
```bash
{
  "db_url":"postgres://username@localhost:5432/gator?sslmode=disable",
  "current_user_name":"username"
}
```
where username is your own username on your computer.

## Usage
Using Gator is simple. Here are the commands you can use:
* Register: ```gator register <username>```. Register allows you to register a new user, and updates the config file to be using the newly registered user.
* Login: ```gator login <username>```. Logs you in as the specified user.
* Reset: ```gator reset```. Resets the database.
* AddFeed: ```gator addfeed <url>```. Creates a new feed and sets current user as the creator of the feed.
* Feeds: ```gator feeds```. Displays all the feeds registered to the database.
* Follow: ```gator follow <url>```. Allows the current user to follow a feed that already exists.
* Unfollow: ```gator unfollow <url>```. Allows the current user to unfollow a feed that they currently follow.
* Following: ```gator following```. Displays all the feeds that the current user follows.
* Agg: ```gator agg <time_between_reqs>```. Fetches posts from the feeds that are followed by the current user. Time between requests specifies how frequently it checks for new posts.
* Browse: ```gator browse <limit>```. Allows you to browse posts saved for the current user. Can specifiy an optional limit. Default limit is 2 posts. 

