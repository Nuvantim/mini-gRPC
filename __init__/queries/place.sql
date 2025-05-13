-- GetPlace
SELECT * FROM place;

-- FindPlace;
SELECT * FROM place WHERE id = $1;

-- CreatePlace
INSERT INTO place (name,location,description,entry_fee) VALUES ($1,$2,$3,$4) RETURNING *;

-- UpdatePlace
UPDATE place SET name=$1, location=$2, description=$3, entry_fee=$4 WHERE id = $5;

-- DeletePlace
DELETE from place WHERE id = $1;
