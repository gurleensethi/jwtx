package main

import (
	"charm.land/bubbles/v2/textarea"
	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

// JWTModel represents the JWT token field that can be in input or output mode
type JWTModel struct {
	TextArea    textarea.Model
	Focused     bool
	Result      *JWTDecodeResult
	EditingMode bool
	Height      int
	Width       int
}

// NewJWTModel creates a new JWT model
func NewJWTModel() JWTModel {
	textArea := textarea.New()
	textArea.Placeholder = PlaceholderJWT
	textArea.Prompt = ""
	textArea.ShowLineNumbers = false

	return JWTModel{
		TextArea:    textArea,
		Focused:     false,
		Result:      nil,
		EditingMode: true, // Default to editing mode,
		Height:      0,
		Width:       0,
	}
}

// Init initializes the JWT model
func (m JWTModel) Init() tea.Cmd {
	return nil
}

// Update handles messages for the JWT model
func (m JWTModel) Update(msg tea.Msg) (JWTModel, tea.Cmd) {
	var cmd tea.Cmd
	m.TextArea, cmd = m.TextArea.Update(msg)
	return m, cmd
}

// View renders the JWT model
func (m JWTModel) View() string {
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
func (m *JWTModel) SetHeight(height int) {
	m.Height = height
	// Account for title, borders, and status bar (typically 3-4 lines total)
	internalHeight := height - 4
	if internalHeight < 1 {
		internalHeight = 1
	}
	m.TextArea.SetHeight(internalHeight)
}

// SetWidth sets the width of the text area
func (m *JWTModel) SetWidth(width int) {
	m.Width = width - 2
	// Account for borders (typically 2 columns total)
	m.TextArea.SetWidth(width - 2)
}

// GetToken returns the current token value
func (m JWTModel) GetToken() string {
	return m.TextArea.Value()
}

// SetToken sets the token value
func (m *JWTModel) SetToken(token string) {
	m.TextArea.SetValue(token)
}

// Focus focuses the text area
func (m *JWTModel) Focus() {
	m.TextArea.Focus()
}

// Blur blurs the text area
func (m *JWTModel) Blur() {
	m.TextArea.Blur()
}

// SetEditingMode sets whether the model is in editing mode
func (m *JWTModel) SetEditingMode(editing bool) {
	m.EditingMode = editing
}

// IsEditing returns whether the model is in editing mode
func (m JWTModel) IsEditing() bool {
	return m.EditingMode
}

// SecretModel represents the secret field that can be in input or output mode
type SecretModel struct {
	TextArea    textarea.Model
	Focused     bool
	Result      *JWTDecodeResult
	EditingMode bool
	Height      int
	Width       int
}

// NewSecretModel creates a new secret model
func NewSecretModel() SecretModel {
	textArea := textarea.New()
	textArea.Placeholder = PlaceholderSecret
	textArea.Prompt = ""
	textArea.ShowLineNumbers = false

	return SecretModel{
		TextArea:    textArea,
		Focused:     false,
		Result:      nil,
		EditingMode: true, // Default to editing mode
		Height:      0,
		Width:       0,
	}
}

// Init initializes the secret model
func (m SecretModel) Init() tea.Cmd {
	return nil
}

// Update handles messages for the secret model
func (m SecretModel) Update(msg tea.Msg) (SecretModel, tea.Cmd) {
	var cmd tea.Cmd
	m.TextArea, cmd = m.TextArea.Update(msg)
	return m, cmd
}

// View renders the secret model
func (m SecretModel) View() string {
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
func (m *SecretModel) SetHeight(height int) {
	m.Height = height
	// Account for title, borders, and status bar (typically 3-4 lines total)
	internalHeight := height - 4
	if internalHeight < 1 {
		internalHeight = 1
	}
	m.TextArea.SetHeight(internalHeight)
}

// SetWidth sets the width of the text area
func (m *SecretModel) SetWidth(width int) {
	m.Width = width - 2
	// Account for borders (typically 2 columns total)
	m.TextArea.SetWidth(width - 2)
}

// GetSecret returns the current secret value
func (m SecretModel) GetSecret() string {
	return m.TextArea.Value()
}

// SetSecret sets the secret value
func (m *SecretModel) SetSecret(secret string) {
	m.TextArea.SetValue(secret)
}

// Focus focuses the text area
func (m *SecretModel) Focus() {
	m.TextArea.Focus()
}

// Blur blurs the text area
func (m *SecretModel) Blur() {
	m.TextArea.Blur()
}

// SetEditingMode sets whether the model is in editing mode
func (m *SecretModel) SetEditingMode(editing bool) {
	m.EditingMode = editing
}

// IsEditing returns whether the model is in editing mode
func (m SecretModel) IsEditing() bool {
	return m.EditingMode
}

// HeaderModel represents the header display that can be in input or output mode
type HeaderModel struct {
	Viewport    viewport.Model
	Title       string
	EditingMode bool
	Height      int
	Width       int
}

// NewHeaderModel creates a new header model
func NewHeaderModel() HeaderModel {
	viewportModel := viewport.New()
	viewportModel.SoftWrap = true

	return HeaderModel{
		Viewport:    viewportModel,
		Title:       "",
		EditingMode: false, // Default to viewing mode
		Height:      0,
		Width:       0,
	}
}

// Init initializes the header model
func (m HeaderModel) Init() tea.Cmd {
	return nil
}

// Update handles messages for the header model
func (m HeaderModel) Update(msg tea.Msg) (HeaderModel, tea.Cmd) {
	var cmd tea.Cmd
	m.Viewport, cmd = m.Viewport.Update(msg)
	return m, cmd
}

// View renders the header model
func (m HeaderModel) View() string {
	width := lipgloss.Width(m.Viewport.View())

	statusBar := styleStatus.Width(width).Render("")

	return styleBox.Render(
		lipgloss.JoinVertical(lipgloss.Left,
			styleTitle.Render(m.Title),
			m.Viewport.View(),
			statusBar,
		),
	)
}

// SetHeight sets the height of the viewport
func (m *HeaderModel) SetHeight(height int) {
	m.Height = height
	// Account for title, borders, and status bar (typically 3-4 lines total)
	internalHeight := height - 4
	if internalHeight < 1 {
		internalHeight = 1
	}
	m.Viewport.SetHeight(internalHeight)
}

// SetWidth sets the width of the viewport
func (m *HeaderModel) SetWidth(width int) {
	m.Width = width - 2
	// Account for borders (typically 2 columns total)
	m.Viewport.SetWidth(width - 2)
}

// SetData sets the content of the viewport
func (m *HeaderModel) SetData(content string) {
	m.Viewport.SetContent(content)
}

// SetEditingMode sets whether the model is in editing mode
func (m *HeaderModel) SetEditingMode(editing bool) {
	m.EditingMode = editing
}

// IsEditing returns whether the model is in editing mode
func (m HeaderModel) IsEditing() bool {
	return m.EditingMode
}

// PayloadModel represents the payload display that can be in input or output mode
type PayloadModel struct {
	Viewport    viewport.Model
	Title       string
	EditingMode bool
	Height      int
	Width       int
}

// NewPayloadModel creates a new payload model
func NewPayloadModel() PayloadModel {
	viewportModel := viewport.New()
	viewportModel.SoftWrap = true

	return PayloadModel{
		Viewport:    viewportModel,
		Title:       "",
		EditingMode: false, // Default to viewing mode
		Height:      0,
		Width:       0,
	}
}

// Init initializes the payload model
func (m PayloadModel) Init() tea.Cmd {
	return nil
}

// Update handles messages for the payload model
func (m PayloadModel) Update(msg tea.Msg) (PayloadModel, tea.Cmd) {
	var cmd tea.Cmd
	m.Viewport, cmd = m.Viewport.Update(msg)
	return m, cmd
}

// View renders the payload model
func (m PayloadModel) View() string {
	width := lipgloss.Width(m.Viewport.View())

	statusBar := styleStatus.Width(width).Render("")

	return styleBox.Render(
		lipgloss.JoinVertical(lipgloss.Left,
			styleTitle.Render(m.Title),
			m.Viewport.View(),
			statusBar,
		),
	)
}

// SetHeight sets the height of the viewport
func (m *PayloadModel) SetHeight(height int) {
	m.Height = height
	// Account for title, borders, and status bar (typically 3-4 lines total)
	internalHeight := max(height-4, 1)
	m.Viewport.SetHeight(internalHeight)
}

// SetWidth sets the width of the viewport
func (m *PayloadModel) SetWidth(width int) {
	m.Width = width - 2
	// Account for borders (typically 2 columns total)
	m.Viewport.SetWidth(width - 2)
}

// SetData sets the content of the viewport
func (m *PayloadModel) SetData(content string) {
	m.Viewport.SetContent(content)
}

// SetEditingMode sets whether the model is in editing mode
func (m *PayloadModel) SetEditingMode(editing bool) {
	m.EditingMode = editing
}

// IsEditing returns whether the model is in editing mode
func (m PayloadModel) IsEditing() bool {
	return m.EditingMode
}
