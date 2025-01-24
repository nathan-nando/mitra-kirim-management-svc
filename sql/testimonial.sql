CREATE TABLE testimonial
(
    ID           int primary key,
    NAME         varchar(255),
    DESCRIPTION  varchar(255),
    IMG          text,
    SORT  int,
    CREATED_DATE timestamp,
    CREATED_BY   varchar(255),
    UPDATED_DATE timestamp,
    UPDATED_BY   varchar(255)
);
