## What is a test?

> A test is a repeatable process that verifies whether or not something is working as intended.

Virtually every other detail is variable

- How a test is performed (manual vs automated)
- What exactly is verified
- How often the test is run
- The scope

Despite this being a pretty decent description of WHAT a test is, a test isn't used solely for verifying that something is working.

## This repo will focus on automated software testing 

Short version - we will write code that runs our testing

```go
if someVal == whatIWant {
    // test passes! =D
} else {
    // test fails. :(
}
```

## Why do tests matter?

1. Tests help find, fix, and prevent mistakes (bugs, side effects, edge cases, ect)
2. Tests document expected behaviour
3. Tests encourage us to write better code
4. Tests can speed up development


## Writing Great Test

- Tests with a purpose 

    ```go
    const SomeString = "this is my string"
    ```

    ```go
    func TestSomeString(t *testing.T)  {
        want := "this is my string"
        if pkg.SomeString != want {
            t.ErrorF("SOmeString = %q; want %q", pkg.SomeString, want)
        }
    }
    ```
- Don't overdo it

> I get paid for code that works, not for tests, so my philosophy is to test as little as possible to reac a given level of confidence (I suspect this level of confidence )
- Testing is a skill, skills evolve


## File Naming Conventions
- Anything with _test is going to be a test and wiull be ignored when the compiler builds the code
- typically you will see a one to one correlation between source files and test files

There are a few caveats to this on topics like 
- `export_test.go` to access unexported variables in external tests
- `xx_internal_tests.go` for internal tests
- `example_xxx_test.go` for examples in isolated files

## Function naming Conventions
The strict rules are 
- When ever you write a test it must be a function
- It must start with a capital T for Test followed by an upercase case for the name of the function or type that is being tested ie. func TestXxx

## Varialbe Naming
When naming the vaiables for the test in following with common practice 
- We often name the variable we take in to the function with arg
- We name the output of the function got 
- We name the expected result as want


