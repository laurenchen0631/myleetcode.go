/*
Write a SQL query to rank scores. If there is a tie between two scores, both should have the same ranking. 

Input:
Scores
+----+-------+
| Id | Score |
+----+-------+
| 1  | 3.50  |
| 2  | 3.65  |
| 3  | 4.00  |
| 4  | 3.85  |
| 5  | 4.00  |
| 6  | 3.65  |
+----+-------+

Output:
+-------+------+
| Score | Rank |
+-------+------+
| 4.00  | 1    |
| 4.00  | 1    |
| 3.85  | 2    |
| 3.65  | 3    |
| 3.65  | 3    |
| 3.50  | 4    |
+-------+------+
*/

SELECT
  Score,
  @rank := @rank + (@prev <> (@prev := Score)) AS Rank
FROM
  Scores,
  (SELECT @rank := 0, @prev := -1) INIT
ORDER BY Score DESC