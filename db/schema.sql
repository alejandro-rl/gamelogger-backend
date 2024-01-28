CREATE TABLE games (
    id INT AUTO_INCREMENT,
    igdb_id INT,
    name VARCHAR(250),
    release_date INT,
    description VARCHAR(2000),
    average_rating FLOAT,
    PRIMARY KEY (id)
);