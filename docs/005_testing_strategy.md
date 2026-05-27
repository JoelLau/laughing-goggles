# Testing Strategy

## Goals

tests in this project have a few objetives:

1. give me confidence that code works as intended
1. document system behavior
1. (incidental) easier to put code under specific scenarios than to set it up
manually - e.g. scenarios involving existing accounts

## Test Coverage

as the goals above DO NOT include 100% test coverage, i WILL NOT be aiming for 100%
coverage because i believe:

1. it has diminishing returns on effort
1. some tests provide less than no value

## Strategy

taking everything else in this document into account, we will focus more on mimicing
"real" use cases and follow Kent C Dodd's testing trophy model,
where much of the focus is on integration testing.

real postgres instances will be used via [testcontainers](https://golang.testcontainers.org/modules/postgres/).

in the interest of time and keeping tests isolated,
a new postgres instance will be spun up per test.

(ideally, all tests will connect to the same postgres instance,
use migrations to create a template database, then each tests creates their
own postgres database with deterministic names)
