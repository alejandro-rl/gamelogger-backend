SET FOREIGN_KEY_CHECKS=0; -- to disable them

DROP TABLE IF EXISTS game;

CREATE TABLE game (
    game_id INT AUTO_INCREMENT,
    igdb_id INT,
    name VARCHAR(250),
    release_date INT,
    description VARCHAR(2500),
    url_name VARCHAR(250),
    average_rating FLOAT DEFAULT 0,
    PRIMARY KEY (game_id)
);

DROP TABLE IF EXISTS genre;

CREATE TABLE genre (
    genre_id INT AUTO_INCREMENT,
    igdb_id INT,
    genre VARCHAR(60),
    PRIMARY KEY (genre_id)
);

DROP TABLE IF EXISTS game_genre;

CREATE TABLE game_genre (
    game_id INT,
    genre_id INT,
    FOREIGN KEY (game_id) REFERENCES game(game_id),
    FOREIGN KEY (genre_id) REFERENCES genre(genre_id),
    PRIMARY KEY (game_id,genre_id)
);

DROP TABLE IF EXISTS platform;

CREATE TABLE platform (
    plat_id INT AUTO_INCREMENT,
    igdb_id INT,
    platform VARCHAR(60),
    PRIMARY KEY (plat_id)
);

DROP TABLE IF EXISTS game_platform;

CREATE TABLE game_platform (
    game_id INT,
    plat_id INT,
    FOREIGN KEY (game_id) REFERENCES game(game_id),
    FOREIGN KEY (plat_id) REFERENCES platform(plat_id),
    PRIMARY KEY (game_id,plat_id)
);

DROP TABLE IF EXISTS game_cover;

CREATE TABLE game_cover (
    game_id INT,
    type INT,
    image_path VARCHAR(1024),
    FOREIGN KEY (game_id) REFERENCES game(game_id),
    PRIMARY KEY (game_id,type)
);

SET FOREIGN_KEY_CHECKS=1; -- to re-enable them
