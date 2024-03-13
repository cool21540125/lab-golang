
# golang 練功

- https://www.bogotobogo.com/GoLang/GoLang_Function_taking_a_slice_and_returning_a_slice.php


Qid  | Date       | Level | File
---- | ---------- | ----- | ---------------
1    | 2023.09.10 | E     | [Two Sum](./leetcode/TwoSum.go)
20   | 2023.09.14 | E     | [Valid Parentheses](./leetcode/ValidParentheses.go)
260  | 2023.09.14 | M     | [Single Number III](./leetcode/SingleNumberIII.go)
136  | 2023.09.16 | E     | [Single Number](./leetcode/SingleNumber.go)
2    | 2023.09.21 | M     | [Add Two Numbers](./leetcode/AddTwoNumbers.go)
643  | 2023.11.26 | E     | [Maximum Average Subarray I](./leetcode/MaximumAverageSubarrayI.go)
2485 | 2024.03.13 | E     | [Find the Pivot Integer](./leetcode/FindThePivotInteger.go)


```bash

go test leetcode/*.go -cover

go test leetcode/*.go -coverprofile=coverage.out && go tool cover -html=coverage.out
```
