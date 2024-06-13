#!/bin/bash

# keys worflow
openssl genpkey -algorithm RSA -out ./backend/private_key.pem
openssl rsa -pubout -in ./backend/private_key.pem -out ./backend/public_key.pem

# keys local
openssl genpkey -algorithm RSA -out private_key.pem
openssl rsa -pubout -in private_key.pem -out public_key.pem
