# Step3 - The mutation testing

The Goal here is to talk about mutation testing

## The situation

We improve our previous tests and handle every case (age/volume).
Take your time to look and play with the very complex program (main.go) and the unit test (main_test.go)

## The issue

We only test basic case and not limit case. For example:
 * Do a 18 years old person can buy alcool ?
 * Can we buy 10 litters of alcool ?

How can I detect ?

## The solution

The mutation testing.

The goal of the mutation testing is to automaticaly test our tests.
It will do some modification on the code (mutation) and check if a test fail.
If a mutants still alive after all tests, we need to improve our tests. [Wikipedia](https://en.wikipedia.org/wiki/Mutation_testing)

One example of mutation is to transfrom `age <= 18` into `age < 18`.

We can run use [go-mutesting](https://github.com/zimmski/go-mutesting)

```
go-mutesting .
```

There is 2 mutants who still alive, `age > 18` and `volume < 10`.

More over if we remove test cases with `volume == 20`, there is another mutant, `age >= 18 && true` who still alive.

So we can detect uncovered branch.
