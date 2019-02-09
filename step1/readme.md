# Step1 - Basic unit test

The Goal here is to talk about code coverage.

## The situation

Take your time to look and play with the very complex program (main.go) and the unit test (main_test.go).

## The issue

What wrong ? We only one case of the function canBuyAlcool !

How can I detect ?

## The solution

Run unit test with the code coverage and check the report.
```
go test -cover -coverprofile=c.out
go tool cover -html=c.out -o coverage.html
```

The coverage.html will show that the line `return false` of the function canBuyAlcool isn't covered.

=> So we can add a new test to cover this case.
