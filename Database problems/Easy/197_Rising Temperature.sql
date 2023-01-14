
-- Table: Weather
--
-- +---------------+---------+
-- | Column Name   | Type    |
-- +---------------+---------+
-- | id            | int     |
-- | recordDate    | date    |
-- | temperature   | int     |
-- +---------------+---------+
-- id is the primary key for this table.
-- This table contains information about the temperature on a certain day.
--
--
-- Write an SQL query to find all dates' Id with higher temperatures compared to its previous dates (yesterday).
--
-- Return the result table in any order.
--
-- The query result format is in the following example.

-- Write your MySQL query statement below

SELECT w1.id AS 'Id'
FROM Weather w1,
     Weather w2
WHERE w1.recordDate = DATE_ADD(w2.recordDate, INTERVAL 1 DAY)
  AND w1.temperature > w2.temperature




