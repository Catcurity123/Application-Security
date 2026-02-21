# Use an existing Python 3 image as the base
FROM python:3

# Install Ansible
RUN pip install ansible

# Optionally, you can install additional dependencies or configure your environment here
RUN apt-get update && \
    apt-get install -y \
    iputils-ping \
    net-tools \
    nano \
    curl \
    wget \
    gnupg \
    putty \
    && rm -rf /var/lib/apt/lists/*
    
# Create home_folder directory and move to it
WORKDIR /home_folder

# Copy essentials files
COPY ansible.cfg hosts-dev /home_folder/

# Specify the entry point
CMD [ "ansible-playbook", "--version" ]
