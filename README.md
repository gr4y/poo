poo [![Build Status](https://travis-ci.org/gr4y/poo.svg)](https://travis-ci.org/gr4y/poo)
---


When I was building an RSS reader I discovered how broken RSS Feeds are and that inspired me to build my own RSS parsing library in Go. I thought about a name long and hard and came up with *poo*, which reflects the way website owners care about their feeds meeting the standard.

## There are still things to do:

* add some working unit tests

## Done

* unmarshalling pubDate into time.Time instead of string- (Works with RFC 1132 and RFC 1132 Z)
* add support for iTunes Podcast Feeds-

If you want to contribute, feel free to fork, hack and submit a PR.
