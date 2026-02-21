FROM golang:1.23

# Install tools + MySQL client
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
        git vim nano jq mariadb-client && \
    rm -rf /var/lib/apt/lists/*

# Working directory inside container
WORKDIR /GoCode

# Expose app port
EXPOSE 4000

# Default command (run Go directly)
CMD ["go", "run", "main.go"]

# docker exec -it ${container_id} bash