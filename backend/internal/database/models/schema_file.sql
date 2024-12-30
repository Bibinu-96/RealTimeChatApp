-- Create enum for message types
CREATE TYPE message_type AS ENUM ('text', 'file');

-- Create Users table
CREATE TABLE Users (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE
);

-- Create Groups table
CREATE TABLE Groups (
    group_id SERIAL PRIMARY KEY,
    group_name VARCHAR(50) NOT NULL
);

-- Create Files table
CREATE TABLE Files (
    file_id SERIAL PRIMARY KEY,
    filename VARCHAR(255) NOT NULL,
    file_size INT NOT NULL CHECK (file_size <= 10485760), -- 10MB in bytes
    file_type VARCHAR(100) NOT NULL,
    file_path VARCHAR(500) NOT NULL,
    uploaded_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    uploaded_by INT NOT NULL REFERENCES Users(user_id)
);

-- Create Messages table with file support
CREATE TABLE Messages (
    message_id SERIAL PRIMARY KEY,
    sender_id INT NOT NULL REFERENCES Users(user_id),
    receiver_id INT REFERENCES Users(user_id),
    group_id INT REFERENCES Groups(group_id),
    message_type message_type NOT NULL DEFAULT 'text',
    content TEXT,
    file_id INT REFERENCES Files(file_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    -- Ensure either content or file_id is present based on message type
    CONSTRAINT valid_message_content CHECK (
        (message_type = 'text' AND content IS NOT NULL AND file_id IS NULL) OR
        (message_type = 'file' AND content IS NULL AND file_id IS NOT NULL)
    ),
    -- Ensure message has either receiver_id or group_id, not both
    CONSTRAINT valid_message_target CHECK (
        (receiver_id IS NOT NULL AND group_id IS NULL) OR
        (receiver_id IS NULL AND group_id IS NOT NULL)
    )
);

-- Create indexes for better performance
CREATE INDEX idx_messages_sender ON Messages(sender_id);
CREATE INDEX idx_messages_receiver ON Messages(receiver_id);
CREATE INDEX idx_messages_group ON Messages(group_id);
CREATE INDEX idx_messages_file ON Messages(file_id);
CREATE INDEX idx_files_uploader ON Files(uploaded_by);
CREATE INDEX idx_files_upload_date ON Files(uploaded_at);

-- Create function for file size validation
CREATE FUNCTION validate_file_size()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.file_size > 10485760 THEN
        RAISE EXCEPTION 'File size exceeds maximum limit of 10MB';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create trigger for file size validation
CREATE TRIGGER validate_file_size_trigger
    BEFORE INSERT OR UPDATE ON Files
    FOR EACH ROW
    EXECUTE FUNCTION validate_file_size();