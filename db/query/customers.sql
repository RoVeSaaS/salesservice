-- name: CreateCustomer :one
INSERT INTO customers (
  "tenant_id",
  "customer_name" ,
  "customer_address",
  "customer_city" ,
  "customer_state" ,
  "customer_country" ,
  "customer_total_value" ,
  "customer_status"
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING * ;

-- name: CreateCustomerDetails :one
INSERT INTO "customer_details" (
  "customer_id",
  "tenant_id",
  "customer_app_type" ,
  "customer_reference" ,
  "customer_app_size"
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING * ;

-- name: CreateCustomerContat :one
INSERT INTO "customer_contact" (
  "customer_id",
  "tenant_id" ,
  "customer_email_1",
  "customer_phone_number_1" ,
  "customer_email_2" ,
  "customer_phone_number_2"
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING * ;

-- name: ListCustomers :many
select c.customer_id, c.tenant_id,c.customer_name,c.customer_address,c.customer_city,c.customer_state,c.customer_country,c.customer_total_value,c.customer_status,
d.customer_app_type,d.customer_reference,d.customer_app_size,
cc.customer_email_1 AS primary_email,cc.customer_phone_number_1 AS primary_phone,
cc.customer_email_2 AS secondary_email,cc.customer_phone_number_2 AS secondary_phone
from customers c
left join customer_details d 
on c.customer_id = d.customer_id and c.tenant_id = d.tenant_id
left join customer_contact cc 
on c.customer_id = cc.customer_id and c.tenant_id = cc.tenant_id
where c.tenant_id = $1 ;


-- name: ListCustomerByID :one
select c.customer_id, c.tenant_id,c.customer_name,c.customer_address,c.customer_city,c.customer_state,c.customer_country,c.customer_total_value,c.customer_status,
       d.customer_app_type,d.customer_reference,d.customer_app_size,
       cc.customer_email_1 AS primary_email,cc.customer_phone_number_1 AS primary_phone,
       cc.customer_email_2 AS secondary_email,cc.customer_phone_number_2 AS secondary_phone
from customers c 
left join customer_details d 
on c.customer_id = d.customer_id and c.tenant_id = d.tenant_id
left join customer_contact cc 
on c.customer_id = cc.customer_id and c.tenant_id = cc.tenant_id
where c.tenant_id = $1 AND c.customer_id = $2 ;

-- name: DeleteCustomerByID :exec
delete from customers
where tenant_id = $1 and customer_id = $2;

-- name: DeleteCustomerDetailsByID :exec
delete from customer_details
where tenant_id = $1 and customer_id = $2;

-- name: DeleteCustomerContactByID :exec
delete from customer_contact
where tenant_id = $1 and customer_id = $2;