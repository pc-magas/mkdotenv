# Creating new Secret Resolver

## Step 1: Implementation

Upon folder `mkdotenv/secret` each secret resolver has its own dedicated file. Each resolver should implement the following Interface:

```go
type Resolver interface {
    Resolve(path string) (string, error)
    ResolveWithParam(path, param string) (string, error)
}
```

The `Resolve` method resolves the secret `path` the returned values are:
* `string` is the actual secret
* `error` the error in case secret cannot be resoilved.

Each secret resolver may need to access a specific field. For example a `keepassx` secret contains these fields:

* `PASSWORD`
* `USERNAME`
* `NOTES`
etc etc

In that case the `ResolveWithParam` is used where upon `param` the nessesary field is provided.

## Step2: Enhance Command Executor

Upon `mkdotenv/core/executor/command_executor.go` upon the `Execute` function place the nessesary logic in order to initialize and call the executor.

The executor itself executes the following command that follows one of the following patterns:

```
#mkdotenv(^environment^):resolve(^secret_path^):^secret_resolver^(^resolver_args^).^item^
#mkdotenv(^environment^):resolve(^secret_path^):^secret_resolver^(^resolver_args^)
```

Where:

* `^environment^`–  is the environment where the secret should resolve upon.
* `^secret_path^`– the resolver used to retrieve the secret
* `^secret_resolver^`–  the resolver used to retrieve the secret 
* `^resolver_args^`– arguments required to initialize the resolver
*  `^item^` (optional) – the specific field of the secret; if provided, this value is passed as the param argument to `ResolveWithParam`
