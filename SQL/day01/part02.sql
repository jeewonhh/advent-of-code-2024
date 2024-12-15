-- Part 02

.read parser.sql

SELECT SUM(freq * a) AS answer
FROM (
    SELECT b, COUNT(*) AS freq
    FROM day01 GROUP BY b
) B
JOIN (
    SELECT a FROM day01
) A
ON B.b=A.a;