CREATE DATABASE projects (
    id BIGINT PRIMARY AUTO_INCREMENT,
    name VARCHAR(100) REQUIRE,
    category VARCHAR(255) REQUIRE,
    tracked BIGINT DEFAULT 0,
    progress INT DEFAULT 0,
    access String
);