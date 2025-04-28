package repository

const (
	countSensorData = `
SELECT COUNT(*) AS total 
FROM sensor_data 
WHERE
(inferred_brand ILIKE $1 OR $1 IS NULL)
AND (inferred_brands @> $2 OR $2 IS NULL)
AND (attestation = $3 OR $3 IS NULL)
AND (
    CASE
        WHEN $4::BOOLEAN IS NULL THEN
            TRUE
        WHEN $4::BOOLEAN = TRUE THEN
            recording_file IS NOT NULL
        WHEN $4::BOOLEAN = FALSE THEN
            recording_file IS NULL
    END
)
AND ( 
    $5::INT IS NULL 
    OR 
    LENGTH(TRIM(transcript)) >= $5::INT
);
	`
	searchSensorData = `
SELECT
    *
FROM
    sensor_data
WHERE
    (inferred_brand ILIKE $1 OR $1 IS NULL)
    AND (inferred_brands @> $2 OR $2 IS NULL)   
    AND (attestation = $3 OR $3 IS NULL)
    AND (
        CASE 
            WHEN $4::BOOLEAN IS NULL THEN
                TRUE
            WHEN $4::BOOLEAN = TRUE THEN
                recording_file IS NOT NULL
            WHEN $4::BOOLEAN = FALSE THEN
                recording_file IS NULL
        END
    )
    AND (
        $5::INT IS NULL 
        OR 
        LENGTH(TRIM(transcript)) >= $5::INT
    )
ORDER BY 
    CASE
        WHEN $6::BOOLEAN = FALSE OR $6::BOOLEAN IS NULL THEN
            CASE $7::TEXT
                WHEN 'created_at' THEN created_at::TEXT
                ELSE NULL
            END
        ELSE NULL
    END ASC,
    CASE
        WHEN $6::BOOLEAN = TRUE THEN
            CASE $7::TEXT
                WHEN 'created_at' THEN created_at::TEXT
                ELSE NULL
            END
        ELSE NULL
    END DESC
LIMIT 
    COALESCE($8::INT, 100)
OFFSET 
    COALESCE($9::INT, 0);
	`
)
