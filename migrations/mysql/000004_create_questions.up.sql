CREATE TABLE questions (
    id CHAR(36) DEFAULT (UUID()) PRIMARY KEY,
    test_id CHAR(36),
    content_question TEXT,
    image_url VARCHAR(255),
    audio_url VARCHAR(255),
    question_type VARCHAR(255),
    points INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (test_id) REFERENCES tests(id) ON DELETE CASCADE
);
