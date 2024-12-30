CREATE TABLE Users (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE
);

CREATE TABLE Groups (
    group_id SERIAL PRIMARY KEY,
    group_name VARCHAR(50) NOT NULL
);

CREATE TABLE Messages (
    message_id SERIAL PRIMARY KEY,
    sender_id INT NOT NULL REFERENCES Users(user_id),
    receiver_id INT REFERENCES Users(user_id),
    group_id INT REFERENCES Groups(group_id),
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
