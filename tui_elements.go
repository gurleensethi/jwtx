package main

import (
	"charm.land/bubbles/v2/textarea"
	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	zone "github.com/lrstanley/bubblezone/v2"
)

// PanelModel is a reusable TUI panel component that can display either
// an editable textarea or a read-only viewport, with optional error status
type PanelModel struct {
	TextArea    textarea.Model
	Viewport    viewport.Model
	Title       string
	Placeholder string
	Focused     bool
	ElementID   Element
	EditingMode bool
	Height      int
	Width       int
	Error       string
	Content     string
}

// NewPanelModel creates a new panel with the specified configuration
func NewPanelModel(ID Element, title, placeholder string, editingMode bool) PanelModel {
	textArea := textarea.New()
	textArea.Placeholder = placeholder
	textArea.Prompt = ""
	textArea.ShowLineNumbers = false

	viewportModel := viewport.New()
	viewportModel.SoftWrap = true

	return PanelModel{
		ElementID:   ID,
		TextArea:    textArea,
		Viewport:    viewportModel,
		Title:       title,
		Placeholder: placeholder,
		Focused:     false,
		EditingMode: editingMode,
		Height:      0,
		Width:       0,
		Error:       "",
		Content:     "",
	}
}

func (m PanelModel) Init() tea.Cmd {
	return nil
}

func (m PanelModel) Update(msg tea.Msg) (PanelModel, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case FocusElementMsg:
		if msg.Element == m.ElementID {
			m.Focused = true
			if m.EditingMode {
				m.TextArea.Focus()
			}
		} else {
			m.Focused = false
			if m.EditingMode {
				m.TextArea.Blur()
			}
		}
		return m, nil
	}

	if m.EditingMode {
		m.TextArea, cmd = m.TextArea.Update(msg)
	} else {
		m.Viewport, cmd = m.Viewport.Update(msg)
	}

	return m, cmd
}

func (m PanelModel) View() string {
	title := styleTitle
	box := styleBox
	if m.Focused {
		title = styleTitleSelected
		box = styleBoxActive
	}

	var content string
	if m.EditingMode {
		content = m.TextArea.View()
	} else {
		content = m.Viewport.View()
	}

	width := lipgloss.Width(content)

	statusBar := styleStatus.Width(width).Render("")
	if m.Error != "" {
		statusBar = styleStatusError.Width(width).Render(m.Error)
	}

	return zone.Mark(
		string(m.ElementID),
		box.Render(
			lipgloss.JoinVertical(lipgloss.Left,
				title.Render(m.Title),
				content,
				statusBar,
			),
		),
	)
}

func (m *PanelModel) SetHeight(height int) {
	m.Height = height
	// Account for title, borders, and status bar (typically 3-4 lines total)
	internalHeight := height - 4
	if internalHeight < 1 {
		internalHeight = 1
	}
	if m.EditingMode {
		m.TextArea.SetHeight(internalHeight)
	} else {
		m.Viewport.SetHeight(internalHeight)
	}
}

func (m *PanelModel) SetWidth(width int) {
	m.Width = width - 2
	// Account for borders (typically 2 columns total)
	if m.EditingMode {
		m.TextArea.SetWidth(width - 2)
	} else {
		m.Viewport.SetWidth(width - 2)
	}
}

func (m *PanelModel) SetValue(content string) {
	m.Content = content
	if m.EditingMode {
		m.TextArea.SetValue(content)
	} else {
		m.Viewport.SetContent(content)
	}
}

func (m PanelModel) GetValue() string {
	if m.EditingMode {
		return m.TextArea.Value()
	}
	return m.Content
}

func (m *PanelModel) SetError(error string) {
	m.Error = error
}

func (m *PanelModel) Blur() {
	m.Focused = false
	if m.EditingMode {
		m.TextArea.Blur()
	}
}

func (m *PanelModel) SetEditingMode(editing bool) {
	m.EditingMode = editing
}

func (m PanelModel) IsEditing() bool {
	return m.EditingMode
}
