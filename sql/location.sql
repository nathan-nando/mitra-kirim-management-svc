CREATE TABLE location
(
    ID           serial primary key,
    NAME         varchar(255),
    DESCRIPTION  varchar(255),
    IFRAME_LINK text,
    CREATED_DATE timestamp NOT NULL DEFAULT now(),
    CREATED_BY   varchar(255),
    UPDATED_DATE timestamp,
    UPDATED_BY   varchar(255)
);
