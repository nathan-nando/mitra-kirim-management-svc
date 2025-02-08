DROP TABLE configuration;

CREATE TABLE configuration
(
    ID           serial primary key,
    KEY          varchar(50) UNIQUE,
    TYPE         varchar(50),
    VALUE        text,
    DESCRIPTION  varchar(255),
    FORM_TYPE    varchar(100),
    CREATED_DATE timestamp NOT NULL DEFAULT now(),
    CREATED_BY   varchar(100),
    UPDATED_DATE timestamp,
    UPDATED_BY   varchar(100)
);

--APPLICATION_CONFIG
INSERT INTO configuration(key, type, value, form_type)
values ('appName', 'APPLICATION_CONFIG', 'PT. Mitra Kirim Horeca', 'text'),
       ('appDescription', 'APPLICATION_CONFIG', 'Menjamin Pasokan Hotel Restoran dan
Kafe dengan Bahan Berkualitas', 'text'),
       ('appLogo', 'APPLICATION_CONFIG', 'logo.png', 'file');

-- SOCIAL_MEDIA_CONFIG
INSERT INTO configuration(key, type, value, form_type)
values ('facebook', 'SOCIAL_MEDIA_CONFIG', 'https://www.facebook.com/profile.php?id=100078752286650', 'text'),
       ('instagram', 'SOCIAL_MEDIA_CONFIG', 'https://www.instagram.com/mitra_kirim?igsh=emt4bmtydnVlcThw', 'text'),
       ('tiktok', 'SOCIAL_MEDIA_CONFIG', 'https://www.tiktok.com/@mitra_kirim', 'text'),
       ('whatsapp', 'SOCIAL_MEDIA_CONFIG', '+6281510085433', 'text');

-- ONLINE_SHOP_CONFIG
INSERT INTO configuration(key, type, value, form_type)
values ('tokopedia', 'ONLINE_SHOP_CONFIG', 'https://www.tokopedia.com/mitrakirim', 'text'),
       ('shopee', 'ONLINE_SHOP_CONFIG', 'https://shopee.co.id/mitrakirim', 'text'),
       ('blibli', 'ONLINE_SHOP_CONFIG', 'https://blibli.onelink.me/GNtk/odrsvzso', 'text'),
       ('bukalapak', 'ONLINE_SHOP_CONFIG', 'https://www.bukalapak.com/u/mitra_kirim_793707', 'text'),
       ('lazada', 'ONLINE_SHOP_CONFIG', 'https://www.lazada.co.id/shop/mitra-kirim/?path=index.htm', 'text');


-- LAYOUT_CONFIG
INSERT INTO configuration(key, type, value, form_type)
values ('heroImg', 'LAYOUT_CONFIG', 'hero.jpg', 'text'),
       ('heroDesc', 'LAYOUT_CONFIG',
        'PT Mitra Kirim Horeca adalah perusahaan yang menyediakan berbagai kebutuhan untuk hotel, restoran, dan kafe. Kami menawarkan produk berkualitas tinggi, mulai dari peralatan dapur, perlengkapan meja, hingga bahan makanan dan minuman, yang dirancang untuk mendukung operasional bisnis hospitality Anda. \n\nDengan pelayanan yang cepat dan profesional, PT Mitra Kirim Horeca menjadi mitra terpercaya dalam memenuhi kebutuhan sektor horeca (hotel, restoran, dan kafe) dengan harga kompetitif dan kualitas terbaik.',
        'textarea'),
       ('services', 'LAYOUT_CONFIG', '[
  {
    "title": "Penyediaan Bahan Berkualitas",
    "description": "Respon cepat terhadap kebutuhan mendesak, keluhan, serta ketersediaan bahan berkualitas",
    "img": "service-1.jpg"
  },
  {
    "title": "Pengiriman Barang Secara Cepat",
    "description": "Pengiriman tepat waktu memastikan waktu operasional bisnis anda tidak terganggu",
    "img": "service-2.jpg"
  },
  {
    "title": "Layanan Konsultasi HORECA",
    "description": "Solusi operasional bisnis maupun pemilihan produk. Membantu menemukan barang yang tepat sesuai kebutuhan dan saran profesional",
    "img": "service-3.jpg"
  }
]', '');
