SET FOREIGN_KEY_CHECKS=0; -- to disable them

-- DROP TABLE IF EXISTS game;

-- CREATE TABLE game (
--     game_id INT AUTO_INCREMENT,
--     igdb_id INT,
--     name VARCHAR(250),
--     release_date INT,
--     description VARCHAR(2500),
--     url_name VARCHAR(250),
--     average_rating FLOAT DEFAULT 0,
--     PRIMARY KEY (game_id)
-- );

-- DROP TABLE IF EXISTS genre;

-- CREATE TABLE genre (
--     genre_id INT AUTO_INCREMENT,
--     igdb_id INT,
--     genre VARCHAR(60),
--     PRIMARY KEY (genre_id)
-- );

-- DROP TABLE IF EXISTS game_genre;

-- CREATE TABLE game_genre (
--     game_id INT,
--     genre_id INT,
--     FOREIGN KEY (game_id) REFERENCES game(game_id),
--     FOREIGN KEY (genre_id) REFERENCES genre(genre_id),
--     PRIMARY KEY (game_id,genre_id)
-- );

-- DROP TABLE IF EXISTS platform;

-- CREATE TABLE platform (
--     plat_id INT AUTO_INCREMENT,
--     igdb_id INT,
--     platform VARCHAR(60),
--     PRIMARY KEY (plat_id)
-- );

-- DROP TABLE IF EXISTS game_platform;

-- CREATE TABLE game_platform (
--     game_id INT,
--     plat_id INT,
--     FOREIGN KEY (game_id) REFERENCES game(game_id),
--     FOREIGN KEY (plat_id) REFERENCES platform(plat_id),
--     PRIMARY KEY (game_id,plat_id)
-- );

-- DROP TABLE IF EXISTS game_image;

-- CREATE TABLE game_image (
--     game_id INT,
--     image_path VARCHAR(1024),
--     FOREIGN KEY (game_id) REFERENCES game(game_id),
--     PRIMARY KEY (game_id)
-- );

DROP TABLE IF EXISTS user;

CREATE TABLE user (
    user_id INT AUTO_INCREMENT,
    email VARCHAR(256),
    username VARCHAR(20),
    description VARCHAR (2000) DEFAULT '',
    hash VARCHAR(255),
    PRIMARY KEY (user_id)
);

-- DROP TABLE IF EXISTS status;

-- CREATE TABLE status (
--     status_id INT AUTO_INCREMENT,
--     status_name VARCHAR(50),
--     PRIMARY KEY (status_id)
-- );

DROP TABLE IF EXISTS log;

CREATE TABLE log (
    log_id INT AUTO_INCREMENT,
    replay INT,
    plat_id INT,
    game_id INT,
    user_id INT,
    status_id INT,

    FOREIGN KEY (plat_id) REFERENCES platform(plat_id),
    FOREIGN KEY (game_id) REFERENCES game(game_id),
    FOREIGN KEY (user_id) REFERENCES user(user_id),
    FOREIGN KEY (status_id) REFERENCES status(status_id),
    PRIMARY KEY (log_id)
);


DROP TABLE IF EXISTS play_session;

CREATE TABLE play_session (
    session_id INT AUTO_INCREMENT,
    note VARCHAR(2000),
    session_date DATE,
    time_played TIME,
    started_game BOOLEAN,
    finished_game BOOLEAN,
    log_id INT,

    FOREIGN KEY (log_id) REFERENCES log(log_id),
    PRIMARY KEY (session_id)
);

DROP TABLE IF EXISTS review;

CREATE TABLE review (
    review_id INT AUTO_INCREMENT,
    score INT,
    favorite BOOLEAN,
    time_played_total TIME,
    review_text VARCHAR(8000),
    total_likes INT DEFAULT 0,
    total_comments INT DEFAULT 0,
    log_id INT,

    FOREIGN KEY (log_id) REFERENCES log(log_id),
    PRIMARY KEY (review_id)

);

DROP TABLE IF EXISTS list;

CREATE TABLE list (
    list_id INT AUTO_INCREMENT,
    name VARCHAR(250),
    description VARCHAR(2000),
    total_games INT,
    total_likes INT,
    total_comments INT,
    user_id INT,

    FOREIGN KEY (user_id) REFERENCES user(user_id),
    PRIMARY KEY (list_id)
);


DROP TABLE IF EXISTS game_list;

CREATE TABLE game_list (
    game_id INT,
    list_id INT,

    FOREIGN KEY (game_id) REFERENCES game(game_id),
    FOREIGN KEY (list_id) REFERENCES list(list_id),
    PRIMARY KEY (game_id,list_id)
 );


DROP TABLE IF EXISTS comment;

CREATE TABLE comment (
    comment_id INT AUTO_INCREMENT,
    text VARCHAR(2000),
    user_id INT,
    review_id INT DEFAULT 0,
    list_id INT DEFAULT 0,

    FOREIGN KEY (user_id) REFERENCES user(user_id),
    FOREIGN KEY (list_id) REFERENCES list(list_id),
    FOREIGN KEY (review_id) REFERENCES review(review_id),
    PRIMARY KEY (comment_id)
    
);




SET FOREIGN_KEY_CHECKS=1; -- to re-enable them
