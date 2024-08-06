-- name: CreateCustomer :one
INSERT INTO customers (
  "tenant_id",
  "customer_id",
  "customer_name" ,
  "customer_address",
  "customer_city" ,
  "customer_state" ,
  "customer_country" ,
  "customer_total_value" ,
  "customer_status",
  "customer_app_type" ,
  "customer_reference" ,
  "customer_app_size" ,
  "customer_primary_email" ,
  "customer_primary_phone" ,
  "customer_secondary_email" ,
  "customer_secondary_phone" 
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9 , $10, $11, $12, $13, $14, $15, $16
) RETURNING * ;

-- name: ListCustomers :many
select * from customers
where tenant_id = $1
LIMIT $2
OFFSET $3;


-- name: ListCustomerByID :one
select *
from customers
where tenant_id = $1 AND customer_uuid = $2;

-- name: DeleteCustomerByID :exec
delete from customers
where tenant_id = $1 and customer_uuid = $2;