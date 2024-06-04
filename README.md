# inscope

A binary to check if a list of DNS is in a scope of IP.

```
cat domains.txt
shoxxdj.fr
github.com

cat scope.txt
192.168.1.0/24
10.10.10.10

inscope -domains domains.txt -scope scope.txt
```

If URL is linked to one IP defined in scope.txt url will be printed.

A verbose output can be obtained with -full parameter.
