-- +goose Up
-- +goose StatementBegin
CREATE TABLE Address (
    SUBJ_ID String NOT NULL,
    LAST_NM String,
    FIRST_NM String,
    MIDDLE_NM String,
    BIRTH_DT Date,
    COUNTRY_REAL String,
    INDEX_REAL String,
    REGION_REAL String,
    CITY_REAL String,
    STREET_REAL String,
    BUILDING_REAL String,
    FLAT_REAL String,
    PHONE_HOME String,
    PHONE_MOBILE String,
    PHONE_WORK String,
    PHONE_CONTACT String,
    PHONE_MOBILE_MAIN String,
    PSP_SERIES String,
    PSP_NO String,
    customer_rk Int32,
    FLG_Actual_ID Int32 NOT NULL
) ENGINE = MergeTree()
ORDER BY SUBJ_ID;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
