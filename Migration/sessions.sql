CREATE TABLE sessions (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    session_index INT NOT NULL,       -- enumeration per user
    jit VARCHAR(512) NOT NULL,        -- current session token
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- Unique index to enumerate sessions per user
CREATE UNIQUE INDEX idx_user_session_index ON sessions(user_id, session_index);

-- Index jit for fast validation
CREATE UNIQUE INDEX idx_jit ON sessions(jit);
