# graphql-test

Testing out union error handling.

```gql
query FindTodos {
  todo(id: "1") {
    ... on Todo {
      id
      text
    }
    ... on Error {
      message
    }
    ... on TodoNotFoundError {
      message
      code
    }
  }
}
```
