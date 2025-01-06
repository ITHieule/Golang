CREATE DATABASE todo_app;
USE todo_app;
CREATE TABLE todos (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    completed BOOLEAN DEFAULT FALSE
);
INSERT INTO todos (title, completed) VALUES
('Học Golang', false),
('Viết API Todo', true),
('Kết nối MySQL', false);
