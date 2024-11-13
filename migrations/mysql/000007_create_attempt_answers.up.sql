CREATE TABLE attempt_answers (
    id CHAR(36) DEFAULT (UUID()) PRIMARY KEY,
    attempt_id CHAR(36),
    question_id CHAR(36),
    selected_answer_option_id CHAR(36),
    score INT,
    attempt_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (attempt_id) REFERENCES attempts(id) ON DELETE CASCADE,
    FOREIGN KEY (question_id) REFERENCES questions(id) ON DELETE CASCADE,
    FOREIGN KEY (selected_answer_option_id) REFERENCES answer_options(id) ON DELETE CASCADE
);