-- name: CreateEntry :one 
INSERT INTO entries (
    journal_id,
    instrument, 
    position, 
    lot_size, 
    opening, 
    closing, 
    stop_loss, 
    take_profit, 
    risk_reward, 
    comments,
    before_image, 
    after_image 
) VALUES ( 
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
)
RETURNING *; 

-- name: GetEntry :one 
SELECT * FROM entries 
WHERE id = $1 LIMIT 1; 

-- name: ListEntries :many 
SELECT * FROM entries 
ORDER BY id 
LIMIT $1 
OFFSET $2; 

-- name: UpdateEntry :one
UPDATE entries 
SET journal_id = $2, 
    instrument = $3, 
    position = $4, 
    lot_size = $5, 
    opening = $6, 
    closing = $7, 
    stop_loss = $8, 
    take_profit = $9, 
    risk_reward = $10, 
    comments = $11, 
    before_image = $12, 
    after_image = $13
WHERE id = $1
RETURNING *; 

-- name: DeleteEntry :exec 
DELETE FROM entries WHERE id = $1;