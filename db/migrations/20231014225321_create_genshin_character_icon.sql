-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE genshin_character_icon (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMPTZ,
  genshin_id VARCHAR(255) DEFAULT '',
  display_name_id UUID REFERENCES genshin_name(id),
  element_id UUID REFERENCES genshin_element(id),
  wepon_type_id UUID REFERENCES genshin_wepon_type(id)
);
CREATE TRIGGER update_genshin_character_icon
BEFORE UPDATE ON genshin_character_icon
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE genshin_character_icon;
-- +goose StatementEnd
