# mkdotenv – Basic Architecture

This document explains the high‑level architecture of **mkdotenv**, focusing on how parsing, execution, and secret resolution work together.

---

## 1. High‑Level Flow

At a high level, mkdotenv processes specially formatted comments (e.g. `#mkdotenv(...)`) and resolves secrets based on those instructions.

The flow is:

1. **Input**: A line of text (typically a comment) and a set of runtime arguments
2. **Parsing**: The line is parsed into a structured command (`MkDotenvCommand`)
3. **Execution**: The command is executed by an `Executor`
4. **Resolution**: A concrete secret resolver fetches the requested secret
5. **Output**: The resolved value (string) or an error

Each step is isolated behind clear interfaces to keep responsibilities separate.

---

## 2. Core Components

### 2.1 Parser (`core/parser`)

The parser is responsible for **interpreting mkdotenv comments** and converting them into a structured command object.

#### Key Type: `MkDotenvCommand`

```go
type MkDotenvCommand struct {
    Environment        string
    SecretResolverType string
    SecretPath         string
    Params             map[string]string
    Item               string
}
```

This struct represents *everything needed* to resolve a secret:

* **Environment** – Logical environment name (defaults to `default`)
* **SecretResolverType** – Resolver identifier (e.g. `keepassx`, `plain`)
* **SecretPath** – Path to the secret in the backend
* **Params** – Resolver‑specific parameters (file, password, etc.)
* **Item** – Optional sub‑item (e.g. a field inside a secret)

#### Parsing Logic

`ParseMkDotenvComment`:

* Uses a regular expression to validate and extract components from a comment
* Supports optional environments and item access (`.item`)
* Parses resolver parameters as key/value pairs
* Allows argument substitution using `$_ARG[name]`

If the line does not match the expected format, parsing fails gracefully by returning `nil`.

---

### 2.2 Executor (`executor`)

The executor is responsible for **orchestrating secret resolution** based on a parsed command.

#### Executor Interface

```go
type Executor interface {
    Execute(command *parser.MkDotenvCommand) (string, error)
}
```

This interface abstracts execution logic, making it easy to introduce alternative executors later (e.g. caching, logging, dry‑run).

#### Default Implementation: `CommandExecutor`

`CommandExecutor` performs three main tasks:

1. **Resolver selection**
2. **Resolver initialization**
3. **Secret resolution**

##### Resolver Selection

```go
switch command.SecretResolverType {
case "keepassx":
    resolver, err = secret.NewKeepassXResolver(...)
case "plain":
    resolver = secret.NewPlaintextResolver()
default:
    return "", fmt.Errorf("resolver %s not found", ...)
}
```

The executor maps the resolver type string to a concrete implementation from the `secret` package.

##### Resolution Strategy

* If `Item` is provided → `ResolveWithParam`
* Otherwise → `Resolve`

This allows uniform handling of simple secrets and structured secrets.

---

### 2.3 Secret Resolvers (`secret`)

Secret resolvers encapsulate **how and where secrets are fetched**.

They are hidden behind a common interface:

```go
type Resolver interface {
    Resolve(path string) (string, error)
    ResolveWithParam(path, param string) (string, error)
}
```

This design allows:

* Multiple secret backends (KeepassX, plaintext, future vaults)
* Executor logic to remain backend‑agnostic
* Easy extension by adding new resolver implementations

---

## 3. Design Principles

### Separation of Concerns

* **Parser**: Syntax and structure only
* **Executor**: Coordination and control flow
* **Resolvers**: Backend‑specific secret access

Each layer knows *just enough* about the next layer to do its job.

### Extensibility

* New resolvers can be added without touching the parser
* New executors can be added without changing resolver implementations
* Resolver parameters are passed generically via `map[string]string`

### Fail Fast

* Invalid syntax → parsing returns `nil`
* Unknown resolvers → explicit error
* Resolver initialization errors propagate immediately

---

## 4. Summary

mkdotenv is structured as a simple but extensible pipeline:

```
Comment → Parser → Command → Executor → Resolver → Value
```

This architecture keeps the system easy to reason about, easy to extend, and safe to evolve as new secret backends or execution strategies are introduced.
