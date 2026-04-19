# How to use Mkdotenv

## Generate Template .env file
First create a template file containing all secret definition per environment for example let us assume we have the file `.env.dist`:

```
#mkdotenv(dev):resolve("value"):plain()
#mkdotenv(prod):resolve("value"):kepassx(file="mysecretdb.kpbx",password="12334").PASSWORD
SECRET=
```

Then execute:

```
mkdotenv --environment prod
```

This would scan for `.env.dist` file and would look for any occurence for `prod` environment. 

Upon our example rhe one that should be resolved is the:
```
#mkdotenv(prod):resolve("value"):kepassx(file="mysecretdb.kpbx",password="12334").PASSWORD
```

The final result would be written upon `.env` file unless specified otherwise. Environment, template path and output path can be overriden via cli arguments.

### Command resolution sequence

In case for the same environment multiple occurences exiust before a variable the lastone is resolved for example if we have:

```
#mkdotenv(dev):resolve("value"):plain()
#mkdotenv(prod):resolve("value"):kepassx(file="mysecretdb.kpbx",password="12334").PASSWORD
#mkdotenv(dev):resolve("value"):kepassx(file="mysecretdb2.kpbx",password="12334").PASSWORD
#mkdotenv(prod):resolve("value"):kepassx(file="mysecretdb3.kpbx",password="12334").PASSWORD
SECRET=
```

And then execute:
```
mkdotenv --environment prod
```

The one resolved is:

```
#mkdotenv(prod):resolve("value"):kepassx(file="mysecretdb3.kpbx",password="12334").PASSWORD
```

Because it is the last occurenct for `prod` environment.

### Extra arguments:

Some parameters are not desired to be hardcoded upon `.env.dist` for these you can use the magic variable `$_ARG`. 
For example let assume that we have:

```
#mkdotenv(dev):resolve("value"):kepassx(file="mysecretdb2.kpbx",password=$_ARG[pass]).PASSWORD
SECRET=
```

We can provide the argument like this:

```
mkdotenv --environment=dev --argument pass=1234
```

Then the `$_ARG[pass]` it would be replaced with `1234`.

## Argument usage

Either run:

```
mkdotenv --help
```

Or if in linux (except alpine) run:

```
man mkdotenv
```