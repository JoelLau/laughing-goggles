# REST APIs

## Documentation

### OpenAPI

OpenAPI has been the leading form of REST endpoint documentation for the decade.
The main painpoint with it comes from maintaining documentation as changes are made
to the underlying implementation.

### OAPI Codegen

By using [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen/), we can make
the OpenAPI spec the source of truth and ensure that endpoints stays true to what
is defined in the doc.

This also speeds up development to an extent because developers simply need to
implement the interfaces that are generated.

In combination with httptest, this makes for a strong development workflow for TDD
because tests that compile can be written first.

### Accessing docs

While the documentation easily accesible via local tools like vscode's
[Swagger Viewer](https://marketplace.visualstudio.com/items?itemName=Arjun.swagger-viewer)
extension, my personal preference is for it to be hosted alongside the other endpoints.

example endpoint design:

```plaintext

GET /api/swagger

GET /api/livez

GET /api/v1/transactions
```
