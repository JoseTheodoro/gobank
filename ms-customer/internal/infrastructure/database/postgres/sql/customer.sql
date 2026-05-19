-- name: Create :one
INSERT INTO customers (customer_id, name, email, document, type, status) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;
