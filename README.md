# cpfcnpj

`cpfcnpj` é um pacote Go para validação, formatação e geração de números de CPF e CNPJ brasileiros. O pacote oferece utilitários para limpar strings de entrada, validar a estrutura e dígitos de CPF (Cadastro de Pessoas Físicas) e CNPJ (Cadastro Nacional da Pessoa Jurídica), aplicar máscaras de formatação e gerar números válidos.

# Instalação

Para usar este pacote no seu projeto Go, instale-o usando:

```bash
go get github.com/itispx/cpfcnpj
```

# Uso

## Validação Geral

Se você precisa validar CPF ou CNPJ sem saber previamente qual é, use a função `Validate()` do pacote principal. Esta função detecta e valida automaticamente CPF ou CNPJ com base no comprimento da entrada e retorna um booleano indicando a validade junto com uma string especificando o tipo de documento ("cpf" ou "cnpj").

```go
package main

import (
	"fmt"
	"github.com/itispx/cpfcnpj"
)

func main() {
	documento := "12.345.678/0001-95" // Pode ser CPF ou CNPJ
	valido, tipoDoc := cpfcnpj.Validate(&cpfcnpj.ValidateInput{
		CPFCNPJ: documento,
	})
	fmt.Printf("O documento é válido? %v, Tipo de documento: %s\n", valido, tipoDoc)
}
```

## Validar CPF

Para validar um número de CPF, use a função `ValidateCPF()`. Esta função espera um `ValidateCPFInput` e retorna um booleano indicando se o CPF é válido.

```go
package main

import (
	"fmt"
	"github.com/itispx/cpfcnpj"
)

func main() {
	cpf := "123.456.789-09"
	valido := cpfcnpj.ValidateCPF(&cpfcnpj.ValidateCPFInput{
		CPF: cpf,
	})
	fmt.Printf("O CPF é válido? %v\n", valido)
}
```

## Validar CNPJ

Para validar um número de CNPJ, use a função `ValidateCNPJ()`. Esta função espera um `ValidateCNPJInput` e retorna um booleano indicando se o CNPJ é válido.

```go
package main

import (
	"fmt"
	"github.com/itispx/cpfcnpj"
)

func main() {
	cnpj := "12.345.678/0001-95"
	valido := cpfcnpj.ValidateCNPJ(&cpfcnpj.ValidateCNPJInput{
		CNPJ: cnpj,
	})
	fmt.Printf("O CNPJ é válido? %v\n", valido)
}
```

## Aplicar Máscara

Para formatar um CPF ou CNPJ com a máscara apropriada, use a função `Mask()`. Esta função detecta automaticamente o tipo de documento com base no comprimento e aplica a formatação correspondente.

```go
package main

import (
	"fmt"
	"github.com/itispx/cpfcnpj"
)

func main() {
	documento := "12345678901" // CPF sem formatação
	documentoFormatado, err := cpfcnpj.Mask(&cpfcnpj.MaskInput{
		CPFCNPJ: documento,
	})
	if err != nil {
		fmt.Printf("Erro: %v\n", err)
		return
	}
	fmt.Printf("Documento formatado: %s\n", documentoFormatado) // Saída: 123.456.789-01
}
```

## Aplicar Máscara em CPF

Para formatar especificamente um CPF, use a função `MaskCPF()`:

```go
package main

import (
	"fmt"
	"github.com/itispx/cpfcnpj"
)

func main() {
	cpf := "12345678901"
	cpfFormatado, err := cpfcnpj.MaskCPF(&cpfcnpj.MaskCPFInput{
		CPF: cpf,
	})
	if err != nil {
		fmt.Printf("Erro: %v\n", err)
		return
	}
	fmt.Printf("CPF formatado: %s\n", cpfFormatado) // Saída: 123.456.789-01
}
```

## Aplicar Máscara em CNPJ

Para formatar especificamente um CNPJ, use a função `MaskCNPJ()`:

```go
package main

import (
	"fmt"
	"github.com/itispx/cpfcnpj"
)

func main() {
	cnpj := "12345678000195"
	cnpjFormatado, err := cpfcnpj.MaskCNPJ(&cpfcnpj.MaskCNPJInput{
		CNPJ: cnpj,
	})
	if err != nil {
		fmt.Printf("Erro: %v\n", err)
		return
	}
	fmt.Printf("CNPJ formatado: %s\n", cnpjFormatado) // Saída: 12.345.678/0001-95
}
```

## Gerar CPF

Para gerar um número de CPF válido, use a função `GenerateCPF()`. Você pode optar por receber o CPF formatado ou não:

```go
package main

import (
	"fmt"
	"github.com/itispx/cpfcnpj"
)

func main() {
	// Gerar CPF sem formatação
	cpf := cpfcnpj.GenerateCPF(&cpfcnpj.GenerateCPFInput{
		Formatted: false,
	})
	fmt.Printf("CPF gerado: %s\n", cpf) // Exemplo: 12345678901

	// Gerar CPF formatado
	cpfFormatado := cpfcnpj.GenerateCPF(&cpfcnpj.GenerateCPFInput{
		Formatted: true,
	})
	fmt.Printf("CPF formatado gerado: %s\n", cpfFormatado) // Exemplo: 123.456.789-01
}
```

## Gerar CNPJ

Para gerar um número de CNPJ válido, use a função `GenerateCNPJ()`. Você pode optar por receber o CNPJ formatado ou não:

```go
package main

import (
	"fmt"
	"github.com/itispx/cpfcnpj"
)

func main() {
	// Gerar CNPJ sem formatação
	cnpj := cpfcnpj.GenerateCNPJ(&cpfcnpj.GenerateCNPJInput{
		Formatted: false,
	})
	fmt.Printf("CNPJ gerado: %s\n", cnpj) // Exemplo: 12345678000195

	// Gerar CNPJ formatado
	cnpjFormatado := cpfcnpj.GenerateCNPJ(&cpfcnpj.GenerateCNPJInput{
		Formatted: true,
	})
	fmt.Printf("CNPJ formatado gerado: %s\n", cnpjFormatado) // Exemplo: 12.345.678/0001-95
}
```

## Função de Limpeza

O pacote também oferece uma função `Clean()` que remove todos os caracteres não numéricos de uma string:

```go
package main

import (
	"fmt"
	"github.com/itispx/cpfcnpj"
)

func main() {
	documentoSujo := "123.456.789-01"
	documentoLimpo := cpfcnpj.Clean(documentoSujo)
	fmt.Printf("Documento limpo: %s\n", documentoLimpo) // Saída: 12345678901
}
```

## Constantes

O pacote define as seguintes constantes para padrões de formatação:

- `CPFFormatPattern`: Padrão regex para formatação de CPF: `([\d]{3})([\d]{3})([\d]{3})([\d]{2})`
- `CNPJFormatPattern`: Padrão regex para formatação de CNPJ: `([\d]{2})([\d]{3})([\d]{3})([\d]{4})([\d]{2})`

## Estrutura do Pacote

O pacote está organizado nos seguintes arquivos:

- `main.go`: Funções gerais de validação e formatação para CPF/CNPJ
- `cpf.go`: Funções específicas para CPF (validação, formatação e geração)
- `cnpj.go`: Funções específicas para CNPJ (validação, formatação e geração)
- `strutil.go`: Funções utilitárias para manipulação de strings
- `constants.go`: Constantes do pacote

## Testes

Para executar os testes deste pacote, use o comando go test:

```bash
go test
```

Isso executará todos os testes unitários do pacote e seus componentes.

## Licença

Este projeto está licenciado sob a Licença MIT - consulte o arquivo [LICENSE](LICENSE) para detalhes.
