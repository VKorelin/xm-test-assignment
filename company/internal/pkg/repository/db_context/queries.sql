-- name: GetCompany :one
SELECT * FROM companies
WHERE id = $1 LIMIT 1;

-- name: DeleteCompany :exec
DELETE FROM companies
WHERE id = $1;

-- name: NewCompany :one
INSERT INTO companies (name, description, amount_of_employees, registered, type) VALUES ($1, $2, $3, $4, $5) RETURNING id;

-- name: UpdateCompany :exec
UPDATE companies SET name = $1, description = $2, amount_of_employees = $3, registered = $4, type = $5 WHERE id = $6;
