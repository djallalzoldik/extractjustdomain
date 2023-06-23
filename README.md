# extractjustdomain
This tool help you to extract just the domains from the list ,

## examples:
```
https://www.google.com/fadslf
https://help.facebook.com/fladsf
https://new.exapmle.com?ldf

```

```
cat list | ./extractjustdomain
```
## result:
```
google.com
facebook.com
example.com
```


## install
```
git clone https://github.com/djallalzoldik/extractjustdomain.git
cd extractjustdomain
go build
```
## use
```
cat list | ./extractjustdomain
```
