# Desafio Técnico para Seleção Globo.com
# Desenvolvedor BackEnd - Integração

## Considerações Gerais

Você deverá usar este repositório como o principal do projeto. Todos os seus commits devem estar registrados aqui.

**Registre tudo**: Ideias que gostaria de implementar se tivesse mais tempo (explique como você as resolveria), decisões tomadas e seus porquês, arquiteturas testadas e os motivos de terem sido modificadas ou abandonadas.
Use arquivos para descrever suas intensões na raiz do repositório, como COMMENTS.md ou HISTORY.md.

A escolha da linguagem é deixada para você, utilize **a que você se sente mais confortável**. A entrada deverá ser por `STDIN` (*standard input*) e a saída por `STDOUT` (*standard output*) na linguagem que você escolher. Sinta-se livre para incluir ferramentas e bibliotecas open source.

Forneça as instruções de instalação e execução do seu sistema, observaremos **principalmente seu *design* de código**. Aspectos como coesão, baixo acoplamento, testabilidade e legibilidade são os principais pontos.

Devemos ser capazes de executar o código do seu backend no nosso ambiente local
ou em uma VM ou máquina limpa com os seguintes comandos, ou algo similar (faça
um INSTALL.md com as instruções):

    git clone seu-repositorio
    cd seu-repositorio
    ./configure (ou algo similar)
    make run (ou algo similar)

Escolha ao menos 3 dos desafios abaixo para resolver de acordo com o seu perfil. Quanto mais dados tivermos para analisar, melhor para o candidato.
Caso já tenha participado do processo seletivo, por favor escolha um desafio diferente do que foi feito anteriormente.


## 1 - Intervalo de Números

Dado uma lista de números inteiros, agrupe a lista em um conjunto de intervalos

**Exemplo**:

    Entrada: 100, 101, 102, 103, 104, 105, 110, 111, 113, 114, 115, 150
    Saída: [100-105], [110-111], [113-115], [150]



## 2 - Troco

Funcionários de empresas comerciais que trabalham como caixa tem uma grande responsabilidade em suas mãos. A maior parte do tempo de seu expediente de trabalho é gasto recebendo valores de clientes e, em alguns casos, fornecendo troco.

Seu desafio é fazer um programa que **leia o valor total** a ser pago e o **valor efetivamente pago**, respondendo o **valor do troco** e o **menor número de cédulas e moedas** que devem ser fornecidas.

Deve-se considerar que há:

    cédulas de R$ 100,00, R$ 50,00, R$ 20,00, R$ 10,00, R$ 5,00 e R$ 2,00;
    moedas de R$ 1,00, R$ 0,50, R$ 0,25, R$ 0,10, R$ 0,05.


**Exemplos**:

    1. Valor a ser pago: R$ 25,00
    2.  Valor efetivamente pago: R$ 40,00
    3. Troco: R$ 15,00
        - 1 x R$ 10,00
        - 1 x R$ 5,00


    1. Valor a ser pago: R$ 35,80
    2. Valor efetivamente pago: R$ 50,00
    3. Troco: R$ 14,20
        - 1 x R$ 10,00
        - 2 x R$ 2,00
        - 2 x R$ 0,10


    1. Valor a ser pago: R$ 40,00
    2. Valor efetivamente pago: R$ 40,00
    3. Troco: R$ 0


## 3 - Breath of Fantasy

Crie um jogo baseado em turnos onde dois personagens lutam entre si. Cada personagem tem `nome`, `pontos de energia` e `pontos de poder`.
Os pontos de energia e poder são valores numéricos inteiros. 

Por exemplo, no **primeiro turno** o `herói` (o atacante da vez) ataca
o `inimigo` (o defensor da vez) o inimigo terá seus pontos de energia diminuídos, no **segundo turno** o `inimigo` vira o atacante e o `herói` se transforma no defensor.

O dano sofrido, ou seja, os pontos de energia perdidos pelo inimigo, dependem do `fator sorte`. O fator sorte é um `número aleatório de 0 a 100` que é gerado a cada turno da batalha.
Há quatro tipos de ataques que dependem logicamente do fator sorte.

### Fator sorte

* Quando a sorte está entre 0 e 15 o ataque é perdido e não causa dano, imprimindo **"Errou - 0 HP"**
* Quando a sorte está entre 16 e 70 o ataque é normal e causa 1/3 dos pontos de poder do atacante em dano, imprimindo **"Normal - X HP"**
* Quando a sorte está entre 71 e 96 o ataque é sorte e causa 20% a mais do que o ataque normal, imprimindo **"Sorte!!! - X HP"**
* Quando a sorte está entre 97 e 100 o ataque é crítico e causa o dobro de um ataque normal, imprimindo **"Crítico! - X HP"**

**X indica o valor de dano sofrido.**

O jogo segue o esquema de turnos, onde a cada turno um jogador ataca o outro. Ao fim de cada turno os papéis de atacante e defensor se alternam. O jogo acaba quando um dos personagens não tem mais energia para lutar.

### Entrada

Cada personagem será informado usando o seguinte padrão:  `nome energia poder`.

```
Entre a primeira personagem:
nome1 40 50
Entre a segunda personagem:
nome2 50 40
```

### Saída

O jogo deverá produzir as seguintes saídas:

```
O jogo começou
Batalha entre nome1 e nome2
nome1 atacou nome2
<mensagem de dano>
nome2 atacou nome1
<mensagem de dano>
...
Jogo acabou, o vencedor foi nome1 com HP restante de Y
```


## 4 - Conversão de medidas

Você terá que criar um sistema que **formate a saída** de uma **dada entrada** para as três unidades de dados: **tempo, bytes e porcentagem**.

### Regras para formatação

#### Tempo

A entrada provida deve ser no formato `<numero> <ms>`, ex: `1340 ms`. A unidade usada para entrada será milisegundo.
* Quando a entrada for menor que 1 segundo a saída deve continuar em milisegundos `ms`.
* Quando a entrada for menor que 60 segundos a saída deve ser em segundos `s`.
* Quando a entrada for menor que 60 minutos a saída deve ser em minutos `min`.
* Quando a entrada for maior ou igual a 60 minutos a saída deve ser em horas `h`.

#### *Bytes*

A entrada provida deve ser no formato `<numero> <bytes>`, ex: `16 B`. A unidade usada para entrada será bytes (8 bits).
* Quando a entrada for menor que 1000 bytes a saída deve continuar em `B`.
* Quando a entrada for menor que 1000 elevado a 2 a saída deve ser em kilobyte `kB`.
* Quando a entrada for maior ou igual a 1000 elevado a 2 a saída deve ser em megabyte `MB`.

#### Porcentagem

A entrada provida deve ser no formato `<numero flutuante>`, ex: `0.16`. A unidade usada para entrada será um número flutuante onde, por exemplo, `0.1` significa `10%` e `0.98` significa `98%`.

### Exemplos de entradas e saídas esperadas pelo seu programa

| Entrada | Saída |
| ------ | ------ |
| 30 ms | `30 ms` |
| 5000 ms | `5 s` |
| 5400000 ms | `1.5 h` |
| 0.14 | `14%` |
| 50 B | `50 B` |
| 2000 B | `2 kB` |


## 5 - Paredão BBB

### Considerações Gerais

### O desafio

Você deve implementar o sistema de votação do Paredão do BBB. No Paredão, dois
ou mais participantes se enfrentam em uma votação popular para permanecer na
casa. Ao final da votação, o participante que recebeu mais votos é eliminado do
programa.

Esse sistema pode ser dividido em três partes:

* API para registro e consulta dos votos.
* Client que permite aos usuários participarem da votação, escolhendo um
  participante para ser eliminado.
* Client que permite à produção do programa consultar algumas informações sobre
  o estado da votação.

As regras de negócio para o desenvolvimento da solução são:

* Cada usuário pode votar quantas vezes quiser com uma taxa máxima de 10 votos
  por minuto.
* A votação é chamada na TV em horário nobre, com isso, é esperado um enorme
  volume de votos concentrados em um curto espaço de tempo. A sua solução deve
  suportar pelo menos 1000 votos por segundo.
* A produção do programa gostaria de poder consultar: o total geral de votos, o
  total por participante e o total de votos por hora de cada Paredão.
* Devido ao grande volume de acessos, podemos enfrentar períodos de lentidão ou
  instabilidade nos serviços. Todo o sistema deve estar preparado para isso.
* Como a produção acessará constantemente o estado da votação durante o programa
  ao vivo, as consultas não podem ser lentas. No entanto, é aceitável que a
  resposta apresente dados defasados em momentos de instabilidade.

#### Observações

* Não é necessário fazer nenhuma interface visual.
* Não se preocupe com o registro de candidatos e Paredões. Você pode popular a
  base da forma que achar mais conveniente.
* Gostamos sempre de ser surpreendidos... :wink:

#### Importante!

Em caso de dúvidas, por favor, pergunte.
