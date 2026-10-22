-- Draft marker for the documents service.
-- Domain tables are intentionally absent.
CREATE TABLE app_meta (
    key text PRIMARY KEY,
    value text NOT NULL,
    updated_at timestamptz NOT NULL DEFAULT now()
);

COMMENT ON TABLE app_meta IS 'Draft schema marker for documents. Not the domain model.';

INSERT INTO app_meta (key, value)
VALUES ('service', 'documents');
