CREATE TABLE products (
	id bigserial NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	"name" text NOT NULL,
	price int8 NOT NULL,
	description text NOT NULL,
	quantity int8 NOT NULL,
	CONSTRAINT products_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_products_deleted_at ON public.products USING btree (deleted_at);

INSERT INTO products (created_at, updated_at, deleted_at, "name", price, description, quantity) VALUES('2020-05-05 16:43:24.437', '2020-05-05 16:43:48.912', NULL, 'OKAY 58D Printer bluetooth mini printer thermal 58MM pos kasir ppob printer termal', 222000, 'desc okay 58d', 101);
INSERT INTO products (created_at, updated_at, deleted_at, "name", price, description, quantity) VALUES('2021-05-05 16:43:48.919', '2021-05-05 16:43:48.919', NULL, 'Thinkplus Lenovo TH10 Headphone Bluetooth Wireless Headset Earphone 5.0', 164900, 'desc thinkplus', 102);
INSERT INTO products (created_at, updated_at, deleted_at, "name", price, description, quantity) VALUES('2022-05-05 16:43:48.920', '2022-05-05 16:43:48.920', NULL, '[Hanya ada di Shopee] Rexus Mouse Wireless Gaming SH10', 139000, 'desc hanya ada di shopee', 103);
