-- name: CreateOnboardingProcess :one
INSERT INTO onboarding_processes
    (onboarding_id, email, document, status)
values ($1, $2, $3, $4)
    RETURNING *;