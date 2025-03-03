package cli

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	"github.com/Ma-Leal/to-do-list/internal/usecase"
)

type TaskHandler struct {
	Usecase *usecase.TaskUseCase
}

func NewTaskHandler(usecase usecase.TaskUseCase) *TaskHandler {
	return &TaskHandler{Usecase: &usecase}
}

func (t *TaskHandler) Run() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		t.showMenu()
		scanner.Scan()
		args := strings.Fields(scanner.Text())

		if len(args) == 0 {
			continue
		}

		command := strings.ToLower(args[0])
		switch command {
		case "add":
			t.handleAdd(args)
			t.handleList()
		case "update":
			t.handleUpdate(args)
			t.handleList()
		case "status":
			t.handleStatus(args)
			t.handleList()
		case "delete":
			t.handleDelete(args)
			t.handleList()
		case "listall":
			t.handleList()
		case "list":
			t.handleListByStatus(args)
		case "exit":
			fmt.Println("👋 Saindo do Task Tracker. Até logo!")
			return
		default:
			fmt.Println("❌ Comando inválido! Digite um comando válido.")
		}
	}

}

func (t *TaskHandler) showMenu() {
	fmt.Println("\n=================================")
	fmt.Println("🎯 Task Tracker CLI")
	fmt.Println("=================================")
	fmt.Println("Comandos disponíveis:")
	fmt.Println("  📌  add <descrição>           → Adicionar uma nova tarefa")
	fmt.Println("  ✏️   update <id> <descrição>   → Atualizar a descrição de uma tarefa")
	fmt.Println("  🔄  status <id> <status>      → Alterar o status da tarefa")
	fmt.Println("  ❌  delete <id>               → Remover uma tarefa")
	fmt.Println("  📋  ListAll               	→ Listar todas as tarefas")
	fmt.Println("  📋  list [status]		→ Listar as tarefas por status")
	fmt.Println("                                   1 → 📝 To-Do")
	fmt.Println("                                   2 → 🚧 In-Progress")
	fmt.Println("                                   3 → ✅ Done")
	fmt.Println("  🚪  exit                      → Sair do programa")
	fmt.Println("=================================")
	fmt.Print("🔹 Digite um comando: ")
}

func (t *TaskHandler) handleAdd(args []string) {
	if len(args) < 2 {
		fmt.Println("Erro: Digite uma descrição para a tarefa.")
		return
	}
	description := strings.Join(args[1:], " ")
	task, err := t.Usecase.CreateTask(description, 1)
	if err != nil {
		fmt.Println("Erro ao adicionar tarefa:", err)
		return
	}
	fmt.Printf("\n✅ Tarefa #%d adicionada: \"%s\" [%s]\n",
		task.ID, task.Description, task.Status)
}

func (t *TaskHandler) handleDelete(args []string) {
	if len(args) < 2 {
		fmt.Println("Erro: Uso correto -> delete <id>")
		return
	}
	id := parseID(args[1])
	if id == -1 {
		return
	}
	err := t.Usecase.DeleteTask(id)
	if err != nil {
		fmt.Println("Erro ao deletar tarefa:", err)
	} else {
		fmt.Println("✅ Tarefa deletada!")
	}
}

func (t *TaskHandler) handleList() {
	clearScreen()
	tasks, err := t.Usecase.ListTasks()
	if err != nil {
		fmt.Println("Erro ao listar tarefas:", err)
		return
	}
	fmt.Println("\n📋 Lista de tarefas:")
	for _, task := range tasks {
		fmt.Printf("🆔 %d | %s | Status: %s\n", task.ID, task.Description, task.Status)
	}
}

func (t *TaskHandler) handleUpdate(args []string) {
	if len(args) < 3 {
		fmt.Println("Erro: Uso correto -> update <id> <nova descrição>")
		return
	}
	id := parseID(args[1])
	if id == -1 {
		return
	}
	description := strings.Join(args[2:], " ")
	_, err := t.Usecase.UpdateTaskDescription(id, description)
	if err != nil {
		fmt.Println("Erro ao atualizar tarefa:", err)
	} else {
		fmt.Println("✅ Tarefa atualizada!")
	}
}

func (t *TaskHandler) handleStatus(args []string) {
	fmt.Println(args)
	if len(args) < 3 {
		fmt.Println("Erro: Uso correto -> status <id> <novo status>")
		return
	}
	id := parseID(args[1])
	if id == -1 {
		return
	}
	status := parseID(args[2])
	if id == -1 {
		return
	}
	_, err := t.Usecase.UpdateTaskStatus(id, status)
	if err != nil {
		fmt.Println("Erro ao atualizar status:", err)
	} else {
		fmt.Println("✅ Status atualizado!")
	}
}

func (t *TaskHandler) handleListByStatus(args []string) {
	clearScreen()
	if len(args) < 2 {
		fmt.Println("Erro: Uso correto -> list <status_id>")
		return
	}
	status_id := parseID(args[1])
	if status_id == -1 {
		return
	}
	tasks, err := t.Usecase.ListTasksByStatus(status_id)
	if err != nil {
		fmt.Println("Erro ao listar tarefas:", err)
		return
	}
	fmt.Println("\n📋 Lista de tarefas:")
	for _, task := range tasks {
		fmt.Printf("🆔 %d | %s | Status: %s\n", task.ID, task.Description, task.Status)
	}
}

func clearScreen() {
	cmd := exec.Command("clear")
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func parseID(input string) int {
	id, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Erro: ID inválido. Deve ser um número.")
		return -1
	}
	return id
}
