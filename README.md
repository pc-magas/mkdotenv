
```text
 __  __ _    _____        _                  
|  \/  | |  |  __ \      | |                 
| \  / | | _| |  | | ___ | |_ ___ _ ____   __
| |\/| | |/ / |  | |/ _ \| __/ _ \ '_ \ \ / /
| |  | |   <| |__| | (_) | ||  __/ | | \ V / 
|_|  |_|_|\_\_____/ \___/ \__\___|_| |_|\_/  
```
                                              
**Simplify Your .env Files – One Variable at a Time!**

MkDotenv is a lightweight and efficient tool for managing your `.env` files. Whether you're adding, updating, or replacing environment variables, MkDotenv makes it easy and hassle-free.

# Usage

## Basic

```
mkdotenv <variable_name> <variable_value>
```

This will output  to *stdout* the contents of a `.env` file with the variable `<variable_name>` having `<variable_value>` instead of the original one.
If no `.env` file exists it will just output the `<variable_name>` having the `<variable_value>`.

### Example:

```
mkdotenv DB_HOST 127.0.0.1
```

This will output:

```
DB_HOST=127.0.0.1
```

If a .env file exists with values:

```
DB_HOST=example.com
DB_USER=xxx
```

The final output would be:

```
DB_HOST=127.0.0.1
DB_USER=xxx
```

## Selecting file to read and write upon

Instead of outputing the .env value you can use the `--output-file` argument in order to write the contents upon a file.
Also you can use the parameter `--input-file` in order to select which file to read upon, if ommited `.env` file is used.

### Example 1 Read a specified file and output its contents to *stdout*:

Assuming we run the command

```
mkdotenv DB_HOST 127.0.0.1 --input-file=.env.example
```

This will read the `.env.example` and output:

```
DB_HOST=127.0.0.1
```


### Example 2 Write file upon a .env file:

```
mkdotenv DB_HOST 127.0.0.1 --output-file=.env.production
```

This would **create** a file named `.env.production` containing:

```
DB_HOST=127.0.0.1
```

### Example 3 Read a specified .env file and output its contents to a seperate .env file:

Assuming we have a file named `.env.template` containing:

```
DB_HOST=example.com
DB_USER=xxx
DB_PASSWORD=zzz
```

And we want to create a file named `.env.production` containing 

```
DB_HOST=127.0.0.1
DB_USER=xxx
DB_PASSWORD=zzz
```

We have to run:

```
mkdotenv DB_HOST 127.0.0.1 --input-file .env.template --output-file .env.production
```

## Piping outputs

You can provide a .env via a pipe. A common use is to replace multiple variables:

```
mkdotenv DB_HOST 127.0.0.1 | mkdotenv DB_USER maiuser | mkdotenv DB_PASSWORD XXXX --output_file .env.production
```

