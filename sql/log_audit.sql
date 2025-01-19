CREATE TABLE LOG_AUDIT
(
    ID   int primary key,
    NAME   varchar(40),
    SUBJECT   varchar(40),
    IP   varchar(20),
    ACTION   varchar(1),
    OLD_DATA   text,
    NEW_DATA   text,
    DATE timestamp
);
