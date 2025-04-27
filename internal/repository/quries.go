package repository

const (
	countSensorData = `
SELECT COUNT(*) AS total 
FROM sensor_data 
WHERE
(inferred_brand ILIKE $1 OR $1 IS NULL)
AND (attestation = $2 OR $2 IS NULL)
AND (
    CASE
        WHEN $3::BOOLEAN IS NULL THEN
            TRUE
        WHEN $3::BOOLEAN = TRUE THEN
            recording_file IS NOT NULL
        WHEN $3::BOOLEAN = FALSE THEN
            recording_file IS NULL
    END
)
AND ( 
    $4::INT IS NULL 
    OR 
    LENGTH(TRIM(transcript)) >= $4::INT
);
	`
	searchSensorData = `
SELECT
    *
FROM
    sensor_data
WHERE
    (inferred_brand ILIKE $1 OR $1 IS NULL)
    AND (attestation = $2 OR $2 IS NULL)
    AND (
        CASE 
            WHEN $3::BOOLEAN IS NULL THEN
                TRUE
            WHEN $3::BOOLEAN = TRUE THEN
                recording_file IS NOT NULL
            WHEN $3::BOOLEAN = FALSE THEN
                recording_file IS NULL
        END
    )
    AND (
        $4::INT IS NULL 
        OR 
        LENGTH(TRIM(transcript)) >= $4::INT
    )
ORDER BY 
    CASE
        WHEN $5::BOOLEAN = FALSE OR $5::BOOLEAN IS NULL THEN
            CASE $6::TEXT
                WHEN 'created_at' THEN created_at::TEXT
                ELSE NULL
            END
        ELSE NULL
    END ASC,
    CASE
        WHEN $5::BOOLEAN = TRUE THEN
            CASE $6::TEXT
                WHEN 'created_at' THEN created_at::TEXT
                ELSE NULL
            END
        ELSE NULL
    END DESC
LIMIT 
    COALESCE($7::INT, 100)
OFFSET 
    COALESCE($8::INT, 0);
	`
)
