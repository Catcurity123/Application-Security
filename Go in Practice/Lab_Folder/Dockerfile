FROM golang:1.23

# Install useful tools
RUN apt-get update && apt-get install -y git vim nano jq && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Use volumes to avoid copying on every change
VOLUME ["/app"]

# Default command: open interactive shell
CMD ["bash"]
