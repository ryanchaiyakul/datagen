# DataGen

A HTTP API that returns the requested permutations of an array or string.

## Usage

In order to start the HTTP API run:

``` Bash
go run datagen/cmd/httpapi/main.go
```

## HTTP API

### Key Parameters

- **key** : 1D string slice that contains the keys of the requested data types
- **permutation_range** : 2 long int array that contains the lower and upper bounds of the permutations returned (lower : inclusive, upper : exclusive)

### Int Slice Parameters

- **dimensions** : 1D int slice that contains the lengths of each dimension
- **valid_values** : 1D int slice that contains all the information that forms the 2D ValidValues slice
- **valid_values_index** : Determines how to generate the 2D slice from valid_values. The length of *valid_values_index* should equal the length of the flattened slice requested.

### Int Parameters

- **valid_values** : 1D int slice that contains the possible values of the integer

### String Slice Parameters

- **dimensions** : 1D int slice that contains the lengths of each dimension
- **lengths** : 1D int slice that contains the length of each string in the requested slice
- **string_values** : The same as *valid_values* but contains strings instead of integers
- **string_values_index** : The sae as *valid_values_index* as a slice parameter

### String Parameters

- **length** : An integer equal to the length of the string requested
- **string_values** : The same as *valid_values* but contains strings instead of integers
- **string_values_index** : The sae as *valid_values_index* as a slice parameter

### Complex Slice Parameters

- **dimensions** : 1D int slice that contains the lengths of each dimension
- **real_values** : The *valid_values* of the real number
- **imaginary_values** : The *valid_values* of the imaginary number
  
### Complex Parameters

- **real_values** : The *valid_values* of the real number
- **imaginary_values** : The *valid_values* of the imaginary number

#### Examples

``` bash
# complex slice and 2D int slice
http://localhost:8080/?funcid=data&key=test,hello&test_type=complex_slice&test_dimensions=2&test_valid_values_index=2,2&test_real_values=0,1,0,2&test_imaginary_values=1,2,3,4&hello_type=int_slice&hello_dimensions=2,2&hello_valid_values=0,1,0,1,0,1,0,1&hello_valid_values_index=2,2,2,2
```
