package repository

const (
	countSensorData = `
SELECT COUNT(*) AS total 
FROM sensor_data 
WHERE
(inferred_brand ILIKE $1 OR inferred_brand IS NULL);
	`
	searchSensorData = `
SELECT
    *
FROM
    sensor_data
WHERE
    (inferred_brand ILIKE $1 OR inferred_brand IS NULL)
ORDER BY 
    CASE
        WHEN $2::BOOLEAN = FALSE OR $2::BOOLEAN IS NULL THEN
            CASE $3::TEXT
                WHEN 'created_at' THEN created_at::TEXT
                ELSE NULL
            END
        ELSE NULL
    END ASC,
    CASE
        WHEN $2::BOOLEAN = TRUE THEN
            CASE $3::TEXT
                WHEN 'created_at' THEN created_at::TEXT
                ELSE NULL
            END
        ELSE NULL
    END DESC
LIMIT 
    COALESCE($4::INT, 100)
OFFSET 
    COALESCE($5::INT, 0);
	`
)
