-- Part 01

.read parser.sql

SELECT SUM(ABS(a-b)) AS answer
FROM (
    SELECT a, row_number() OVER(ORDER BY a) as rnk from day01
) a_col
JOIN (
    SELECT b, row_number() OVER(ORDER BY b) as rnk from day01
) b_col
ON a_col.rnk=b_col.rnk;

