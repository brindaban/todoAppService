
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE tasks (
  description   VARCHAR(255),
  priority      VARCHAR(255),
  finished      BOOLEAN
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE tasks;
