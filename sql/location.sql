CREATE TABLE location
(
    ID           int primary key,
    NAME         varchar(255),
    DESCRIPTION  varchar(255),
    LATITUDE     decimal(8, 6),
    LONGITUDES   decimal(9, 6),
    SORT         int,
    CREATED_DATE timestamp,
    CREATED_BY   varchar(255),
    UPDATED_DATE timestamp,
    UPDATED_BY   varchar(255)
);
