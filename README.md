# Challenge: Portão do Enem

Esse projeto é uma possível solução escrita em `Golang` para o desafio `Portão do Enem` do [CodeAspiras](https://codeaspiras.dev/discord). Sinta-se a vontade para resolver o mesmo desafio em outras linguagens!

## Desafio

O governo decidiu instalar catracas automáticas nas escolas onde vão ser ministradas as provas do Enem. Agora não é mais necessário que uma pessoa fique abrindo e/ou fechando os portões, apenas monitorando para ver se alguém não está pulando a catraca.

Para que o programa funcione, a catraca precisa acionar uma função para verificar se a entrada pode ser liberada ou não. Você irá desenvolver essa função.

As entradas são liberadas entre 10h da manhã e meio-dia. Em casos de horário de verão, a entrada fica entre 9h e 11h.

Faça uma função que recebe uma string informando a hora (em formato hh:mm) e um boolean informando se é ou não horário de verão. Essa função deve retornar um outro  boolean informando se a entrada está liberada (true) ou não (false), respeitando as regras da descrição do desafio

## Nota

O programa deve validar se a hora está preenchida corretamente.
Por exemplo, 11:60 não existe, mas 11:59 sim.

## Como rodar a solução

1. Você precisar ter o [Golang v1.20+](https://go.dev/) instalado na sua máquina. Caso não o tenha, acesse o link informado anteriormente, baixe e instale!
2. Abra um terminal no diretório raiz desse projeto;
3. Rode `go run .` e siga as instruções do programa!

Você também pode compilar o projeto em um executável rodando `go build .`. Um arquivo executável será criado na raiz do projeto para que possa fazer o que quiser.