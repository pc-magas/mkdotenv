# Plaintext resolver

Just place a value as is. Value needs no resolution.


# Usage

```
#mdotenv():resolve(^value^):plain()
```

## Example

Asuming upon .env file we have:

```
#mkdotenv():resolve("Hello World"):plain()
VALUE=
```

The final result is:

```
#mkdotenv():resolve("Hello World"):plain()
VALUE="Hello World"
```
