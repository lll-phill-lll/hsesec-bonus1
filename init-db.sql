CREATE TABLE users(
                      id INTEGER PRIMARY KEY,
                      login VARCHAR (255) UNIQUE NOT NULL,
                      money_amount INTEGER NOT NULL,
                      card_number VARCHAR (255) UNIQUE NOT NULL,
                      status BOOLEAN NOT NULL
);
CREATE TABLE pswd(
                   id INTEGER PRIMARY KEY,
                   password VARCHAR (255) UNIQUE NOT NULL
);
INSERT INTO users VALUES (1, 'anton', 10, '76969', TRUE);
INSERT INTO users VALUES (2, 'petr', 20, '9898998', TRUE);
INSERT INTO users VALUES (3, 'misha', 39, '777777', TRUE);
INSERT INTO users VALUES (4, 'kolya', 40, '66666', TRUE);
INSERT INTO users VALUES (5, 'matvey', 76, '878787878', FALSE);
INSERT INTO users VALUES (6, 'sasha', 343, '999999', FALSE);
INSERT INTO users VALUES (7, 'vasya', 767, '7777', TRUE);
INSERT INTO users VALUES (8, 'ini', 99, '8786776t6', FALSE);
INSERT INTO users VALUES (9, 'python', 777, '65656', TRUE);
INSERT INTO users VALUES (10050, 'c++', 000, '779988778', FALSE);

INSERT INTO pswd VALUES (1, 'aaaa');
INSERT INTO pswd VALUES (2, 'nnnjn');
INSERT INTO pswd VALUES (3, 'mom');
INSERT INTO pswd VALUES (4, 'dad');
INSERT INTO pswd VALUES (5, 'granny');
INSERT INTO pswd VALUES (6, 'nina viktorovna');
INSERT INTO pswd VALUES (7, 'petya');
INSERT INTO pswd VALUES (8, 'dog');
INSERT INTO pswd VALUES (9, 'cat');
INSERT INTO pswd VALUES (10050, 'tratata');
