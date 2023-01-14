-- Table: Logs
--
-- +-------------+---------+
-- | Column Name | Type    |
-- +-------------+---------+
-- | id          | int     |
-- | num         | varchar |
-- +-------------+---------+
-- id is the primary key for this table.
-- id is an autoincrement column.
--
--
-- Write an SQL query to find all numbers that appear at least three times consecutively.
--
-- Return the result table in any order.
--
-- The query result format is in the following example.

-- Write your MySQL query statement below.

SELECT DISTINCT num AS ConsecutiveNums
FROM Logs
WHERE id BETWEEN 3 AND (SELECT MAX(id) FROM Logs)
AND num IN (SELECT num FROM Logs WHERE id BETWEEN 1 AND (SELECT MAX(id) - 1 FROM Logs))
AND num IN (SELECT num FROM Logs WHERE id BETWEEN 2 AND (SELECT MAX(id) FROM Logs))





