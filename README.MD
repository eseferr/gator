This is a guided gator project from Boot.dev

This is an RSS feed aggregator called "Gator" in go it's a CLI tool that allows users to:

    Add RSS feeds from across the internet to be collected
    Store the collected posts in a PostgreSQL database
    Follow and unfollow RSS feeds that other users have added
    View summaries of the aggregated posts in the terminal, with a link to the full post

RSS feeds are a way for websites to publish updates to their content. You can use this project to keep up with your favorite blogs, news sites, podcasts, and more!

Project Description:
In this guided project you'll practice building a CLI in Go, and you'll use production-ready database tools like PostgreSQL, SQLc, Goose, and psql. This won't just be another CLI utility, but a service that has a long-running service worker that reaches out over the internet to fetch data from remote locations.

## Prerequisites

- Go 1.16 or higher
- PostgreSQL 12 or higher

This program can be installed by go install github.com/eseferr/gator

## Configuration

Create a config.json file in your home directory:

````json
## Configuration

Create a `config.json` file in your home directory with the following structure:

```json
{
  "host": "localhost",
  "port": 5432,
  "database": "your_database_name",
  "user": "your_username",
  "password": "your_password"
  // Add any other required fields
}
```
Commands:
gator start
  - Starts the RSS feed monitoring service that:
  - Reads the configured RSS feeds from the database
  - Periodically fetches new posts from these feeds
  - Stores new posts in the PostgreSQL database
  - Continues running until stopped
gator status - Check current status
gator help - Show available commands

````
