# Release procedute


## STEP 1 Bump version

```
nano VERSION
bash ./tools/bump_version.sh
```

## STEP 2 push into dev

```
git push origin dev
```

## Step 3 merge and push into master

```
git checkout master
git merge dev
git push origin master
```

# Step 4 release into alpine:

See: https://github.com/pc-magas/mkdotenv-alpine-staging

# STEP 5 Release into ARCH

https://aur.archlinux.org/packages/mkdotenv
