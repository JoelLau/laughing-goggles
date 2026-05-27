# Concurrency / Performance

## Performance

Ledger based systems are often complex because current state has to be derived
from many events (event sourcing pattern).

For example, the account balance is currently derived from processing all the
entries an account has ever been involved in.

## Concurrency

To make the situation more complex, the likelihood of concurrency issues occurring
grows as the number of requests per seconds increases.

Picture a scenario where we perform simple checks for an account's balance before
attempting to deduct from it. 2 concurrent requests to perform deductions might
result in the account's balance going negative because both checks happened before
any deductions happen.

To combat this, we will require mutexes to "hold" the account's balance while
a transaction is being processed (e.g. postgres' SELECT .. FOR UPDATE).

However, by fixing one concurrency problem, we will introduce another.
Transactions that involve more than 1 account might cause deadlocks as they
might both wait for each other resource to be free-ed (unlocked).

Due to the complexity of deadlocks and concurrency (and this being a take home assignment
with extremely low requests per second), I have chosen to forego implementing
solutions concurrency and performance for this project.
