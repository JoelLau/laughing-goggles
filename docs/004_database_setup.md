# Database Setup

## Migrations

keep things readable by using database migrations.
the migration files were edited during development for readability and
since there is no other users in any deployed environments (no prod-like)

similarly, no down migrations were generated since docker-compose was used.

## Database Access

raw driver (pgx) vs ORM (gorm) vs query builder (goqu) vs sqlc

raw driver - too tedious

ORM - extra layer provides little value for too much abstraction

query builder - choice of database is unlikely to change,
we'd also want database-specific optimizations / data types

✅ sqlc - feels the most "native" in that raw sql is written.
requires additional step in build process (code generation).
