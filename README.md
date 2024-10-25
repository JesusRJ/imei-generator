# Gerador de IMEIs com Metadados

Este projeto foi desenvolvido para gerar números de IMEI (Identificação Internacional de Equipamento Móvel) válidos, além de fornecer metadados relacionados a dispositivos móveis, como marca, modelo, cor, memória e CPU. O gerador permite que o usuário forneça uma parte do número IMEI (TAC e serial parcial), e então gera IMEIs válidos usando o [algoritmo de Luhn](https://en.wikipedia.org/wiki/Luhn_algorithm) como verificação. O projeto é modular, podendo ser utilizado tanto como uma aplicação de linha de comando quanto como uma biblioteca em outros projetos.

## O que é um Código IMEI?

O [**IMEI (International Mobile Equipment Identity)**](https://en.wikipedia.org/wiki/International_Mobile_Equipment_Identity) é um número exclusivo que identifica dispositivos móveis em redes celulares. Ele possui 15 dígitos e é usado pelos operadores de redes para verificar a validade dos dispositivos e, se necessário, bloquear aparelhos roubados ou inválidos. O IMEI é dividido em três partes principais:

1. **TAC (Type Allocation Code)**: Composto pelos primeiros 8 dígitos, identifica o fabricante e o modelo do dispositivo.
2. **Número de Série**: Composto pelos 6 dígitos seguintes, que são únicos para cada dispositivo fabricado.
3. **Dígito Verificador**: O último dígito é calculado usando o algoritmo de Luhn, que valida o IMEI.

### Estrutura de um IMEI

Um número IMEI típico tem a seguinte estrutura:

```
123456 78 901234 5
|     |  |     | |
TAC   |  |     | Check Digit
      |  |     Serial Number
      |  |
      |  Manufacturer Code
      Model Code
```

## Como Executar o Projeto

Você precisará ter o [Go](https://golang.org/dl/) instalado no seu ambiente.

1. Clone este repositório:

   ```bash
   git clone https://github.com/seuusuario/gerador-imei.git
   cd gerador-imei
   ```

2. Execute o projeto diretamente:

   ```bash
   go run main.go
   ```

   Isso iniciará a aplicação de linha de comando, que solicitará ao usuário os primeiros 8 a 12 dígitos do IMEI e quantos números de IMEI deseja gerar.

### Exemplo de Execução

```bash
$ go run main.go
Enter the first 8 - 12 digits: 12345678
Enter the number of IMEI numbers to generate: 3

IMEI: 123456789012345
Brand: Apple
Model: iPhone 13
Color: Black
Memory: 128GB
CPU: A15 Bionic

IMEI: 123456789012346
Brand: Samsung
Model: Galaxy S22
Color: White
Memory: 128GB
CPU: Exynos 2200

IMEI: 123456789012347
Brand: Xiaomi
Model: Mi 11
Color: Gray
Memory: 128GB
CPU: Snapdragon 888
```

O programa gerará números IMEI válidos e associará cada um a um modelo de celular aleatório, fornecendo informações como a marca, cor, memória e CPU.

### Como Usar como Biblioteca

Se desejar utilizar este projeto como uma biblioteca em outro código Go, basta importar o pacote `imei` e usar a função `GenerateIMEIs`:

```bash
$ go get github.com/jesusrj/imei-generator
```

```go
package main

import (
    "fmt"
    "github.com/jesusrj/imei-generator/pkg/imei"
)

func main() {
    devices, err := imei.GenerateIMEIs("12345678", 5)
    if err != nil {
        fmt.Println("Erro:", err)
        return
    }

    for _, device := range devices {
        fmt.Printf("IMEI: %s\nBrand: %s\nModel: %s\n", device.IMEI, device.Brand, device.Model)
    }
}
```

Isso permitirá que você gere IMEIs e receba os metadados relacionados de forma fácil e modular.
