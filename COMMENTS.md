## 1 - Intervalo de Números

### Histórico
  - Primeira versão para validar o raciocionio
  - Inclusão de testes
  - Refatoração

### Validações
  

### Melhorias
  - Aceitar uma lista desordenada, ordenar e agrupar os intervalos


## 2 - Troco

### Histórico

- Fiz um código usando estrutura de repetição depois alterei para recursão
- Para conversão da moeda no poderiamos usar o [currency](https://godoc.org/golang.org/x/text/currency#example-Query), mas optei em fazer o replace ",", "." para evitar usar outros pacotes.
- Para deixar o código testável eu precisei refator, fazer teste pegando o print no console não acho legal. 
  - Acabei criando mais tipagens para poder fazer map de troco mais legível e testavel
- Criei um método só para printar o recibo
- Precisei de um work around, ao criar o teste onde tinha que dividir o 14.20 /10.0 retorna 4.1999999999999, então precisei criar o método formatFloat 

### Validações
 - Inclui validação quando os valores digitados estão incorretos
 - A momentos em que o troco pode ser inferir a 0.05 que a menor fração, neste caso assumi que o retorno é zero. Não temos bala! =D

### Melhorias

- Conversão da moeda para o BRL usando [currency](https://godoc.org/golang.org/x/text/currency#example-Query).
- Criar uma lista de moeda conforme a type Unit struct. No caso BRL.
  
    Additional common currencies as defined by CLDR.
    BRL Unit = Unit{brl}
- Fazer a quebra de caixa com o arredamento quando houver troco inferior ao menor valor da lista de moedas.
