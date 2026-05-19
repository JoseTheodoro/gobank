-- name: CreateOnboardingProcess :one
INSERT INTO onboarding_processes
    (onboarding_id, email, document, status)
values ($1, $2, $3, $4)
    RETURNING *;

-- name: SetCustomerOnboardingProcess :exec
UPDATE onboarding_processes
SET updated_at = now(),
    customer_id = $1,
    status = $2
WHERE id = $3
    RETURNING *;

-- name: SetAccountOnboardingProcess :exec
UPDATE onboarding_processes
SET update_at = now(),
    account_id = $1,
    status = $2
WHERE id = $3
    RETURNING *;

