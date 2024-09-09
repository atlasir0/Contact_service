-- +goose Up
-- +goose StatementBegin
CREATE TABLE Phone (
    SUBJ_ID String NOT NULL,
    LAST_NM String,
    FIRST_NM String,
    MIDDLE_NM String,
    BIRTH_DT Date,
    Contact_Type String NOT NULL,
    Contact String,
    PSP_SERIES String,
    PSP_NO String,
    customer_rk Int32,
    FLG_Actual_ID Int32 NOT NULL
) ENGINE = MergeTree()
ORDER BY SUBJ_ID;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE TABLE Phone 
-- +goose StatementEnd
