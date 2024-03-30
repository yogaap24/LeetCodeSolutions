# Write your MySQL query statement below
SELECT
    cust.name AS Customers
FROM Customers AS cust
    WHERE cust.id NOT IN (
        SELECT Orders.customerId FROM Orders
    );