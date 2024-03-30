# Write your MySQL query statement below
SELECT
    Emp.name,
    Bns.bonus
FROM Employee AS Emp
    LEFT JOIN
Bonus AS Bns ON Bns.empId = Emp.empId
    WHERE Bns.Bonus < 1000 OR Bns.bonus IS NULL;