# latest

Return the latest version given a go module path.

```
$ ./latest github.com/go-git/go-billy
v1.0.0
```

```
$ ./latest github.com/go-git/go-billy/v4
v4.3.2
```

NOTE: This tool doesn't yet understand how to navigate vanity repos with a meta
tag.

