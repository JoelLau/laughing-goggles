# Error Handling

## General Strategy

bubble errors up the stack before handling / logging.

wrap errors with more information at every step (`fmt.errorf`)

## REST EndPoints

avoid re-inventing the wheel; follow existing standards in RFC 9457

have a static, append-only wiki of event codes that can be referenced.

this will serve as documentation for both devs and users.

give each group exactly enough information they need to ease
support / operational burden.

_(NOTE: example links were added to server responses,
but they do not exist because this is a demo project)_
