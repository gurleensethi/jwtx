package main

import (
	"charm.land/bubbles/v2/textarea"
	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	zone "github.com/lrstanley/bubblezone/v2"
)

type JWTTokenModel struct {
	ZoneID      string
	TextArea    textarea.Model
	Viewport    viewport.Model
	Focused     bool
	ElementID   Element
	EditingMode bool
	Height      int
	Width       int
	Error       string
	Content     string
}

func NewJWTTokenModel() JWTTokenModel {
	textArea := textarea.New()
	textArea.Placeholder = PlaceholderJWT
	textArea.Prompt = ""
	textArea.ShowLineNumbers = false

	viewportModel := viewport.New()
	viewportModel.SoftWrap = true

	return JWTTokenModel{
		TextArea:    textArea,
		Viewport:    viewportModel,
		Focused:     false,
		ElementID:   "",
		EditingMode: true, // Default to editing mode,
		Height:      0,
		Width:       0,
		Error:       "",
		Content:     "",
	}
}

func (m JWTTokenModel) Init() tea.Cmd {
	return nil
}

func (m JWTTokenModel) Update(msg tea.Msg) (JWTTokenModel, tea.Cmd) {
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
	default:
		if m.EditingMode {
			m.TextArea, cmd = m.TextArea.Update(msg)
		} else {
			m.Viewport, cmd = m.Viewport.Update(msg)
		}
	}

	return m, cmd
}

func (m JWTTokenModel) View() string {
	title := styleTitle
	box := styleBox
	if m.Focused {
		title = styleTitleSelected
		box = styleBoxActive
	}

	statusBar := styleStatus.Width(m.Width).Render("")
	if m.Error != "" {
		statusBar = styleStatusError.Width(m.Width).Render(m.Error)
	}

	var content string
	if m.EditingMode {
		content = m.TextArea.View()
	} else {
		content = m.Viewport.View()
	}

	return zone.Mark(
		m.ZoneID,
		box.Render(
			lipgloss.JoinVertical(lipgloss.Left,
				title.Render(TitleJWTToken),
				content,
				statusBar,
			),
		),
	)
}

func (m *JWTTokenModel) SetHeight(height int) {
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

func (m *JWTTokenModel) SetWidth(width int) {
	m.Width = width - 2
	// Account for borders (typically 2 columns total)
	if m.EditingMode {
		m.TextArea.SetWidth(width - 2)
	} else {
		m.Viewport.SetWidth(width - 2)
	}
}

func (m *JWTTokenModel) SetToken(token string) {
	m.Content = token
	if m.EditingMode {
		m.TextArea.SetValue(token)
	} else {
		m.Viewport.SetContent(token)
	}
}

func (m JWTTokenModel) GetToken() string {
	if m.EditingMode {
		return m.TextArea.Value()
	} else {
		return m.Content
	}
}

func (m JWTTokenModel) GetTokenText() string {
	if m.EditingMode {
		return m.TextArea.Value()
	} else {
		return m.Content
	}
}

func (m *JWTTokenModel) SetError(error string) {
	m.Error = error
}

func (m *JWTTokenModel) Blur() {
	if m.EditingMode {
		m.TextArea.Blur()
	} else {
		// For viewport, no blur needed
	}
}

func (m *JWTTokenModel) SetEditingMode(editing bool) {
	m.EditingMode = editing
}

func (m JWTTokenModel) IsEditing() bool {
	return m.EditingMode
}

// ================================================================================

type JWTSecretModel struct {
	ZoneID      string
	TextArea    textarea.Model
	Focused     bool
	ElementID   Element
	EditingMode bool
	Height      int
	Width       int
	Error       string
}

func NewJWTSecretModel() JWTSecretModel {
	textArea := textarea.New()
	textArea.Placeholder = PlaceholderSecret
	textArea.Prompt = ""
	textArea.ShowLineNumbers = false

	return JWTSecretModel{
		TextArea:    textArea,
		Focused:     false,
		ElementID:   "",
		EditingMode: true, // Default to editing mode
		Height:      0,
		Width:       0,
		Error:       "",
	}
}

func (m JWTSecretModel) Init() tea.Cmd {
	return nil
}

func (m JWTSecretModel) Update(msg tea.Msg) (JWTSecretModel, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case FocusElementMsg:
		if msg.Element == m.ElementID {
			m.Focused = true
			m.TextArea.Focus()
		} else {
			m.Focused = false
			m.TextArea.Blur()
		}
		return m, nil
	}

	m.TextArea, cmd = m.TextArea.Update(msg)

	return m, cmd
}

func (m JWTSecretModel) View() string {
	title := styleTitle
	box := styleBox
	if m.Focused {
		title = styleTitleSelected
		box = styleBoxActive
	}

	width := lipgloss.Width(m.TextArea.View())

	statusBar := styleStatus.Width(width).Render("")
	if m.Error != "" {
		statusBar = styleStatusError.Width(width).Render(m.Error)
	}

	return zone.Mark(
		m.ZoneID,
		box.Render(
			lipgloss.JoinVertical(lipgloss.Left,
				title.Render(TitleSecret),
				m.TextArea.View(),
				statusBar,
			),
		),
	)
}

func (m *JWTSecretModel) SetHeight(height int) {
	m.Height = height
	// Account for title, borders, and status bar (typically 3-4 lines total)
	internalHeight := height - 4
	if internalHeight < 1 {
		internalHeight = 1
	}
	m.TextArea.SetHeight(internalHeight)
}

func (m *JWTSecretModel) SetWidth(width int) {
	m.Width = width - 2
	// Account for borders (typically 2 columns total)
	m.TextArea.SetWidth(width - 2)
}

func (m JWTSecretModel) GetSecret() string {
	return m.TextArea.Value()
}

func (m JWTSecretModel) SetSecret(secret string) {
	m.TextArea.SetValue(secret)
}

func (m *JWTSecretModel) Blur() {
	m.TextArea.Blur()
}

func (m *JWTSecretModel) SetEditingMode(editing bool) {
	m.EditingMode = editing
}

func (m JWTSecretModel) IsEditing() bool {
	return m.EditingMode
}

func (m *JWTSecretModel) SetError(error string) {
	m.Error = error
}

// ================================================================================

type JWTHeaderModel struct {
	ZoneID      string
	TextArea    textarea.Model
	Viewport    viewport.Model
	Title       string
	EditingMode bool
	Focused     bool
	ElementID   Element
	Height      int
	Width       int
	Error       string
	Content     string
}

func NewJWTHeaderModel() JWTHeaderModel {
	textArea := textarea.New()
	textArea.Placeholder = "Enter header JSON here..."
	textArea.Prompt = ""
	textArea.ShowLineNumbers = false

	viewportModel := viewport.New()
	viewportModel.SoftWrap = true

	return JWTHeaderModel{
		TextArea:    textArea,
		Viewport:    viewportModel,
		Title:       "",
		EditingMode: true, // Default to editing mode for encoder
		Focused:     false,
		ElementID:   "",
		Height:      0,
		Width:       0,
		Error:       "",
		Content:     "",
	}
}

func (m JWTHeaderModel) Init() tea.Cmd {
	return nil
}

func (m JWTHeaderModel) Update(msg tea.Msg) (JWTHeaderModel, tea.Cmd) {
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

	var cmd tea.Cmd
	if m.Focused && m.EditingMode {
		m.TextArea, cmd = m.TextArea.Update(msg)
	} else if m.Focused && !m.EditingMode {
		m.Viewport, cmd = m.Viewport.Update(msg)
	}
	return m, cmd
}

func (m JWTHeaderModel) View() string {
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
		m.ZoneID,
		box.Render(
			lipgloss.JoinVertical(lipgloss.Left,
				title.Render(m.Title),
				content,
				statusBar,
			),
		),
	)
}

func (m *JWTHeaderModel) SetHeight(height int) {
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

func (m *JWTHeaderModel) SetWidth(width int) {
	m.Width = width - 2
	// Account for borders (typically 2 columns total)
	if m.EditingMode {
		m.TextArea.SetWidth(width - 2)
	} else {
		m.Viewport.SetWidth(width - 2)
	}
}

func (m *JWTHeaderModel) SetData(content string) {
	m.Content = content
	if m.EditingMode {
		m.TextArea.SetValue(content)
	} else {
		m.Viewport.SetContent(content)
	}
}

func (m JWTHeaderModel) GetData() string {
	if m.EditingMode {
		return m.TextArea.Value()
	} else {
		return m.Content
	}
}

func (m *JWTHeaderModel) SetEditingMode(editing bool) {
	m.EditingMode = editing
}

func (m JWTHeaderModel) IsEditing() bool {
	return m.EditingMode
}

func (m *JWTHeaderModel) SetError(error string) {
	m.Error = error
}

func (m *JWTHeaderModel) Blur() {
	m.Focused = false
	if m.EditingMode {
		m.TextArea.Blur()
	}
}

// ================================================================================

type JWTPayloadModel struct {
	ZoneID      string
	TextArea    textarea.Model
	Viewport    viewport.Model
	Title       string
	EditingMode bool
	Focused     bool
	ElementID   Element
	Height      int
	Width       int
	Error       string
	Content     string
}

func NewJWTPayloadModel() JWTPayloadModel {
	textArea := textarea.New()
	textArea.Placeholder = "Enter payload JSON here..."
	textArea.Prompt = ""
	textArea.ShowLineNumbers = false

	viewportModel := viewport.New()
	viewportModel.SoftWrap = true

	return JWTPayloadModel{
		TextArea:    textArea,
		Viewport:    viewportModel,
		Title:       "",
		EditingMode: true, // Default to editing mode for encoder
		Focused:     false,
		ElementID:   "",
		Height:      0,
		Width:       0,
		Error:       "",
		Content:     "",
	}
}

func (m JWTPayloadModel) Init() tea.Cmd {
	return nil
}

func (m JWTPayloadModel) Update(msg tea.Msg) (JWTPayloadModel, tea.Cmd) {
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

	var cmd tea.Cmd
	if m.Focused && m.EditingMode {
		m.TextArea, cmd = m.TextArea.Update(msg)
	} else if m.Focused && !m.EditingMode {
		m.Viewport, cmd = m.Viewport.Update(msg)
	}
	return m, cmd
}

func (m JWTPayloadModel) View() string {
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
		m.ZoneID,
		box.Render(
			lipgloss.JoinVertical(lipgloss.Left,
				title.Render(m.Title),
				content,
				statusBar,
			),
		),
	)
}

func (m *JWTPayloadModel) SetHeight(height int) {
	m.Height = height
	// Account for title, borders, and status bar (typically 3-4 lines total)
	internalHeight := max(height-4, 1)

	if m.EditingMode {
		m.TextArea.SetHeight(internalHeight)
	} else {
		m.Viewport.SetHeight(internalHeight)
	}
}

func (m *JWTPayloadModel) SetWidth(width int) {
	m.Width = width - 2
	// Account for borders (typically 2 columns total)
	if m.EditingMode {
		m.TextArea.SetWidth(width - 2)
	} else {
		m.Viewport.SetWidth(width - 2)
	}
}

func (m *JWTPayloadModel) SetData(content string) {
	m.Content = content
	if m.EditingMode {
		m.TextArea.SetValue(content)
	} else {
		m.Viewport.SetContent(content)
	}
}

func (m JWTPayloadModel) GetData() string {
	if m.EditingMode {
		return m.TextArea.Value()
	} else {
		return m.Content
	}
}

func (m *JWTPayloadModel) SetEditingMode(editing bool) {
	m.EditingMode = editing
}

func (m JWTPayloadModel) IsEditing() bool {
	return m.EditingMode
}

func (m *JWTPayloadModel) SetError(error string) {
	m.Error = error
}

func (m *JWTPayloadModel) Blur() {
	m.Focused = false
	if m.EditingMode {
		m.TextArea.Blur()
	}
}
