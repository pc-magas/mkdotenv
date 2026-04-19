# KeepassX Secret Resolver

The **KeepassX Secret Resolver** allows you to fetch secrets stored in a KeePassX (KDBX) database using the **MkDotenv Markup Language (MDML)**. It can extract fields like `USERNAME`, `PASSWORD`, `URL`, and `NOTES` from entries based on a specified key.

---

## MkDotenv Markup Language Definition

The general syntax for resolving a KeePassX secret is:

```
#mkdotenv()::resolve(^key^)::keepassx(file=^path_to_kpbx.file^,password=^password^).PASSWORD
#mkdotenv()::resolve(^key^)::keepassx(file=^path_to_kpbx.file^,password=^password^).USERNAME
#mkdotenv()::resolve(^key^)::keepassx(file=^path_to_kpbx.file^,password=^password^).URL
#mkdotenv()::resolve(^key^)::keepassx(file=^path_to_kpbx.file^,password=^password^).URL

```

---

### Syntax Components

| Component | Description |
|-----------|-------------|
| `#mkdotenv()` | Initializes the MkDotenv processing context. |
| `resolve(^key^)` | The path to the secret. Use slashes for nested groups. The final segment must be the Entry Title (e.g., "group/subgroup/entry"). |
| `keepassx(...)` | Specifies that the resolver should use KeePassX as the secret source. |
| `file=^kpbx.file^` | Path to the `.kdbx` KeePassX database file. |
| `password=^password^` | Master password for decrypting the KeePassX database. |


### Supported Fields

- **PASSWORD** – The password stored in the entry.  
- **USERNAME** – The username associated with the entry.  
- **URL** – The URL or associated link for the entry.  
- **NOTES** – Any notes attached to the entry.  


### Examples

**NOTES:**
In the examples bellow:

1. `1234` is a dummy password pleach choose your own.
2. Upon `mkdotenv()` you can pass the environment inside the `()`. it is ommited for simplicity.

Let me assume that inside the **mypassword.kdbx** file following records exist:

```
- databases 
-- db1
--- login1
```

The `login1` is our password entry.

**1. Fetch only the password:**

```
#mkdotenv()::resolve("databases/db1/login1")::keepassx(file="mypassword.kdbx",password="1234").PASSWORD
```

**2. Fetch username:**

```
#mkdotenv():resolve("databases/db1/login1"):keepassx(file="mypassword.kdbx",password="1234").USERNAME
```

### Pro-Tip: Dynamic Passwords

In order to avoid hardcoding the password you can `$_ARG` entries:

```
#mkdotenv():resolve("databases/db1/login1"):keepassx(file="mypassword.kdbx",password=$_ARG[password]).PASSWORD
```

Then you can place the password like this:

```
mkdotenv --arg password=1234
```

---

### Notes

- The `^key^` placeholder should uniquely identify the entry in the KeePassX database. Nested groups can be referenced using slashes, e.g., `^group1/subgroup1/entry1^`.  
- Ensure that the file path and password are kept secure and not hard-coded in production environments. Consider using environment variables for sensitive information.