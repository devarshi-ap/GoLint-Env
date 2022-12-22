# .env linter

---

##### 💭What is a .env file ?

- .env (or dotenv) files are simple text files containing a project’s env-vars.
  
- .env files have a simple key-value format
  
  ```textile
  FOO=BAR
  ```
  

##### 🚨The Problem :

- The human element in manually managing .env files can easily lead to typos and misconfiguration errors, posing a risk to production stability.

##### 💡Proposed Solution :

- ***`golint-env`*** - checks all `.env` files in the current dir:
  
  ```bash
  $ golint-env
  
  Checking .env
  .env:<line> <error>: <error-msg>
  .env:3 UnorderedKey: The BAR key should go before the FOO key
  
  Checking .env.test
  .env.test:1 LeadingCharacter: Invalid leading character detected
  
  Found 3 problems
  ```
  
  Pre-commit .env linter that checks for:
  
- ***✅ Duplicate Keys*** - detects if a key is not unique:
  
  ```textile
  ❌ Wrong
  FOO=BAR
  FOO=BAR
  
  ✅ Correct
  FOO=BAR
  BAR=FOO
  ```
  
- ***✅ Valueless Keys*** - detects if a line has a key without a value:
  
  ```textile
  ❌ Wrong
  FOO
  FOO=
  
  ✅ Correct
  FOO=BAR
  ```
  
- ***✅ Incorrect Delimiter*** - detects if key doesn't use an _ to separate words:
  
  ```textile
  ❌ Wrong
  FOO-BAR=FOOBAR
  
  ✅ Correct
  FOO_BAR=FOOBAR
  ```
  
- ***✅ Incorrect Leading Character*** - detects if line starts w/ unallowed char [^A-Z__]:
  
  ```textile
  ❌ Wrong
   FOO=BAR
  .FOO=BAR
  *FOO=BAR
  1FOO=BAR
  
  ✅ Correct
  FOO=BAR
  _FOO=BAR
  ```
  
- ***✅ Lowercase Keys*** - detects if a key has lowercase characters:
  
  ```textile
  ❌ Wrong
  FOo_BAR=FOOBAR
  foo_bar=FOOBAR
  
  ✅ Correct
  FOO_BAR=FOOBAR
  ```
  
- ***Invalid Quotes*** - detects if a value has non-matching quotes (if any):
  
  ```textile
  ❌ Wrong
  FOO=BAR BAR
  FOO="BAR'
  FOO='B"AR'
  
  ✅ Correct
  FOO=BAR
  FOO="BAR BAR"
  FOO='BAR BAR'
  ```
  
- ***✅ Incorrect Spaces*** - detects lines with whitespaces around equal sign character:
  
  ```textile
  ❌ Wrong
  FOO =BAR
  FOO= BAR
  FOO = BAR
  
  ✅ Correct
  FOO=BAR
  ```
  
- ***Invalid Substitution Key*** - detects invalid substitutions on key assignments:
  
  ```textile
  ❌ Wrong
  ABC=${BAR
  FOO="$BAR}"
  
  ✅ Correct
  ABC=${BAR}
  FOO="$BAR"
  ```
  
- ***✅ Incorrectly-Spaced Comments*** - comments should be preceded with a space:
  
  ```textile
  ❌ Wrong
  #comment
  FOO=BAR
  
  ✅ Correct
  # comment
  FOO=BAR
  ```
  
- ***(✅ comments) Trailing Whitespace*** - detects if a line has trailing whitespace
  
  ```textile
  ❌ Wrong
  FOO=BAR<whitespace>
  
  ✅ Correct
  FOO=BAR
  ```
  
- ***Unordered Key*** - detects if a key is not alphabetically ordered:
  
  ```textile
  ❌ Wrong
  FOO=BAR
  BAR=FOO
  
  ✅ Correct
  BAR=FOO
  FOO=BAR
  ```
  
- ***✅ Empty Key*** - detects if a key is empty
  
  ```textile
  ❌ Wrong
  =BAR
   =FOO
  
  ✅ Correct
  FOO=BAR
  F=BAR
  ```
  

***`golint-env compare`*** - compare 2 `.env` files w/ each other & output difference

```bash
$ golint-env compare .env .env.example

Comparing .env
Comparing .env.example
.env is missing keys: BAR
.env.example is missing keys: FOO
```