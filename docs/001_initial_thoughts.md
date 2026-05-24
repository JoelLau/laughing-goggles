# Initial Thoughts

> maintain data integrity and consistency ...

this is a simple banking rest api that i've built a number of times by now.
the key design decision lies with the database and how it uses the
event sourcing pattern and ledger.

> ... clean code principles is crucial ...

notable principls:

- [KISS] keep it simple, stupid
  - build the simplest structure
- [YAGNI] you aren't going to need it
  - only build what is immediately necessary
- [SRP] single responsibility principle
  - code should only do one thing
- unit tests

> ... well-documented

forms of documentation:

- behaviour: automated tests (unit and integration tests)
- REST endpoints: openapi 3
- database schemas: migrations, `COMMENT ON` statements
