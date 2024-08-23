# cpfcnpj

`cpfcnpj` is a Go package for validating Brazilian CPF and CNPJ numbers. The package provides utilities to clean up input strings and validate the structure and digits of CPF (Cadastro de Pessoas Físicas) and CNPJ (Cadastro Nacional da Pessoa Jurídica) numbers.

## Installation

To use this package in your Go project, install it using:

```bash
go get github.com/itispx/cpfcnpj
```

# Usage

# General Validation

If you need to validate either CPF or CNPJ without knowing beforehand which one it is, use the `Validate()` function in the main package. This function automatically detects and validates CPF or CNPJ based on the length of the input.

```go
package main

import (
	"fmt"
	"github.com/itispx/cpfcnpj"
)

func main() {
	document := "12.345.678/0001-95" // Can be CPF or CNPJ
	isValid := cpfcnpj.Validate(document)
	fmt.Printf("Is document valid? %v\n", isValid)
}
```

### Validate CPF

To validate a CPF number, use the `cpfvalidator.Validate()` function. This function expects a string and returns a boolean indicating whether the CPF is valid.

```go
package main

import (
	"fmt"
	"github.com/itispx/cpfcnpj/cpfvalidator"
)

func main() {
	cpf := "123.456.789-09"
	isValid := cpfvalidator.Validate(cpf)
	fmt.Printf("Is CPF valid? %v\n", isValid)
}
```

# Validating CNPJ

To validate a CNPJ number, use the `cnpjvalidator.Validate()` function. This function expects a string and returns a boolean indicating whether the CNPJ is valid.

```go
package main

import (
	"fmt"
	"github.com/itispx/cpfcnpj/cnpjvalidator"
)

func main() {
	cnpj := "12.345.678/0001-95"
	isValid := cnpjvalidator.Validate(cnpj)
	fmt.Printf("Is CNPJ valid? %v\n", isValid)
}
```

# Testing

To run the tests for this package, use the go test command:

```bash
go test ./...
```

This will run all the unit tests for the package and its components.

# License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.
