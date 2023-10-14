-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE genshin_icon_type (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMPTZ,
  genshin_id VARCHAR(255) DEFAULT '',
  name VARCHAR(255) NOT NULL,
  code int NOT NULL
);
CREATE TRIGGER update_genshin_icon_type
BEFORE UPDATE ON genshin_icon_type
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE genshin_icon_type;
-- +goose StatementEnd
