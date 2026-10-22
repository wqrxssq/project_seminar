-- Draft marker for the history service.
-- Domain tables are intentionally absent.
CREATE TABLE app_meta (
    key text PRIMARY KEY,
    value text NOT NULL,
    updated_at timestamptz NOT NULL DEFAULT now()
);

COMMENT ON TABLE app_meta IS 'Draft schema marker for history. Not the domain model.';

INSERT INTO app_meta (key, value)
VALUES ('service', 'history');
