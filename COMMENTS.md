## 1 - Intervalo de Números

### Histórico
  - Primeira versão para validar o raciocinio
  - Inclusão de testes
  - Refatoração
  - Incluir Makefile
  - Aumentar cobertura de código/Refatoração

### Validações
  - Validar se a lista é válida:
    - Vazia
    - Somente números inteiros
    Retorna Entrada Inválida
  

### Melhorias
  - Aceitar uma lista desordenada, ordenar e agrupar os intervalos / Done!


## 2 - Troco

### Histórico

  - Fiz um código usando estrutura de repetição depois alterei para recursão
  - Para conversão da moeda no poderiamos usar o [currency](https://godoc.org/golang.org/x/text/currency#example-Query), mas optei em fazer o replace ",", "." para evitar usar outros pacotes.
  - Para deixar o código testável eu precisei refator, fazer teste pegando o print no console não acho legal. 
  - Acabei criando mais tipagens para poder fazer map de troco mais legível e testavel
  - Criei um método só para printar o recibo
  - Precisei de um work around, ao criar o teste onde tinha que dividir o 14.20 /10.0 retorna 4.1999999999999, então precisei criar o método formatFloat 
  - Ajuste passagem parametro os.stdIn
  - Incluir Makefile
  - Aumentar cobertura de código/Refatoração

### Validações
 - Inclui validação quando os valores digitados estão incorretos
 - A momentos em que o troco pode ser inferir a 0.05 que a menor fração, neste caso assumi que o retorno é zero. Não temos bala! =D

### Melhorias

- Conversão da moeda para o BRL usando [currency](https://godoc.org/golang.org/x/text/currency#example-Query).
- Criar uma lista de moeda conforme a type Unit struct. No caso BRL.
  
    Additional common currencies as defined by CLDR.
    BRL Unit = Unit{brl}
- Fazer a quebra de caixa com o arredamento quando houver troco inferior ao menor valor da lista de moedas.


## 3 - Breath Fantasy

### Histórico
  - Primeira versão para validar o raciocínio
  - Inclusão de testes
  - Refatoração com recursão
  - Fix pacote rand
  - Incluir Makefile
  - Aumentar cobertura de código/Refatoração

### Validações
  - Validar se os dados dos personagens foram preenchidos corretamente:
    - Nome não vazio
    - Poder não vazio e somente número inteiro
    - Energia não vazio e somente número inteiro
  Retorna Entrada Inválida
  

### Melhorias
  - Fazer W.0 ou limitar o número de **rounds**, colocar timer ou correlacionados.
  - Ataque simultaneo


## 5 - Paredão BBB

Para mim é o desafio mais interessante, e seria o preferido para desenvolver se fosse apenas 1 deles dentro do tempo estipulado.
Um sistema desta proporção que recebe uma carga imensa de requisições em poucos minutos e com picos inesperados causados talvez por 
um comentário nas redes sociais, propaganda de TV e até mesmo durante a transmissão da final com certeza terá muitos desafios de escalabilidade,
tolerância a falha e também de segurança, pois a cada dia que passa os bots estão mais sofisticados.

Eu dividiria em mais partes das que foram propostas. Por quê?

Acho que dividiria em 3.
- Serviço de autenticação
  - Banco de dado exclusivo
- Registro do voto   - recebe a requisição de voto / manda para o serviço de mensageria
  - Precisaria de autenticação
  - Banco só com dados dos participantes do paredão
  
- Consulta de votos  - recebe a requisição de voto
  - Banco com os dados dos votos computados
  - Precisaria de autenticação, exclusiva dos participantes do programa

- Computação do voto - processa as mensagens, computando os votos
  - Rabbit MQ ou Kafka

Partiria da premissa de fazer o on-premisse usando o Kubernets, tornando agnóstico a algum provedor de Cloud que já nos dá muitos
recursos que precisamos, mas que no fim acaba acoplando sua a solução a elas. E temos o fator pico, que pode gerar um gasto
não previsto.

No quesito segurança eu nunca tive uma experiência diferente do OAuth. 
Neste caso por causa dos bots eu não saberia como lidar com isso neste momento, pois muitos deles já conseguem quebrar os captchas.
Acho que um proxy na frente poderia fazer o papel de filter e criar algumas limitações por token, ip do requisitante entre outros.code

Por que eu usaria um barramento entre os dois, aceitaria uma requisição de voto, responderia imediatamente para o cliente que o voto será computado e encaminharia este voto para um fila, onde eu poderia conectador vários consumers para computador este voto.

Tenho mais conhecimento do SpringBoot, mas quando começamos a colocar no Kubernets, percebemos que ele consome bastante recurso como memória e demora a subir os pods quando a necessidade de replica. No caso tentaria utilizar outros microframeworks, como Quarkus ou Micronault por exemplo.

Veja o diagrama abaixo, é um protótipo de como eu penso nesta solução.
Digamos que é um primeiro desenho e que claro com as discussões técnicas de um grupo, poderia sofrer muitas mudanças.


### Diagrama
![Diagrama Votação](/diagrama_do_serviço_de_votação.png)

Escolhi o Golang para fazer o testes porque tenho gostado da linguagem e queria ter desafios para fixar ainda mais.

Caso queira ver meu código em JAVA vou deixar 2 repositórios que eu codei recentemente em outros desafios.

[Java Swagger RestApi](https://github.com/SamuelsSantos/controle-patrimonial)

[Quarkus Codenation](https://github.com/SamuelsSantos/codenation-desafio)
