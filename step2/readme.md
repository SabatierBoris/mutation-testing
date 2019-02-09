# Step2 - The branch coverage

The Goal here is to talk about branch coverage

## The situation

We improve our previous app.
Take your time to look and play with the very complex program (main.go) and the unit test (main_test.go)

## The issue

The code coverage say the function canBuyAlcool is fully complet, so what's the problem ?

=> We only test with a volume of alcool of 5 litters.

How can I detect ?

## The solution

The branch coverage is a solution where every branch of a statement is evaluated separately.

But, at this time, we can't do a branch coverage in go. see [cover.go](https://golang.org/src/cmd/cover/cover.go#L142) and a [closed issue](https://github.com/golang/go/issues/28888).

So we only need to be meticulous when we write test and when other review them.
