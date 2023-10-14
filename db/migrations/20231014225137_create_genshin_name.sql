-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE genshin_name (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMPTZ,
  genshin_id VARCHAR(255) DEFAULT '',
  name VARCHAR(255) NOT NULL,
  language_id UUID NOT NULL REFERENCES language(id)
);
CREATE TRIGGER update_genshin_name
BEFORE UPDATE ON genshin_name
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE genshin_name;
-- +goose StatementEnd
