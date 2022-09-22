-- name: CreateSubscriptionTier :one 
INSERT INTO subscription_tiers ( 
    name
) VALUES ( 
    $1 
)
RETURNING *; 

-- name: GetSubscriptionTier :one 
SELECT * FROM subscription_tiers
WHERE id = $1 LIMIT 1; 

-- name: ListSubscriptionTiers :many 
SELECT * FROM subscription_tiers
ORDER BY id
LIMIT $1 
OFFSET $2; 

-- name: UpdateSubscriptionTier :one 
UPDATE subscription_tiers 
SET name = $2
WHERE id = $1
RETURNING *; 

-- name: DeleteSubscriptionTier :exec 
DELETE FROM subscription_tiers WHERE id = $1;