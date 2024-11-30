
# HTTP Load Tester CLI 🚀

Sistema CLI em Go para realizar testes de carga em serviços web. O usuário pode especificar a URL do serviço, o número total de requisições e o nível de concorrência. Inclui relatórios detalhados após os testes. Desenvolvido por Paulo Nunes.

## Objetivo 🎯

O objetivo deste projeto é criar um sistema CLI em Go para realizar testes de carga em serviços web. O usuário fornece a URL do serviço, o número total de requisições e o número de chamadas simultâneas. Após os testes, o sistema gera um relatório detalhado com informações específicas da execução.

## Funcionalidades 📋

- **Configuração via linha de comando (parâmetros obrigatórios):**
  - URL do serviço (`--url`)
  - Total de requisições (`--requests`)
  - Concorrência (`--concurrency`)
- **Execução do teste:**
  - Realizar requisições HTTP para a URL especificada.
  - Distribuir os requests de acordo com o nível de concorrência definido.
  - Garantir que o número total de requests seja cumprido.
- **Relatórios detalhados:**
  - Tempo total de execução.
  - Total de requisições realizadas.
  - Quantidade de respostas com status HTTP 200.
  - Distribuição de outros códigos HTTP (404, 500, etc.).
- Flexível para execução via Docker, Podman, Docker Compose ou Podman Compose.

## Requisitos 📦

- Golang >= 1.23.1
- Docker ou Podman
- Docker Compose ou Podman Compose

## Como Executar 🚀

### Usando Docker Compose

1. Configure o arquivo `.env` com a URL desejada para o teste:

   ```
   TARGET_URL=https://example.com
   ```

2. Suba o serviço com:

   ```
   docker compose up --build
   ```

3. Para rodar o CLI com parâmetros personalizados, use:

   ```
   docker compose run httploadtester --url=https://example.com --requests=100 --concurrency=10
   ```

### Usando Makefile

Execute os comandos abaixo para maior automação:

- **Compilação Local**

  ```
  make build
  ```

- **Testes**

  ```
  make test
  ```

- **Executar via CLI**

  ```
  make run -- --url=https://example.com --requests=100 --concurrency=10
  ```

- **Limpeza**

  ```
  make clean
  ```

### Executando com `docker run`

A aplicação também pode ser executada diretamente com o comando `docker run`:

```
docker image ls
docker run <imagem_docker> --url=http://google.com --requests=1000 --concurrency=10
```

Certifique-se de substituir `<imagem_docker>` pelo nome ou ID da imagem Docker construída.

## Exemplos de Uso 🛠️

- **Executar Teste com Parâmetros Padrão**

  ```
  docker compose run httploadtester --url=https://example.com --requests=100 --concurrency=10
  ```

- **Executar com Personalização**

  ```
  docker compose run httploadtester --url=https://example.com --requests=500 --concurrency=20 --method=POST --timeout=10s
  ```

## Relatórios 📊

Após cada teste, será gerado um relatório no seguinte formato:

```
+----------------+--------------+------------+--------+
| TOTAL REQUESTS |   TIME SPENT | SUCCESSFUL | HTTP 0 |
+----------------+--------------+------------+--------+
|            500 |    5.002s    |        450 |     50 |
+----------------+--------------+------------+--------+
```

O relatório é exibido diretamente no console após a execução dos testes.

## Estrutura do Projeto 📂

```
.
├── build
│   └── Dockerfile
├── cmd
│   └── root.go
├── go.mod
├── internal
│   └── stresstest
│       ├── reports.go
│       ├── request.go
│       ├── stress.go
│       ├── utils.go
├── main.go
├── Makefile
├── README.md
└── .env
```

## Validação

Todos os requisitos do desafio foram atendidos:

- Configuração via CLI ✅
- Teste de carga com concorrência ✅
- Geração de relatórios detalhados ✅
- Execução via Docker/Podman ✅
- Flexibilidade para personalização ✅

## Autor

**Paulo Henrique Nunes Vanderley**

- 🌐 [Site Pessoal](https://www.paulonunes.dev/)
- 🌐 [GitHub](https://github.com/paulnune)
- ✉️ Email: [paulo.nunes@live.de](mailto:paulo.nunes@live.de)
- 🚀 Aluno da Pós GoExpert 2024 pela FullCycle

## 🎉 Agradecimentos

Este projeto foi desenvolvido como parte de um desafio técnico, combinando aprendizado e prática com Go. 🚀