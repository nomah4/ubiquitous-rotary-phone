#!/usr/bin/bash

echo "TEST video dl"
yt-dlp "https://www.youtube.com/watch?v=jNQXAC9IVRw" --force-overwrites
echo "YOUTUBE download status: $?"

echo "TEST get thumbnail"
# TODO hardcoded for demo
curl "https://yt3.ggpht.com/fxGKYucJAVme-Yz4fsdCroCFCrANWqw0ql4GYuvx8Uq4l_euNJHgE-w9MTkLQA805vWCi-kE0g=s88-c-k-c0x00ffffff-no-rj-mo" -m 1
echo "YOUTUBE GET thumbnail status: $?"
# TODO traceroute