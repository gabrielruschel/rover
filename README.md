# rover

Este repositório implementa a solução para o problema de navegação de *rovers* em um *plateau*. A solução foi desenvolvida em Go (golang) utilizando apenas funcionalidades disponíveis na biblioteca nativa (possui a dependência da biblioteca `github.com/stretchr/testify v1.10.0` apenas para realizar *asserts* nos testes unitários).

## Funcionamento Interno

Para os dados de entrada, o programa realiza a leitura de um arquivo de texto ([exemplo de arquivo de entrada](input.example.txt)) e disponibiliza a saída através da saída padrão e também em um arquivo de texto.

### Configurações

É possível realizar algumas configurações via variáveis de ambiente:

- `INPUT_FILE`: Define o nome do arquivo de texto entrada dos dados
  -  Padrão: `input.txt`
- `OUTPUT_FILE`: Define o nome do arquivo de texto de saída da posição final dos *rovers*
  -  Padrão: `output.txt`
 - `LOG_LEVEL`: Nível dos logs a serem mostrados
   - Padrão: `ERROR` (apenas logs de erro são mostrados)
   - Valores aceitos: `DEBUG`, `INFO`, `WARN` e `ERROR` 

### Leitura de Input e Navegação

Seguindo o enunciado do problema, assume-se que o *plateau* de navegação (e todas as suas coordenadas) são valores positivos, portanto, o parse destas informações é feito para números inteiros sem sinal de 64bits (`uint64`). Por conta disso o programa não permite definir coordenadas negativas (tanto para o `plateau` quanto para a posição inicial dos *rovers*).

A interpretação das coordenadas e instruções para cada *rover* é feita de forma sequencial, seguindo a mesma estrutura do enunciado do problema.

- Caso um *rover* possua coordenadas inciais inválidas, é apresentado o erro no log e as instruções do *rover* são ignoradas, seguindo para a próxima navegação
  - Coordenadas negativas são consideradas inválidas
  - Coordenadas fora dos limites do *plateau* são consideradas inválidas
  - Coordenadas que são as mesmas da posição final de um *rover* anterior são consideradas inválidas
- Caso um *rover* receba uma instrução não conhecida, ela é ignorada
- Caso a movimentação de um *rover* seja para uma coordenada inválida, um log de aviso é exibido e a coordenada atual é preservada, prosseguindo para a próxima instrução

Finalmente, como mencionado anteriormente, ao final da navegação de cada *rover* é exibido na saída padrão as coordenadas finais e orientação (`2 3 W`, por exemplo). Estas informações também são escritas no arquivo de saída ao final da execução.

## Execução

Para buildar/executar o programa, é disponibilizado um `Makefile`

Para realizar o build (diretório `/bin`)
```
make build
```

Para buildar e executar (utilizando configurações padrão)
```
make run
```

## Testes

Para executar os testes
```
make test
```
