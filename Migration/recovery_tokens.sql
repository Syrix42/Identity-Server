CREATE TABLE recovery_tokens (
    id INT AUTO_INCREMENT PRIMARY KEY,
    jit VARCHAR(512) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP NOT NULL
);

-- Index for fast lookup by jit
CREATE UNIQUE INDEX idx_jit ON recovery_tokens(jit);

-- Optional: index by expires_at if you frequently query expired tokens
CREATE INDEX idx_expires_at ON recovery_tokens(expires_at);
