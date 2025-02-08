DROP TABLE suggestion;

CREATE TABLE suggestion(
    ID varchar(40) primary key ,
    NAME varchar(80),
    EMAIL varchar(80),
    MESSAGE text,
    REPLY text,
    HAS_REPLIED int,
    CREATED_DATE timestamp,
    CREATED_BY varchar(80),
    UPDATED_DATE timestamp,
    UPDATED_BY varchar(80)
);
