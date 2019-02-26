use cyfir;

DROP TABLE IF EXISTS person;
create table person (
    id INT NOT NULL,
    fname VARCHAR(20),
    lname VARCHAR(20),
    sex CHAR(1),
    birth DATE,
    PRIMARY KEY (id)
);

INSERT INTO PERSON VALUES
    (1, 'jane', 'smith', 'F', '1990-06-21'),
    (2, 'sue', 'wyatt', 'F', '1998-07-22'),
    (3, 'ken', 'doll', 'M', '1997-08-23'),
    (4, 'barbie', 'doll', 'F', '1996-09-24'),
    (5, 'john', 'doe', 'M', '1995-10-25');


DROP TABLE IF EXISTS account;
create table account  (
    id INT NOT NULL,
    account_type VARCHAR(10),
    balance DECIMAL,
    person_id INT NOT NULL,
    FOREIGN KEY (person_id)
        REFERENCES person(id)
        ON DELETE CASCADE
);

INSERT INTO ACCOUNT VALUES
    (1001, 'checking', 1200, 1),
    (1002, 'checking', 1200, 2),
    (1003, 'checking', 1200, 3),
    (1004, 'checking', 1200, 4),
    (1005, 'checking', 1200, 5);