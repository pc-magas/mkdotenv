# Mkdotenv Description Language (MDL)

The **Mkdotenv Description Language** provides a standardized way to annotate `.env` file entries so the `mkdotenv` tool can automatically resolve or generate environment variable values.

## Annotation Syntax

```
#mkdotenv(^environment^):resolve(^path^):^resolver^(^args^).^item^
```

### Components

1. **environment**  
   - Specifies the target environment (e.g., `development`, `production`) for which this rule applies.  
   - Default: `default` if omitted.  
   - `*` indicates the rule applies to all environments.

2. **path**  
   - A filesystem or project-relative path that guides where the resolver should operate or obtain additional data.

3. **resolver**  
   - The mechanism used to retrieve or generate the value. Currently supported:
     - `keepassx` → fetches secrets from a KeePassX database.
     - `plain` → uses plaintext values.

4. **args**  
   - Arguments passed to the resolver, formatted as key-value pairs:
     ```
     argument=value
     ```
   - These provide additional details the resolver may require, like file paths, entry names, or usernames.

5. **item**  
   - Specifies the exact key or sub-value from the resolved output to use.  
   - Example: For KeePassX, this could be `password` or `username`.

---

## Magic Variables

Inspired by PHP-style syntax. Currently supported:

```
$_ARG[^argument^]
```

- `^argument^` refers to a CLI argument provided by the user.
- Allows dynamic resolution based on runtime input.

---

## File Paths as Resolver Arguments

If a resolver requires a file path, it should be **relative to the template directory**.

*Example:*

Assuming we have upon a template file the following:

```
#mkdotenv():resolve(secret_path):myresolver(file="myfile.txt")
```

With project structure:

```
- templates
  |-- .env.dist
  |-- myfile.txt
```

Then `myfile.txt` should be resolved as `templates/myfile.txt`.
