#!/usr/bin/env bash

ssh-keygen -t rsa -b 4096 -m PEM -f ./keys/app.key
openssl rsa -in ./keys/app.key -pubout -outform PEM -out ./keys/app.key.pub.pem
