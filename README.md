# DataGen

A HTTP API that returns the requested permutations of an array or string.

## Usage

In order to start the HTTP API run:

``` Bash
go run datagen/cmd/httpapi/main.go
```

## HTTP API

Key:

- **Bold** : A parameter that is necessary to pass in.
- *Italic* : A paramter that is passed in by different means.

### General Parameters

- **funcid** : Determines the data type that will be returned. Each funcid has different paramters.

### Slice Parameters

- **dimensions** : A one dimensional slice that contains the lengths of each dimension.
- *permutation_range* : A 2 long array that contains the bounds of the permutations requested. When using the same configuration, two different calls will output the same permutations but possibly in different order. Reccomended to be used when working with large slices or when parallelizing the workflow.
- *validValues* : A two dimensional slice containing the possible values according to each index in the dimensions. The length of validValues should equal the length of the flattened slice requested.

#### validValues

1. **valid_values_range** : A 2 long array that contains the bounds of the validValues(inclusive). This will apply to all indexs like valid_values_unanimous but less customizable because you can only submeit a single range.
2. **valid_values_unanimous** : Duplicates the given validValues so that each index will have the same validValues
3. *validValuesUnique* : For every index in the requested slice, there must be a one dimensional slice serving as its *validValues*

##### validValuesUnique

- **valid_values** : Contains all the actually information that forms the 2D ValidValues slice
- **valid_values_index** : Determines how to generate the 2D slice from valid_values. The length of *valid_values_index* should equal the length of the flattened slice requested.

#### Slice Examples

``` Bash
# using valid_values_range
localhost:8080?funcid=slice&dimensions=2,2&valid_values_range=0,2

# using valid_values_unanimous
localhost:8080?funcid=slice&dimensions=2,2&valid_values_unanimous=0,1,2

# using valid_values_unique
localhost:8080?funcid=slice&dimensions=2,2&valid_values=0,1,2,0,1,2,0,1,2,0,1,2&valid_values_index=3,3,3,3

# using valid_values_unique for control
localhost:8080?funcid=slice&dimensions=2,2&valid_values=0,2,1,2,3,5,4,3,2,0&valid_values_index=1,2,3,4
```

### String Parameters

- **length** : An integer equal to the length of the string requested
- **permutation_range** : Acts the same as *permutaiton_range* as a slice pararmter
- *stringValues* : The *validValues* of a string.

#### stringValues

ascii_values follows the same idea as valid_values, but the resulting slice is converted to a string.

1. **string_values_unanimous**
2. *string_values_unique*

##### stringValuesUnique

1. **string_values**
2. **string_values_index**

#### String Examples

``` Bash

# using ascii_values_unanimous
localhost:8080?funcid=string&length=2&string_values_unanimous=a,b

# using ascii_values_unique
localhost:8080?funcid=sring&length=2&string_values=a,b&string_values_index=1,1

# using ascii_values_unique for control
localhost:8080?funcid=string&length=2&string_values=a,a,b&string_values_index=1,2
```
