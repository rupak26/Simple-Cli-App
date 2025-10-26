package main

func main() {
	todos := Todos{}
	storage := NewStroage[Todos]("todos.json")
	storage.Load(&todos)
	
    cmfFlags := NewCmdFlags()
	cmfFlags.Excute(&todos)
	storage.Save(todos)
}