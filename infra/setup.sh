#!/bin/bash

"""
Script meant to show what dependencies to use with Server
"""

sudo apt update -y
sudo apt upgrade -y

sudo apt install nginx golang-go python3 git make wget curl -y
