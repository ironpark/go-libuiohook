#!/bin/zsh

LIBUIOHOOK_VERSION=1.2
# Update libuiohook source code tar
# https://github.com/kwhat/libuiohook
rm -rf libuiohook && mkdir libuiohook
curl -L https://github.com/kwhat/libuiohook/tarball/$LIBUIOHOOK_VERSION -o libuiohook.tar.gz
tar -xvf libuiohook.tar.gz -C ./libuiohook --strip-components=1
rm -rf libuiohook.tar.gz libuiohook/.github libuiohook/demo libuiohook/test libuiohook/man libuiohook/pc