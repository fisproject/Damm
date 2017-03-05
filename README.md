# Damm

## Overview
A Go implementation of the Damm algorithm, a check digit algorithm created by H. Michael Damm. It detects all single-digit errors and adjacent transposition errors (swapping adjacent numbers).


## Example
```go
d, err := damm.Append("54994384990") // "549943849904"

result, err := damm.Verify(d) // true
```

## Licence
[MIT](http://opensource.org/licenses/MIT)
