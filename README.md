# DataGen

A HTTP API that returns the requested permutations of a custom data type. Currently supports integer, string, and complex numbers as well as their array variations(Up to 4 dimensions).

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
http://localhost:8080/?key=c_slice,i_slice&c_slice_type=complex_slice&c_slice_dimensions=2&c_slice_valid_values_index=2,2&c_slice_real_values=0,1,0,2&c_slice_imaginary_values=1,2,3,4&i_slice_type=int_slice&i_slice_dimensions=2,2&i_slice_valid_values=0,1,0,1,0,1,0,1&i_slice_valid_values_index=2,2,2,2

# string slice and string
http://localhost:8080/?key=s_slice,str&s_slice_type=string_slice&s_slice_dimensions=2,2&s_slice_lengths=1,2,1,2&s_slice_string_values=a,b,a,b,c,a,b,a,b,c,a,b,a,b,c,a,b,a,b,c,a,b&s_slice_string_values_index=2,3,2,3,2,3,2,3,2&str_type=string&str_length=4&str_string_values=a,b,c,a,b,c,d,a,b,a&str_string_values_index=3,4,2,1
```
