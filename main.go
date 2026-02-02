package main

import (
	tea "charm.land/bubbletea/v2"
)

func main() {
	// ctx := context.Background()
	model := BubbleTeaModel{}

	_, err := tea.NewProgram(model).Run()
	if err != nil {
		panic(err)
	}
}
