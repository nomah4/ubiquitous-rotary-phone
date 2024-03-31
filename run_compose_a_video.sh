#!/usr/bin/bash

function runTests() {
    # main tests
sudo docker-compose up;
# video dl test
./messengers/discord/run.sh;
./messengers/vk/run.sh;
./messengers/telegram/run.sh;
./messengers/youtube/run.sh;
}

time runTests;
# without video dl (containers pre builded
# real    0m7,728s
#user    0m0,006s
#sys     0m0,006s

# with video dl
#real    0m20,516s
#user    0m4,386s
#sys     0m0,691s