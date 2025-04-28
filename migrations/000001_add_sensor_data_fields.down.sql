BEGIN;

ALTER TABLE sensor_data
    DROP COLUMN IF EXISTS context,
    DROP COLUMN IF EXISTS risk_rating,
    DROP COLUMN IF EXISTS inferred_brands;

DROP INDEX IF EXISTS idx_sensor_data_inferred_brand;
DROP INDEX IF EXISTS idx_sensor_data_inferred_brands;

COMMIT;