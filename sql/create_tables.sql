CREATE TABLE IF NOT EXISTS series (
                                           id INT PRIMARY KEY NOT NULL,
                                           title VARCHAR (200),
                                           description VARCHAR (1000)
                                           );

CREATE TABLE IF NOT EXISTS comic_books (
                                           id INT PRIMARY KEY NOT NULL,
                                           title VARCHAR (200),
                                           description VARCHAR (1000),
                                           series_id INT NOT NULL,

    CONSTRAINT FK_series_id_comic_books FOREIGN KEY (series_id)
    REFERENCES series(id)  
    ON DELETE CASCADE
    ON UPDATE CASCADE
);      

CREATE TABLE IF NOT EXISTS characters (
                                           id INT PRIMARY KEY NOT NULL,
                                           name VARCHAR (100),
                                           description VARCHAR (1000)
                                           );                           

CREATE TABLE IF NOT EXISTS characters_comic_books (
                                           character_id INT NOT NULL,
                                           comic_book_id INT NOT NULL,
                                           character_name VARCHAR (100),
                                           comic_book_title VARCHAR (200),
                                           PRIMARY KEY (character_id, comic_book_id),
                                           

    CONSTRAINT FK_character_comic_book FOREIGN KEY (character_id)
    REFERENCES characters(id)  
    ON DELETE CASCADE
    ON UPDATE CASCADE,

    CONSTRAINT FK_comic_book_character FOREIGN KEY (comic_book_id)
    REFERENCES characters(id)  
    ON DELETE CASCADE
    ON UPDATE CASCADE
);
CREATE TABLE IF NOT EXISTS characters_series (
                                           character_id INT NOT NULL,
                                           series_id INT NOT NULL,
                                           character_name VARCHAR (100),
                                           series_title VARCHAR (200),
                                           PRIMARY KEY (character_id, series_id),
                                             

    CONSTRAINT FK_character_series FOREIGN KEY (character_id)
    REFERENCES characters(id)  
    ON DELETE CASCADE
    ON UPDATE CASCADE,

    CONSTRAINT FK_series_character FOREIGN KEY (series_id)
    REFERENCES characters(id)  
    ON DELETE CASCADE
    ON UPDATE CASCADE
);