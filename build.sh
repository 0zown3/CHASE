#!/bin/bash

echo "Building CHASE!"
sudo docker build -t chase:latest .

echo "Starting CHASE container and exposing port 7000 on localhost!"
sudo docker run -d -p 7000:7000 chase:latest
