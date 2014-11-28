# RemindMe
I often find my self opening a lot of tabs to URLs I later forget to read. This is usually in the morning after opening up [Hacker News](https://news.ycombinator.com/). I figured if I had somewhere to stash the URL until later, I could come back and run through a list, reading the content whenever I have time, on any device, regardless of browser being used or syncing features available.

Bookmarks in browsers are, in my opinion, clunky. I want to pop a URL in a box, put in an optional comment, and save it until later. This will then be a list I can read on any device.

## Authentication
This is most certainly something that will be implemented in the HTTP system (see below).

## CLI
There is a simple CLI tool to demonstrate/test the library's APIs. Feel free to knock your self out with it.

## HTTP
Soon an HTTP server will be developed that one can deploy anywhere. This server will serve up a simple (Bootstrap) UI that you can use to add, view and delete URLs from database.

I don't intend on supporting HTTPS directly from the application as I feel NginX off-loading your SSL certificate is the better, safer option.

## Database
SQLite3 is the default, with the current, very simple, implementation of [Gorm](https://github.com/jinzhu/gorm) not allowing any other selection yet. This will change if I feel the need or there is deman.

## Sharing
I intend on adding a feature which allows people to share their lists in a read-only fashion via the web UI or perhaps an RSS feed.

## Likes
Maybe a simple "thumbs up" feature so people can show approval of a particular URL in a shared list?

## Encryption
Maybe add encryption into the mix, allowing one to keep a more secure database on their servers. Could be useful, but given the links are public in the first place, might be a waste of time and resources.
