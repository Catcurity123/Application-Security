#!/bin/bash

set -e  # Exit on error

# Update package lists
sudo apt update -y

# Install dependencies
sudo apt install -y \
    docker.io \
    git \
    curl \
    unzip \
    jq

# Start and enable Docker
sudo systemctl enable --now docker

# Install AWS CLI v2
curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
unzip awscliv2.zip
sudo ./aws/install
rm -rf aws awscliv2.zip

# Verify installations
docker --version
git --version
aws --version
curl --version