-- +goose Up
-- +goose StatementBegin
CREATE TABLE Applications (
    CUSTOMER_RK Int32 NOT NULL,
    SOURSE_SYSTEM_ID_APPLICATION String,
    SOURCE_SYSTEM_CD_APPLICATION String
) ENGINE = MergeTree()
ORDER BY CUSTOMER_RK;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
