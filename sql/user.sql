CREATE TABLE USERS
(
    ID           SERIAL primary key,
    NAME         varchar(255),
    EMAIL        varchar(255),
    TITLE        varchar(255),
    GENDER       varchar(6),
    PHONE        varchar(20),
    STATUS       int,
    IMG          varchar(255),
    CREATED_DATE timestamp,
    CREATED_BY   varchar(255),
    UPDATED_DATE timestamp,
    UPDATED_BY   varchar(255)
);

INSERT INTO USERS(name, email, title, gender, phone, status, img) VALUES( 'Admin MItra Kirim 2', 'admin@mkhoreca.co.id', 'Staff Operational', 'pria', '0813291212', 1, 'user.jpg');