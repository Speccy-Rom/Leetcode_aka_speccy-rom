-- Table: Employee
--
-- +-------------+---------+
-- | Column Name | Type    |
-- +-------------+---------+
-- | id          | int     |
-- | name        | varchar |
-- | salary      | int     |
-- | managerId   | int     |
-- +-------------+---------+
-- id is the primary key column for this table.
-- Each row of this table indicates the ID of an employee, their name, salary, and the ID of their manager.
--
--
-- Write an SQL query to find the employees who earn more than their managers.
--
-- Return the result table in any order.
--
-- The query result format is in the following example.

--
-- Table: Employee
--
-- Write your MySQL query statement below

SELECT e1.name AS Employee
FROM Employee e1, Employee e2
WHERE e1.managerId = e2.id AND e1.salary > e2.salary

SELECT e1.name AS Employee
FROM Employee e1
JOIN Employee e2
ON e1.managerId = e2.id
WHERE e1.salary > e2.salary

SELECT e1.name AS Employee
FROM Employee e1
INNER JOIN Employee e2
ON e1.managerId = e2.id
WHERE e1.salary > e2.salary

SELECT e1.name AS Employee
FROM Employee e1
WHERE e1.salary > (SELECT e2.salary FROM Employee e2 WHERE e2.id = e1.managerId)


