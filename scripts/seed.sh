#!/bin/bash

curl -v \
     -X POST \
     -H "Content-Type: application/json" \
     -d '{"url": "https://nathanberry.co.uk/index.xml"}' \
     localhost:8080/rss_feed

curl -v \
    -X POST \
    -H "Content-Type: application/json" \
    -d '{"url": "https://aws.amazon.com/blogs/aws/feed/"}' \
    localhost:8080/rss_feed

curl -v \
    -X POST \
    -H "Content-Type: application/json" \
    -d '{"url": "https://feeds.bbci.co.uk/news/technology/rss.xml"}' \
    localhost:8080/rss_feed
