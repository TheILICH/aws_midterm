#!/bin/bash

sudo su
apt update -y && apt upgrade -y
apt install git

sudo snap install go --classic

go get .
go run .


