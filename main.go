package main

import (
	tea "charm.land/bubbletea/v2"
	zone "github.com/lrstanley/bubblezone/v2"
)

func main() {
	// ctx := context.Background()

	zone.NewGlobal()

	_, err := tea.NewProgram(NewBubbleTeamModel()).Run()
	if err != nil {
		panic(err)
	}
}
