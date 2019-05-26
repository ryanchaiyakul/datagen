# DataGen

A local HTTP API that returns the requested permutations of an array or string.

## Usage

go run datagen/cmd/httpapi/main.go

## HTTP API

```
localhost:8080?funcid=array&dimensions=4,4&valid_values=0,1&permutation_range=0,100
localhost:8080?funcid=string&length=4&permutation_range=0,100
```

**permutation_range is optional but reccomended for strings as it gets rather large.**
