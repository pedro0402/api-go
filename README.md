
# 🧠 API CRUD em Go

Essa é uma API básica de uma To Do List desenvolvida em Go. O projeto permite operações CRUD (Create, Read, Update, Delete) em tarefas, utilizando um banco de dados PostgreSQL.

---

## 📑 Índice

- [Tecnologias Utilizadas](#-tecnologias-utilizadas)
- [Instalação](#-instalação)
- [Uso](#-uso)
- [Endpoints](#-endpoints)
- [Configuração / Variáveis de Ambiente](#-configuração--variáveis-de-ambiente)
- [Contribuição](#-contribuição)
- [Licença](#-licença)
- [Autor](#-autor)

---

## 🚀 Tecnologias Utilizadas

- Go 1.22.1
- PostgreSQL

---

## ⚙️ Instalação

1. Clone o repositório:
```bash
git clone https://github.com/pedro0402/api-go.git
cd api-go
```

2. Instale as dependências:
```bash
go mod download
```

3. Configure o banco de dados PostgreSQL e ajuste as variáveis de ambiente conforme necessário.

---

## ▶️ Uso

Para rodar a aplicação localmente:

```bash
go run main.go
```

A API será iniciada em:  
`http://localhost:8080`

---

## 🔗 Endpoints

- `GET    /tasks` – Retorna todas as tarefas  
- `POST   /tasks` – Cria uma nova tarefa  
- `PUT    /tasks/1` – Atualiza a tarefa com ID 1  
- `DELETE /tasks/1` – Deleta a tarefa com ID 1  

---

## ⚙️ Configuração / Variáveis de Ambiente

Crie um arquivo `.env` com as seguintes variáveis:

```
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=todo
```

Certifique-se de que o banco esteja rodando com as credenciais corretas.

---

## 🤝 Contribuição

Contribuições são bem-vindas!  
Para colaborar:

1. Faça um fork do projeto
2. Crie uma branch com sua feature: `git checkout -b minha-feature`
3. Commit suas mudanças: `git commit -m 'feat: minha nova feature'`
4. Envie para o repositório remoto: `git push origin minha-feature`
5. Abra um Pull Request

---

## 📄 Licença

*Este projeto ainda não possui uma licença definida.*  
> Adicione uma licença (ex: MIT, Apache 2.0) e atualize esta seção.

---

## 👨‍💻 Autor

Pedro Henrique  
- GitHub: [github.com/pedro0402](https://github.com/pedro0402)  
- LinkedIn: https://www.linkedin.com/in/pedro-moraes-142884252/   

---
