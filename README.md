
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

## Relatórios 📊

Após cada teste, o sistema gera um relatório com os seguintes dados:

- **Tempo total gasto na execução**: Quanto tempo levou para concluir todas as requisições.
- **Quantidade total de requisições realizadas**: Número total de requisições enviadas.
- **Quantidade de requisições com status HTTP 200**: Requisições bem-sucedidas.
- **Distribuição de outros códigos de status HTTP**: Contagem de códigos como 404, 500, etc.

---

### Como Executar 🚀

### Usando Docker Compose

1. Configure o arquivo `.env` com as variáveis desejadas para o teste:

   ```
   TARGET_URL=https://paulonunes.dev
   REQUESTS=100
   CONCURRENCY=10
   ```

   - **`TARGET_URL`**: A URL que será testada.
   - **`REQUESTS`**: O número total de requisições que serão realizadas durante o teste.
   - **`CONCURRENCY`**: O número de requisições simultâneas que serão enviadas durante o teste.

2. Suba o serviço com:

   ```
   docker compose up --build
   ```

3. Para rodar o CLI com parâmetros personalizados, use:

   ```
   docker compose run httploadtester --url=https://paulonunes.dev --requests=70 --concurrency=16
   ```

### Executando com `docker run`

A aplicação também pode ser executada diretamente com o comando `docker run`:

```
docker run httploadtester --url=http://google.com --requests=200 --concurrency=4
```

Certifique-se de utilizar a imagem correta e que ela tenha sido buildada. 

---

### Evidências de Execução

#### **1. Teste Básico com 100 Requisições para um Site Existente**
![Teste 1](.assets/1.png)
Comando executado:
```bash
sudo docker compose run httploadtester --url=https://example.com --requests=100 --concurrency=1
```

#### **2. Teste com Concorrência de 10 Requisições Simultâneas**
![Teste 2](.assets/2.png)
Comando executado:
```bash
sudo docker compose run httploadtester --url=https://example.com --requests=100 --concurrency=10
```

#### **3. Teste com 70 Requisições em Alta Concorrência**
![Teste 3](.assets/3.png)
Comando executado:
```bash
sudo docker compose run httploadtester --url=https://paulonunes.dev --requests=70 --concurrency=16
```

#### **4. Teste com 200 Requisições e Baixa Concorrência**
![Teste 4](.assets/4.png)
Comando executado:
```bash
sudo docker run httploadtester --url=http://google.com --requests=200 --concurrency=4
```

#### **5. Teste com HTTP 405 (Método Não Permitido)**
![Teste 5](.assets/5.png)
Comando executado:
```bash
sudo docker compose run httploadtester --url=https://example.com --requests=500 --concurrency=20 --method=POST --timeout=10s
```

#### **6. Teste com HTTP 404 (Não Encontrado)**
![Teste 6](.assets/6.png)
Comando executado:
```bash
sudo docker compose run httploadtester --url=http://httpstat.us/404 --requests=80 --concurrency=6
```

#### **7. Teste com HTTP 500 (Erro Interno do Servidor)**
![Teste 7](.assets/7.png)
Comando executado:
```bash
sudo docker compose run httploadtester --url=http://httpstat.us/500 --requests=100 --concurrency=20
```

---

## Autor

**Paulo Henrique Nunes Vanderley**

- 🌐 [Site Pessoal](https://www.paulonunes.dev/)
- 🌐 [GitHub](https://github.com/paulnune)
- ✉️ Email: [paulo.nunes@live.de](mailto:paulo.nunes@live.de)
- 🚀 Aluno da Pós GoExpert 2024 pela FullCycle

## 🎉 Agradecimentos

Este projeto foi desenvolvido como parte de um desafio técnico, combinando aprendizado e prática com Go. 🚀