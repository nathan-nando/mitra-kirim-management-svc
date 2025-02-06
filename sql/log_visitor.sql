DROP TABLE LOG_VISITOR;

CREATE TABLE LOG_VISITOR
(
    ID   serial primary key,
    IP   varchar(20) not null ,
    DATE timestamp not null
);
