# Crossbuild for windows using linux

# Command:

```
make OS=windows
```

This would result a file named `mkdotenv_^VERSION^.exe` (replace `^VERSION^` with the version that exists upon `VERSION` file). 

# Testing windows execution upon Linux

## Preparation

Install the latest wine version, curwently it is tested using `wine-10.0` but later versions sould suffice. 
In case that wine's latest version is not upon repositories, and using Mint/Ubuntu/Debian install `winehq-stable` package following rthese instruction https://www.tecmint.com/install-wine-on-ubuntu-and-linux-mint/.

## Executing

Once wine is installed run:

```
wine ./mkdotenv_^VERSION^.exe
```

At command above replace `^VERSION^` with the version that is listed upon `VERSION` file.