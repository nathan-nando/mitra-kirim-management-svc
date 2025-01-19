CREATE TABLE configuration
(
    ID           int primary key,
    KEY          varchar(255),
    TYPE         varchar(255),
    VALUE        text,
    DESCRIPTION  varchar(255),
    SORT         int,
    CREATED_DATE timestamp,
    CREATED_BY   varchar(255),
    UPDATED_DATE timestamp,
    UPDATED_BY   varchar(255)
);
