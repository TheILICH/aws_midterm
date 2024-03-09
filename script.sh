#!/bin/bash

sudo su
apt update -y && apt upgrade -y
apt install git
git clone https://github.com/TheILICH/aws_midterm.git

sudo snap install go --classic

go get .
go run .


