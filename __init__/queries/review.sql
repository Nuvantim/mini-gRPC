-- CreateReview
INSERT INTO review (place_id,rating,comment) VALUES($1,$2,$3) RETURNING *;

-- UpdateReview
UPDATE review SET rating = $1, comment = $2
WHERE place_id = $3 RETURNING *; 

-- DeleteReview
DELETE FROM review WHERE id = $1;