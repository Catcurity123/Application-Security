# Use an official OpenJDK 11 as the base image
FROM openjdk:11-jdk-slim

# Set environment variables for Maven
ENV MAVEN_VERSION=3.5.0
# Dir to install Maven
ENV MAVEN_HOME=/opt/maven 

# Install required packages: curl, git, and Maven
RUN apt update && apt-get install -y \
    curl \
    git \
    apt-utils \
    jq \
    bash \
# Above are for installing needed packages
    && curl -fsSL https://archive.apache.org/dist/maven/maven-3/${MAVEN_VERSION}/binaries/apache-maven-${MAVEN_VERSION}-bin.tar.gz \
    | tar -xz -C /opt \
    && mv /opt/apache-maven-${MAVEN_VERSION} ${MAVEN_HOME} \
    # Creat link
    && ln -s /opt/maven/bin/mvn /usr/bin/mvn \
    # Remove cache and delete downloading package
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*


# Set default working directory
WORKDIR /app

# Print versions for verification
RUN java -version && mvn -version && git --version && curl --version

# Default command to keep the container running
CMD ["bash"] 