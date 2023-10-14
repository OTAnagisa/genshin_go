-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE FUNCTION set_updated_at()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = now();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP FUNCTION set_updated_at;
-- +goose StatementEnd
