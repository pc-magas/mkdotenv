
First run:
```
bash runAlpine.sh
```

Then run:

```
su packager
abuild-keygen -a -i
```

Afterwards run:

```
abuild checksum
abuild -r
```
