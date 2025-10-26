CREATE DATABASE IF NOT EXISTS myappdb;
USE myappdb;

-- Users table
CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO users (name, email)
VALUES 
    ('Jash', 'jash@example.com'),
    ('Jane', 'jane@example.com');

-- Spaces table (AUTO_INCREMENT replaces sequence)
CREATE TABLE IF NOT EXISTS spaces (
    space_id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    owner VARCHAR(30) NOT NULL
);

-- Messages table
CREATE TABLE IF NOT EXISTS messages (
    msg_id INT AUTO_INCREMENT PRIMARY KEY,
    space_id INT NOT NULL,
    author VARCHAR(30) NOT NULL,
    msg_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    msg_text VARCHAR(1024) NOT NULL,
    FOREIGN KEY (space_id) REFERENCES spaces(space_id)
);

-- Indexes
CREATE INDEX msg_timestamp_idx ON messages(msg_time);

-- Sample data
INSERT INTO spaces (name, owner)
VALUES 
    ('General Chat', 'Jash'),
    ('Dev Room', 'Jane');

INSERT INTO messages (space_id, author, msg_text)
VALUES
    (1, 'Jash', 'Hello everyone!'),
    (1, 'Jane', 'Hi Jash, good to see you here.'),
    (2, 'Jane', 'Deploying the new version now.'),
    (2, 'Jash', 'Awesome, ping me once it’s live.');
