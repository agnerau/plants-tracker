CREATE TABLE plants (
                        id INT AUTO_INCREMENT PRIMARY KEY,
                        name VARCHAR(255) NOT NULL,
                        nickname VARCHAR(255),
                        bought_at DATE NULL,
                        planted_at DATE NULL,
                        died_at DATE NULL,
                        created_at DATETIME DEFAULT CURRENT_TIMESTAMP;
);

CREATE TABLE heights (
                         id INT AUTO_INCREMENT PRIMARY KEY,
                         value FLOAT NOT NULL,
                         created_at DDATETIME DEFAULT CURRENT_TIMESTAMP;,
                         plant_id INT NOT NULL,
                         FOREIGN KEY (plant_id) REFERENCES plants(id) ON DELETE CASCADE
);