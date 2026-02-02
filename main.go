package main

import (
	tea "charm.land/bubbletea/v2"
)

func main() {
	// ctx := context.Background()

	_, err := tea.NewProgram(NewBubbleTeamModel()).Run()
	if err != nil {
		panic(err)
	}
}
