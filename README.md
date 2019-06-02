# DataGen

A local HTTP API that returns the requested permutations of an array or string.

## Usage

go run datagen/cmd/httpapi/main.go

## HTTP API

```
localhost:80?funcid=slice&dimensions=4,4&valid_values=0,1&permutation_range=0,100
localhost:80?funcid=string&length=4&ascii_range=60,90&permutation_range=0,100
```

permutation_range is optional but reccomended for strings as it gets rather large.
leave ascii_range empty for all lowercase and uppercase characters

## Genlib API

(DOCUMENTATION TODO)
