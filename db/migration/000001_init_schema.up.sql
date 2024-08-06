CREATE TABLE "customers" (
  "customer_uuid" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "tenant_id" varchar NOT NULL,
  "customer_id" varchar NOT NULL unique,
  "customer_name" varchar,
  "customer_address" text,
  "customer_city" varchar,
  "customer_state" varchar,
  "customer_country" varchar,
  "customer_total_value" bigint NOT NULL,
  "customer_status" varchar,
  "customer_app_type" varchar,
  "customer_reference" varchar,
  "customer_app_size" varchar,
  "customer_primary_email" varchar,
  "customer_primary_phone" varchar,
  "customer_secondary_email" varchar,
  "customer_secondary_phone" varchar
);

CREATE TABLE "customer_transcations" (
  "transaction_id" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "invoice_id" varchar NOT NULL unique,
  "customer_uuid" uuid,
  "tenant_id" varchar NOT NULL,
  "amount" bigint,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now()),
  "transcation_debit" bool
);

CREATE TABLE "vendor_transcations" (
  "transaction_id" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "invoice_id" varchar NOT NULL unique,
  "vendor_uuid" uuid,
  "tenant_id" varchar NOT NULL,
  "amount" bigint,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now()),
  "transaction_type" varchar,
  "transcation_debit" bool
);

CREATE TABLE "internal_transcations" (
  "transaction_id" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "invoice_id" varchar NOT NULL unique,
  "user_name" varchar,
  "tenant_id" varchar NOT NULL,
  "amount" bigint,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now()),
  "transcation_debit" bool
);

CREATE TABLE "vendors" (
  "vendor_uuid" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "vendor_name" varchar,
  "vendor_id" varchar NOT NULL unique,
  "vendor_address" text,
  "vendor_city" varchar,
  "vendor_state" varchar,
  "vendor_country" varchar,
  "vendor_email" varchar,
  "vendor_phone_number" varchar,
  "tenant_id" varchar NOT NULL
);

CREATE TABLE "quotations" (
  "quotation_uuid" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "customer_uuid" uuid,
  "quotation_id" varchar NOT NULL unique,
  "tenant_id" varchar NOT NULL,
  "quotation_value" bigint,
  "quatation_type" varchar,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now()),
  "converted" bool
);

CREATE TABLE "service_details" (
  "customer_uuid" uuid,
  "under_warranty" bool,
  "next_service_date" date,
  "service_amount" bigint,
  "tenant_id" varchar NOT NULL
);

CREATE INDEX ON "customers" ("tenant_id");

CREATE INDEX ON "customers" ("customer_status");

CREATE INDEX ON "customer_transcations" ("tenant_id");

CREATE INDEX ON "customer_transcations" ("customer_uuid");

CREATE INDEX ON "vendor_transcations" ("tenant_id");

CREATE INDEX ON "vendor_transcations" ("transaction_type");

CREATE INDEX ON "internal_transcations" ("tenant_id");

CREATE INDEX ON "vendors" ("tenant_id");

CREATE INDEX ON "quotations" ("tenant_id");

CREATE INDEX ON "service_details" ("tenant_id");

ALTER TABLE "customer_transcations" ADD FOREIGN KEY ("customer_uuid") REFERENCES "customers" ("customer_uuid");

ALTER TABLE "vendor_transcations" ADD FOREIGN KEY ("vendor_uuid") REFERENCES "vendors" ("vendor_uuid");

ALTER TABLE "quotations" ADD FOREIGN KEY ("customer_uuid") REFERENCES "customers" ("customer_uuid");

ALTER TABLE "service_details" ADD FOREIGN KEY ("customer_uuid") REFERENCES "customers" ("customer_uuid");