drop table testimonial;

CREATE TABLE testimonial
(
    ID           serial primary key ,
    NAME         varchar(255),
    DESCRIPTION  varchar(255),
    IMG          text,
    IS_CAROUSEL   int2,
    CREATED_DATE timestamp,
    CREATED_BY   varchar(255),
    UPDATED_DATE timestamp,
    UPDATED_BY   varchar(255)
);
