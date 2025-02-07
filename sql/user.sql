DROP TABLE USERS;

CREATE TABLE USERS
(
    ID           varchar(40),
    NAME         varchar(80),
    USERNAME     varchar(50) UNIQUE,
    PASSWORD         TEXT,
    EMAIL        varchar(100),
    TITLE        varchar(50),
    GENDER       varchar(1),
    PHONE        varchar(20),
    STATUS       int,
    IMG          varchar(80),
    CREATED_DATE timestamp,
    CREATED_BY   varchar(80),
    UPDATED_DATE timestamp,
    UPDATED_BY   varchar(80)
);

INSERT INTO USERS(name, email, title, gender, phone, status, img)
VALUES ('Admin MItra Kirim 2', 'admin@mkhoreca.co.id', 'Staff Operational', 'L', '0813291212', 1, 'user.jpg');