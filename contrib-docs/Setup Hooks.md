# Setup Local git hooks

Upon `./tools/git-hooks` there are scripts that could be copied  upon `.git/hooks`.
Upon this folder it should contain scripts named as git-hooks expects (https://git-scm.com/book/en/v2/Customizing-Git-Git-Hooks).

In order to install them run:

```
bash ./tools/setup-hooks.sh
```

That scripts symlinks any script that is residing inside `./tools/git-hooks` upon `.git/hooks`.