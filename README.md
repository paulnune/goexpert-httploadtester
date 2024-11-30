# go-expert-labs-stress-test

Ferramenta para testes de stress HTTP.

## Build local

```bash
make build
```

> Este comand irá gerar o executável `stress-test` na raiz do repositório.

## Executando localmente via Docker

```bash
docker run --network host betonetotbo/stress-test:latest -u http://localhost:8080 -H API_KEY:abc123 -r 1200 -c 50
```
