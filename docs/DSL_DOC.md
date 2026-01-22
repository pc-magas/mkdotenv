# Mkdotenv Description Language

The **Mkdotenv Description Language** defines a standard way to annotate entries in a `.env` file so that mkdotenv tool can understand how each value should be resolved or generated.

Each annotation is expressed as a comment using the following format:

```
#mkdotenv(^environment^):resolve(^path^):^resolver^(^args^).^item^
```

## Components

- **environment**  
  The environment name (e.g., `development`, `production`) to which this rule applies. If not value `default` is assumed, `*` indicates any environment

- **path**  
  A filesystem or project-relative path that indicates where the resolver should operate or obtain additional information.

- **resolver**  
  A secret resolver mechanism indication currently supported:
  * `keepassx` for a file oppened by keepassx password manager
  * `plain` for a plantext value.

- **args**  
  Arguments passed to the resolver, formatted according to the resolver’s requirements. the argument follow this format:
  ```
   argument=value
  ```

- **item**  
  The specific key or sub-value within the resolved output that should be used to populate the corresponding environment variable. For example upon keppasx it is defined whether we need the password or the username of an entry.

# Magic Variables

This functionality is heavilty inspired from php. Currently the one supported is:

```
$_ARG[^argument^]
```

Where `^argument^` is an argument name rovided from user as cli argument.
