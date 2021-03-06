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

CREATE TABLE "log_files" (
  "id" serial NOT NULL,
  "file_id" character varying(100) PRIMARY KEY,
  "row_id" character varying(100) NOT NULL,
  "tipe" character varying(20) NOT NULL,
  "url" character varying(250) NOT NULL,
  "filename" character varying(250) NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "log_surat_masuks" (
  "id" serial NOT NULL,
  "row_id" character varying(100) PRIMARY KEY,
  "no_surat" character varying(100) NOT NULL,
  "tanggal_surat" character varying(100) NOT NULL,
  "dari" character varying(250) NOT NULL,
  "perihal" character varying(250) NOT NULL,
  "isi" text NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "log_surat_keluars" (
  "id" serial NOT NULL,
  "row_id" character varying(100) PRIMARY KEY,
  "no_surat" character varying(100) NOT NULL,
  "tanggal_surat" character varying(100) NOT NULL,
  "tujuan" character varying(250) NOT NULL,
  "perihal" character varying(250) NOT NULL,
  "isi" text NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "log_unit_kerjas" (
  "id" serial NOT NULL PRIMARY KEY,
  "nama" character varying(250) NOT NULL,
  "alamat" character varying(250) NOT NULL,
  "telepon" character varying(250) NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "log_user_unit_kerja" (
  "user_id" integer NOT NULL,
  "unit_kerja_id" integer NOT NULL
);

ALTER TABLE "log_surat_masuks"
ADD "unit_kerja_id" integer NOT NULL;

ALTER TABLE "log_surat_keluars"
ADD "unit_kerja_id" integer NOT NULL;

CREATE TABLE "log_metric" (
  "id" serial NOT NULL PRIMARY KEY,
  "nama" character varying(250) NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "log_jukniss" (
  "id" serial NOT NULL,
  "row_id" character varying(100) PRIMARY KEY,
  "unit_kerja_id" integer NOT NULL,
  "no_surat" character varying(100) NOT NULL,
  "tanggal_surat" character varying(100) NOT NULL,
  "perihal" character varying(250) NOT NULL,
  "isi" text NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "log_barangs" (
  "id" serial NOT NULL,
  "row_id" character varying(100) PRIMARY KEY,
  "unit_kerja_id" integer NOT NULL,
  "kategori_id" integer NOT NULL,
  "metric_id" integer NOT NULL,
  "kode" character varying(100) NOT NULL,
  "nama" character varying(250) NOT NULL,
  "jumlah" bigint NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "log_barang_masuks" (
  "id" serial NOT NULL,
  "row_id" character varying(100) PRIMARY KEY,
  "barang_id" integer NOT NULL,
  "tanggal" character varying(20) NOT NULL,
  "jumlah" integer NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE "log_barang_masuks"
ALTER "barang_id" TYPE character varying(100),
ALTER "barang_id" DROP DEFAULT,
ALTER "barang_id" SET NOT NULL;
COMMENT ON COLUMN "log_barang_masuks"."barang_id" IS '';
COMMENT ON TABLE "log_barang_masuks" IS '';

ALTER TABLE "log_barang_masuks"
ADD "unit_kerja_id" character varying(100) NOT NULL;
COMMENT ON TABLE "log_barang_masuks" IS '';

ALTER TABLE "log_barang_masuks"
DROP "unit_kerja_id";
COMMENT ON TABLE "log_barang_masuks" IS '';

ALTER TABLE "log_barang_masuks"
ADD "unit_kerja_id" integer NOT NULL;
COMMENT ON TABLE "log_barang_masuks" IS '';

CREATE TABLE "log_barang_keluars" (
  "id" serial NOT NULL,
  "row_id" character varying(100) PRIMARY KEY,
  "barang_id" character varying(100) NOT NULL,
  "unit_kerja_id" integer NOT NULL,
  "tanggal" character varying(20) NOT NULL,
  "jumlah" integer NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

