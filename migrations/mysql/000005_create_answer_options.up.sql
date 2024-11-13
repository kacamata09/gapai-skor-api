CREATE TABLE answer_options (
    id CHAR(36) DEFAULT (UUID()) PRIMARY KEY,
    question_id CHAR(36),
    content_answer TEXT,
    image_url VARCHAR(255),
    audio_url VARCHAR(255),
    is_correct BOOLEAN,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (question_id) REFERENCES questions(id) ON DELETE CASCADE
);
