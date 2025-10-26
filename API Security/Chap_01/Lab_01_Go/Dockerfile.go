FROM golang:1.23

# Install tools + MySQL client
RUN apt-get update && \
    apt-get install -y git vim nano jq mariadb-client && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Copy Go source files
COPY . .



EXPOSE 8080
CMD ["bash"]
