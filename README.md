# SGA: Sistema de Gestão de Alocação

API REST escrita em **Go** com o framework **Gin** para gerenciar salas, alunos e turmas de uma instituição de ensino. Ela cuida da matrícula de alunos em turmas e da alocação de turmas em salas, validando capacidade e conflitos de horário.

---

## Contexto acadêmico

| | |
|---|---|
| **Disciplina** | Introdução a Go |
| **Tipo de trabalho** | Atividade prática |
| **Objetivo** | Aplicar os fundamentos da linguagem Go na construção de uma API REST |
| **Ponto de partida** | Projeto de exemplo disponibilizado em aula: [tiagoravache/exemplo_gin](https://github.com/tiagoravache/exemplo_gin) |

Conceitos de Go praticados no projeto:

- structs, tags (`json` e `binding`) e visibilidade por letra maiúscula/minúscula
- ponteiros e zero value
- maps e slices
- métodos com receiver
- concorrência com goroutines e sync.RWMutex
- tratamento de erros como valores (errors.Is, fmt.Errorf com %w)
- API HTTP com Gin: rotas, grupos, leitura e validação de JSON

---

## Funcionalidades

- **Salas:** cadastro com capacidade e recursos (projetor, computadores…), listagem e grade de uso por sala.
- **Alunos:** cadastro com matrícula, nome e e-mail, listagem e busca por matrícula.
- **Turmas:** cadastro com disciplina e professor, e listagem com a quantidade de alunos e o status de alocação.
- **Matrícula:** inclusão de alunos em turmas e listagem dos alunos de uma turma.
- **Alocação:** associação de uma turma a uma sala em um dia da semana e horário.
- **Health check:** status do serviço, data/hora do servidor e versão.

---

## Tecnologias

- [Go](https://go.dev/) 1.22 ou superior
- [Gin](https://gin-gonic.com/) 1.9.1
- Armazenamento **em memória**, sem banco de dados

---

## Estrutura do projeto

```
ams-gin/
├── main.go                  ponto de entrada: monta as dependências e registra as rotas
└── internal/
    ├── domain/              tipos do negócio e erros
    │   ├── sala.go
    │   ├── aluno.go
    │   ├── turma.go
    │   └── erros.go
    ├── service/             regras de negócio e dados em memória
    │   ├── servico.go
    │   ├── sala.go
    │   ├── aluno.go
    │   └── turma.go
    └── handler/             camada HTTP: recebe a requisição, chama o service e responde
        ├── sala_handler.go
        ├── aluno_handler.go
        ├── turma_handler.go
        └── erros.go
```

O caminho de uma requisição é sempre o mesmo:

```
HTTP → handler → service → resposta
```

---

## Como executar

**Pré-requisito:** Go 1.22 ou superior ([download](https://go.dev/dl/)).

```bash
git clone https://github.com/GuiMarobo/ams-gin.git
cd ams-gin
go run .
```

A API sobe em `http://localhost:8080`.

---

## Endpoints

Todas as rotas começam com `/api/v1`.

| Método | Rota | Descrição | Sucesso |
|---|---|---|---|
| GET | `/health` | Status, data/hora e versão do serviço | 200 |
| POST | `/salas` | Cadastra uma sala | 201 |
| GET | `/salas` | Lista as salas | 200 |
| GET | `/salas/:id/agenda` | Grade de uso da sala (filtro opcional `?dia=segunda`) | 200 |
| POST | `/alunos` | Cadastra um aluno | 201 |
| GET | `/alunos` | Lista os alunos | 200 |
| GET | `/alunos/:id` | Busca um aluno pela matrícula | 200 |
| POST | `/turmas` | Cadastra uma turma | 201 |
| GET | `/turmas` | Lista as turmas com quantidade de alunos e alocação | 200 |
| POST | `/turmas/:id/alunos` | Matricula um aluno na turma | 201 |
| GET | `/turmas/:id/alunos` | Lista os alunos da turma | 200 |
| POST | `/turmas/:id/alocar` | Aloca a turma em uma sala | 200 |

---

## Respostas de erro

| Status | Quando acontece |
|---|---|
| **400** Bad Request | JSON inválido, campo obrigatório ausente, capacidade ≤ 0, e-mail inválido, dia ou horário inválido, início depois do fim |
| **404** Not Found | Sala, aluno ou turma não encontrados |
| **409** Conflict | Cadastro duplicado, aluno já matriculado, choque de horário na sala ou na agenda do aluno |
| **422** Unprocessable Entity | Capacidade da sala insuficiente para os alunos da turma |

---

## Limitações

Por ser um projeto de estudo, algumas simplificações foram feitas de propósito:

- Os dados ficam **em memória** e são perdidos quando o servidor reinicia.
- Não há autenticação.
- Cada turma tem **uma única alocação** (um dia e horário). Alocar de novo substitui a anterior.
- As mensagens de validação de formato vêm do validador do Gin e aparecem em inglês.

---

## Autor

**Guilherme Marobo**: [@GuiMarobo](https://github.com/GuiMarobo)

Projeto de uso exclusivamente educacional, desenvolvido na disciplina Introdução a Go.
