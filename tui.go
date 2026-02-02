package main

import tea "charm.land/bubbletea/v2"

type BubbleTeaModel struct{}

func (m BubbleTeaModel) Init() tea.Cmd {
	return nil
}

func (m BubbleTeaModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		// Quit Program
		if msg.Key().String() == "ctrl+c" || msg.Key().String() == "ctrl+q" {
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m BubbleTeaModel) View() tea.View {
	v := tea.View{}

	v.SetContent("hello world")

	return v
}
