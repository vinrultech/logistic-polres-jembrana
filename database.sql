CREATE TABLE "log_users" (
  "id" serial NOT NULL PRIMARY KEY,
  "nama" character varying(250) NOT NULL,
  "username" character varying(20) NOT NULL,
  "password" character varying(100) NOT NULL,
  "role" character varying(10) NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE "log_users"
ADD "hp" character varying NULL,
ADD "foto" character varying NULL;

INSERT INTO log_users 
    (nama, username, password, role) 
    VALUES ('Superuser', 'superuser', 
     '$2a$10$/f1QhCtKF2uFo94yZm1cdeAJ2tO.Mfd.K8LWLcPhOdr/Jh1t7TiAa', 
      'superuser'
    );

CREATE TABLE "log_kategoris" (
  "id" serial NOT NULL PRIMARY KEY,
  "kode" character varying(20) NOT NULL,
  "nama" character varying(250) NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);