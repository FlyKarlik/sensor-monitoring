BEGIN;

ALTER TABLE sensor_data
    ADD COLUMN IF NOT EXISTS context TEXT,
    ADD COLUMN IF NOT EXISTS risk_rating FLOAT,
    ADD COLUMN IF NOT EXISTS inferred_brands TEXT[];

CREATE INDEX IF NOT EXISTS idx_sensor_data_inferred_brand ON sensor_data(inferred_brand);
CREATE INDEX IF NOT EXISTS idx_sensor_data_inferred_brands ON sensor_data USING GIN (inferred_brands);

COMMIT;