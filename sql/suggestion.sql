
CREATE TABLE suggestion(
    ID varchar(40) primary key ,
    NAME varchar(255),
    EMAIL varchar(255),
    MESSAGE text,
    HAS_REPLIED int,
    CREATED_DATE timestamp,
    CREATED_BY varchar(255),
    UPDATED_DATE timestamp,
    UPDATED_BY varchar(255)
);
