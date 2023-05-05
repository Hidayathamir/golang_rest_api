CREATE TABLE public.products (
	id bigserial NOT NULL,
	"name" text NOT NULL,
	price int8 NOT NULL,
	description text NOT NULL,
	quantity int8 NOT NULL,
	CONSTRAINT products_pkey PRIMARY KEY (id)
);

INSERT INTO products ("name", price, description, quantity) VALUES('OKAY 58D Printer bluetooth mini printer thermal 58MM pos kasir ppob printer termal', 222000, 'desc okay 58d', 101);
INSERT INTO products ("name", price, description, quantity) VALUES('Thinkplus Lenovo TH10 Headphone Bluetooth Wireless Headset Earphone 5.0', 164900, 'desc thinkplus', 102);
INSERT INTO products ("name", price, description, quantity) VALUES('[Hanya ada di Shopee] Rexus Mouse Wireless Gaming SH10', 139000, 'desc hanya ada di shopee', 103);
