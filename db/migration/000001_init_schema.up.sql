CREATE TABLE "customers" (
  "customer_id" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "tenant_id" varchar NOT NULL,
  "customer_name" varchar,
  "customer_address" text,
  "customer_city" varchar,
  "customer_state" varchar,
  "customer_country" varchar,
  "customer_total_value" bigint,
  "customer_status" varchar
);

CREATE TABLE "customer_details" (
  "customer_id" uuid,
  "tenant_id" varchar NOT NULL,
  "customer_app_type" varchar,
  "customer_reference" varchar,
  "customer_app_size" varchar
);

CREATE TABLE "customer_contact" (
  "customer_id" uuid,
  "tenant_id" varchar NOT NULL,
  "customer_email_1" varchar,
  "customer_phone_number_1" varchar,
  "customer_email_2" varchar,
  "customer_phone_number_2" varchar
);

CREATE TABLE "customer_transcations" (
  "transaction_id" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "invoice_id" varchar,
  "customer_id" uuid,
  "tenant_id" varchar NOT NULL,
  "amount" bigint,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now()),
  "transcation_debit" bool
);

CREATE TABLE "vendor_transcations" (
  "transaction_id" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "invoice_id" varchar,
  "vendor_id" uuid,
  "tenant_id" varchar NOT NULL,
  "amount" bigint,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now()),
  "transaction_type" varchar,
  "transcation_debit" bool
);

CREATE TABLE "internal_transcations" (
  "transaction_id" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "invoice_id" varchar,
  "user_name" varchar,
  "tenant_id" varchar NOT NULL,
  "amount" bigint,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now()),
  "transcation_debit" bool
);

CREATE TABLE "vendors" (
  "vendor_id" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "vendor_name" varchar,
  "vendor_address" text,
  "vendor_city" varchar,
  "vendor_state" varchar,
  "vendor_country" varchar,
  "vendor_email" varchar,
  "vendor_phone_number" varchar,
  "tenant_id" varchar NOT NULL
);

CREATE TABLE "quotations" (
  "quotation_id" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "customer_id" uuid,
  "tenant_id" varchar NOT NULL,
  "quotation_value" bigint,
  "quatation_type" varchar,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now()),
  "converted" bool
);

CREATE TABLE "service_details" (
  "customer_id" uuid,
  "under_warranty" bool,
  "next_service_date" date,
  "service_amount" bigint,
  "tenant_id" varchar NOT NULL
);

CREATE INDEX ON "customers" ("tenant_id");

CREATE INDEX ON "customers" ("customer_status");

CREATE INDEX ON "customer_details" ("tenant_id");

CREATE INDEX ON "customer_contact" ("tenant_id");

CREATE INDEX ON "customer_transcations" ("tenant_id");

CREATE INDEX ON "customer_transcations" ("customer_id");

CREATE INDEX ON "vendor_transcations" ("tenant_id");

CREATE INDEX ON "vendor_transcations" ("transaction_type");

CREATE INDEX ON "internal_transcations" ("tenant_id");

CREATE INDEX ON "vendors" ("tenant_id");

CREATE INDEX ON "quotations" ("tenant_id");

CREATE INDEX ON "service_details" ("tenant_id");

ALTER TABLE "customer_details" ADD FOREIGN KEY ("customer_id") REFERENCES "customers" ("customer_id");

ALTER TABLE "customer_contact" ADD FOREIGN KEY ("customer_id") REFERENCES "customers" ("customer_id");

ALTER TABLE "customer_transcations" ADD FOREIGN KEY ("customer_id") REFERENCES "customers" ("customer_id");

ALTER TABLE "vendor_transcations" ADD FOREIGN KEY ("vendor_id") REFERENCES "vendors" ("vendor_id");

ALTER TABLE "quotations" ADD FOREIGN KEY ("customer_id") REFERENCES "customers" ("customer_id");

ALTER TABLE "service_details" ADD FOREIGN KEY ("customer_id") REFERENCES "customers" ("customer_id");