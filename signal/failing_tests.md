## Signalling Failure
The `testing.T` Type has a `Log` and `Logf` method. These work similarly to `Print` and `Printf` in the fmt package

There are two basic ways to signal that a test has failed:
- Fail = fail, but keep running 
- FailNow = fail now and stop running tests

You will rarely see this called though, because most of the time people will call the following methods which combine failures and logging
- Error = Log + Fail
- Errorf = Logf + Fail
- Fatal = Log + FailNow
- Fatalf = LogF + FailNow

Which do you use?
- If you can let a test continue to run use Error/Errorf
- If a test is completely over and running further won't help at all, use Fatal/Fatalf

If not using subtests, Fatal will prevent other tests in that function from running. 

## Writing useful error messages 
- The goal of an error message is to quickly identify what went wrong in a given test and make it easier for a developer find, debug and fix 
- It should then look something like this 

t.Errorf("SomeFunc(%v) err = %v; want %v")

Common error sequence if 
if got != want
if actual != expected
