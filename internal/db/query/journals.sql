-- name: CreateJournal :one 
INSERT INTO journals (
    user_id,
    title, 
    strategy, 
    initial_balance,
    current_balance, 
    risk_tolerance 
) VALUES (
    $1, $2, $3, $4, $5, $6 
) 
RETURNING *; 

-- name: GetJournal :one
SELECT * FROM journals
WHERE id = $1 LIMIT 1; 

-- name: ListJournals :many 
SELECT * FROM journals 
ORDER BY id 
LIMIT $1 
OFFSET $2; 

-- name: UpdateJournal :one 
UPDATE journals 
SET user_id = $2, 
    title = $3, 
    strategy = $4, 
    initial_balance = $5, 
    current_balance = $6, 
    risk_tolerance = $7 
WHERE id = $1
RETURNING *; 

-- name: DeleteJournal :exec 
DELETE FROM journals WHERE id = $1; 
