CREATE TABLE business_location
(
    ID           serial primary key,
    NAME         varchar(255),
    EMAIL        varchar(255),
    PHONE        varchar(255),
    ADDRESS      text,
    DESCRIPTION  varchar(255),
    IFRAME_LINK  text,
    CREATED_DATE timestamp NOT NULL DEFAULT now(),
    CREATED_BY   varchar(255),
    UPDATED_DATE timestamp,
    UPDATED_BY   varchar(255)
);

INSERT INTO business_location (name, phone, email, ADDRESS, description, iframe_link)
VALUES ('Cabang Taman Duren Sawit', '0813223344', 'cabang-1@gmail.com',
        'Jl. Taman No.12 Blok E6, Duren Sawit, Durensawit, East Jakarta City, Jakarta',
        'Ini adalah cabang taman duren sawit',
        'https%3A%2F%2Fwww.google.com%2Fmaps%2Fembed%3Fpb%3D!1m14!1m8!1m3!1d3966.162716210069!2d106.91244269311521!3d-6.242275082535274!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x2e698d01021dd817%253A0xa938efce571d975c!2sSupplier%2520Hotel%2520Resto%2520Kafe!5e0!3m2!1sid!2sid!4v1737750657426!5m2!1sid!2sid%22%2520width%3D%22600%22%2520height%3D%22450%22%2520style%3D%22border%3A0%3B%22%2520allowfullscreen%3D%22%22%2520loading%3D%22lazy%22%2520referrerpolicy%3D%22no-referrer-when-downgrade'),
       ('Cabang Pulo Mas',
        '0813223344', 'cabang-1@gmail.com',
        'Jl. Panca Wardi No.2, RT.2/RW.10, Kayu Putih, Kec. Pulo Gadung, Kota Jakarta Timur, Daerah Khusus Ibukota Jakarta 13210',
        'Ini adalah cabang pulo mas',
        'https%3A%2F%2Fwww.google.com%2Fmaps%2Fembed%3Fpb%3D!1m18!1m12!1m3!1d3966.6323026040245!2d106.88088169999999!3d-6.1799462!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x2e69f5002a9ac153%253A0xd6d0f772f4f2567c!2sMitra%2520Kirim!5e0!3m2!1sid!2sid!4v1737750884654!5m2!1sid!2sid%22%2520width%3D%22600%22%2520height%3D%22450%22%2520style%3D%22border%3A0%3B%22%2520allowfullscreen%3D%22%22%2520loading%3D%22lazy%22%2520referrerpolicy%3D%22no-referrer-when-downgrade');
