CREATE OR REPLACE TABLE day01 AS (
SELECT
    CAST(regexp_extract(raw, '(\d+)\s+?') AS INT) AS a,
    CAST(regexp_extract(raw, '\s+(\d+)') AS INT) AS b
FROM  READ_CSV_AUTO('../inputs/puzzle-01.txt', header=0, columns={'raw': varchar})
);
