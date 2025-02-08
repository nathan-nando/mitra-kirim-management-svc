drop table testimonial;

CREATE TABLE testimonial
(
    ID           serial primary key,
    NAME         varchar(50),
    DESCRIPTION  varchar(100),
    IMG          text,
    IS_CAROUSEL  int,
    CREATED_DATE timestamp NOT null DEFAULT now(),
    CREATED_BY   varchar(50),
    UPDATED_DATE timestamp,
    UPDATED_BY   varchar(50)
);
