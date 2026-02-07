package main

import (
	"charm.land/bubbles/v2/textarea"
	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

// JWTTokenModel represents the JWT token field that can be in input or output mode
type JWTTokenModel struct {
	TextArea    textarea.Model
	Focused     bool
	Result      *JWTDecodeResult
	EditingMode bool
	Height      int
	Width       int
}

// NewJWTModel creates a new JWT model
func NewJWTTokenModel() JWTTokenModel {
	textArea := textarea.New()
	textArea.Placeholder = PlaceholderJWT
	textArea.Prompt = ""
	textArea.ShowLineNumbers = false

	return JWTTokenModel{
		TextArea:    textArea,
		Focused:     false,
		Result:      nil,
		EditingMode: true, // Default to editing mode,
		Height:      0,
		Width:       0,
	}
}

// Init initializes the JWT model
func (m JWTTokenModel) Init() tea.Cmd {
	return nil
}

// Update handles messages for the JWT model
func (m JWTTokenModel) Update(msg tea.Msg) (JWTTokenModel, tea.Cmd) {
	var cmd tea.Cmd
	m.TextArea, cmd = m.TextArea.Update(msg)
	return m, cmd
}

// View renders the JWT model
func (m JWTTokenModel) View() string {
	title := styleTitle
	box := styleBox
	if m.Focused {
		title = styleTitleSelected
		box = styleBoxActive
	}

	statusBar := styleStatus.Width(m.Width).Render("")
	if m.Result != nil {
		if m.Result.IsTokenValid {
			statusBar = styleStatusSuccess.Width(m.Width).Render(StatusValidJWT)
		} else {
			statusBar = styleStatusError.Width(m.Width).Render(StatusInvalidToken)
		}
	}

	return box.Render(
		lipgloss.JoinVertical(lipgloss.Left,
			title.Render(TitleJWTToken),
			m.TextArea.View(),
			statusBar,
		),
	)
}

// SetHeight sets the height of the text area
func (m *JWTTokenModel) SetHeight(height int) {
	m.Height = height
	// Account for title, borders, and status bar (typically 3-4 lines total)
	internalHeight := height - 4
	if internalHeight < 1 {
		internalHeight = 1
	}
	m.TextArea.SetHeight(internalHeight)
}

// SetWidth sets the width of the text area
func (m *JWTTokenModel) SetWidth(width int) {
	m.Width = width - 2
	// Account for borders (typically 2 columns total)
	m.TextArea.SetWidth(width - 2)
}

// GetToken returns the current token value
func (m JWTTokenModel) GetToken() string {
	return m.TextArea.Value()
}

// SetToken sets the token value
func (m *JWTTokenModel) SetToken(token string) {
	m.TextArea.SetValue(token)
}

// Focus focuses the text area
func (m *JWTTokenModel) Focus() {
	m.TextArea.Focus()
}

// Blur blurs the text area
func (m *JWTTokenModel) Blur() {
	m.TextArea.Blur()
}

// SetEditingMode sets whether the model is in editing mode
func (m *JWTTokenModel) SetEditingMode(editing bool) {
	m.EditingMode = editing
}

// IsEditing returns whether the model is in editing mode
func (m JWTTokenModel) IsEditing() bool {
	return m.EditingMode
}

// JWTSecretModel represents the secret field that can be in input or output mode
type JWTSecretModel struct {
	TextArea    textarea.Model
	Focused     bool
	Result      *JWTDecodeResult
	EditingMode bool
	Height      int
	Width       int
}

// NewSecretModel creates a new secret model
func NewJWTSecretModel() JWTSecretModel {
	textArea := textarea.New()
	textArea.Placeholder = PlaceholderSecret
	textArea.Prompt = ""
	textArea.ShowLineNumbers = false

	return JWTSecretModel{
		TextArea:    textArea,
		Focused:     false,
		Result:      nil,
		EditingMode: true, // Default to editing mode
		Height:      0,
		Width:       0,
	}
}

// Init initializes the secret model
func (m JWTSecretModel) Init() tea.Cmd {
	return nil
}

// Update handles messages for the secret model
func (m JWTSecretModel) Update(msg tea.Msg) (JWTSecretModel, tea.Cmd) {
	var cmd tea.Cmd
	m.TextArea, cmd = m.TextArea.Update(msg)
	return m, cmd
}

// View renders the secret model
func (m JWTSecretModel) View() string {
	title := styleTitle
	box := styleBox
	if m.Focused {
		title = styleTitleSelected
		box = styleBoxActive
	}

	width := lipgloss.Width(m.TextArea.View())

	statusBar := styleStatus.Width(width).Render("")
	if m.Result != nil {
		if m.Result.IsSignatureValid {
			statusBar = styleStatusSuccess.Width(width).Render(StatusSignatureVerified)
		} else {
			statusBar = styleStatusError.Width(width).Render(StatusSignatureVerificationFailed)
		}
	}

	return box.Render(
		lipgloss.JoinVertical(lipgloss.Left,
			title.Render(TitleSecret),
			m.TextArea.View(),
			statusBar,
		),
	)
}

// SetHeight sets the height of the text area
func (m *JWTSecretModel) SetHeight(height int) {
	m.Height = height
	// Account for title, borders, and status bar (typically 3-4 lines total)
	internalHeight := height - 4
	if internalHeight < 1 {
		internalHeight = 1
	}
	m.TextArea.SetHeight(internalHeight)
}

// SetWidth sets the width of the text area
func (m *JWTSecretModel) SetWidth(width int) {
	m.Width = width - 2
	// Account for borders (typically 2 columns total)
	m.TextArea.SetWidth(width - 2)
}

// GetSecret returns the current secret value
func (m JWTSecretModel) GetSecret() string {
	return m.TextArea.Value()
}

// SetSecret sets the secret value
func (m *JWTSecretModel) SetSecret(secret string) {
	m.TextArea.SetValue(secret)
}

// Focus focuses the text area
func (m *JWTSecretModel) Focus() {
	m.TextArea.Focus()
}

// Blur blurs the text area
func (m *JWTSecretModel) Blur() {
	m.TextArea.Blur()
}

// SetEditingMode sets whether the model is in editing mode
func (m *JWTSecretModel) SetEditingMode(editing bool) {
	m.EditingMode = editing
}

// IsEditing returns whether the model is in editing mode
func (m JWTSecretModel) IsEditing() bool {
	return m.EditingMode
}

// JWTHeaderModel represents the header display that can be in input or output mode
type JWTHeaderModel struct {
	Viewport    viewport.Model
	Title       string
	EditingMode bool
	Focused     bool
	Height      int
	Width       int
}

// NewHeaderModel creates a new header model
func NewJWTHeaderModel() JWTHeaderModel {
	viewportModel := viewport.New()
	viewportModel.SoftWrap = true

	return JWTHeaderModel{
		Viewport:    viewportModel,
		Title:       "",
		EditingMode: false, // Default to viewing mode
		Focused:     false,
		Height:      0,
		Width:       0,
	}
}

// Init initializes the header model
func (m JWTHeaderModel) Init() tea.Cmd {
	return nil
}

// Update handles messages for the header model
func (m JWTHeaderModel) Update(msg tea.Msg) (JWTHeaderModel, tea.Cmd) {
	if !m.Focused {
		return m, nil
	}

	var cmd tea.Cmd
	m.Viewport, cmd = m.Viewport.Update(msg)
	return m, cmd
}

// View renders the header model
func (m JWTHeaderModel) View() string {
	title := styleTitle
	box := styleBox
	if m.Focused {
		title = styleTitleSelected
		box = styleBoxActive
	}

	width := lipgloss.Width(m.Viewport.View())

	statusBar := styleStatus.Width(width).Render("")

	return box.Render(
		lipgloss.JoinVertical(lipgloss.Left,
			title.Render(m.Title),
			m.Viewport.View(),
			statusBar,
		),
	)
}

// SetHeight sets the height of the viewport
func (m *JWTHeaderModel) SetHeight(height int) {
	m.Height = height
	// Account for title, borders, and status bar (typically 3-4 lines total)
	internalHeight := height - 4
	if internalHeight < 1 {
		internalHeight = 1
	}
	m.Viewport.SetHeight(internalHeight)
}

// SetWidth sets the width of the viewport
func (m *JWTHeaderModel) SetWidth(width int) {
	m.Width = width - 2
	// Account for borders (typically 2 columns total)
	m.Viewport.SetWidth(width - 2)
}

// SetData sets the content of the viewport
func (m *JWTHeaderModel) SetData(content string) {
	m.Viewport.SetContent(content)
}

// SetEditingMode sets whether the model is in editing mode
func (m *JWTHeaderModel) SetEditingMode(editing bool) {
	m.EditingMode = editing
}

// IsEditing returns whether the model is in editing mode
func (m JWTHeaderModel) IsEditing() bool {
	return m.EditingMode
}

// Focus focuses the header model
func (m *JWTHeaderModel) Focus() {
	m.Focused = true
}

// Blur blurs the header model
func (m *JWTHeaderModel) Blur() {
	m.Focused = false
}

// JWTPayloadModel represents the payload display that can be in input or output mode
type JWTPayloadModel struct {
	Viewport    viewport.Model
	Title       string
	EditingMode bool
	Focused     bool
	Height      int
	Width       int
}

// NewPayloadModel creates a new payload model
func NewJWTPayloadModel() JWTPayloadModel {
	viewportModel := viewport.New()
	viewportModel.SoftWrap = true

	return JWTPayloadModel{
		Viewport:    viewportModel,
		Title:       "",
		EditingMode: false, // Default to viewing mode
		Focused:     false,
		Height:      0,
		Width:       0,
	}
}

// Init initializes the payload model
func (m JWTPayloadModel) Init() tea.Cmd {
	return nil
}

// Update handles messages for the payload model
func (m JWTPayloadModel) Update(msg tea.Msg) (JWTPayloadModel, tea.Cmd) {
	if !m.Focused {
		return m, nil
	}

	var cmd tea.Cmd
	m.Viewport, cmd = m.Viewport.Update(msg)
	return m, cmd
}

// View renders the payload model
func (m JWTPayloadModel) View() string {
	title := styleTitle
	box := styleBox
	if m.Focused {
		title = styleTitleSelected
		box = styleBoxActive
	}

	width := lipgloss.Width(m.Viewport.View())

	statusBar := styleStatus.Width(width).Render("")

	return box.Render(
		lipgloss.JoinVertical(lipgloss.Left,
			title.Render(m.Title),
			m.Viewport.View(),
			statusBar,
		),
	)
}

// SetHeight sets the height of the viewport
func (m *JWTPayloadModel) SetHeight(height int) {
	m.Height = height
	// Account for title, borders, and status bar (typically 3-4 lines total)
	internalHeight := max(height-4, 1)
	m.Viewport.SetHeight(internalHeight)
}

// SetWidth sets the width of the viewport
func (m *JWTPayloadModel) SetWidth(width int) {
	m.Width = width - 2
	// Account for borders (typically 2 columns total)
	m.Viewport.SetWidth(width - 2)
}

// SetData sets the content of the viewport
func (m *JWTPayloadModel) SetData(content string) {
	m.Viewport.SetContent(content)
}

// SetEditingMode sets whether the model is in editing mode
func (m *JWTPayloadModel) SetEditingMode(editing bool) {
	m.EditingMode = editing
}

// IsEditing returns whether the model is in editing mode
func (m JWTPayloadModel) IsEditing() bool {
	return m.EditingMode
}

// Focus focuses the payload model
func (m *JWTPayloadModel) Focus() {
	m.Focused = true
}

// Blur blurs the payload model
func (m *JWTPayloadModel) Blur() {
	m.Focused = false
}
