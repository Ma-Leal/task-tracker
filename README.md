# Task Tracker

Um gerenciador de tarefas, desenvolvido em Go, seguindo boas práticas de Clean Architecture.

## 📌 Funcionalidades
- Criar, listar, atualizar e excluir tarefas
- Persistência baseada em arquivos JSON
- Organização modular com separação clara de responsabilidades
- Fácil extensibilidade e manutenção

## 🏗 Arquitetura
Este projeto segue a arquitetura hexagonal, separando lógica de negócio da infraestrutura. As principais camadas incluem:

- **Entities**: Modelos de domínio
- **Use Cases**: Regras de negócio
- **Repositories**: Persistência em JSON
- **Handlers**: Interface com CLI

## 📖 Referências
- [Roadmap Task Tracker](https://roadmap.sh/projects/task-tracker)

## 📜 Menu de Comandos
```
=================================
🎯 Task Tracker CLI
=================================
Comandos disponíveis:
  📌  add <descrição>           → Adicionar uma nova tarefa
  ✏️   update <id> <descrição>   → Atualizar a descrição de uma tarefa
  🔄  status <id> <status>      → Alterar o status da tarefa
  ❌  delete <id>               → Remover uma tarefa
  📋  ListAll                   → Listar todas as tarefas
  📋  list [status]             → Listar as tarefas por status
                                   1 → 📝 To-Do
                                   2 → 🚧 In-Progress
                                   3 → ✅ Done
  🚪  exit                      → Sair do programa
=================================
```