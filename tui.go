package main

import (
	"charm.land/bubbles/v2/help"
	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

func NewBubbleTeamModel() BubbleTeaModel {
	decoderJWTModel := NewJWTTokenModel()
	decoderSecretModel := NewJWTSecretModel()
	decoderHeaderModel := NewJWTHeaderModel()
	decoderHeaderModel.Title = TitleDecodedHeader
	decoderPayloadModel := NewJWTPayloadModel()
	decoderPayloadModel.Title = TitleDecodedPayload
	decoderHelpModel := help.New()

	return BubbleTeaModel{
		SelectedView:           ViewJWTDecoder,
		FocusedElement:         ElementDecoderJWTTextArea,
		DecoderJWTModel:        decoderJWTModel,
		DecoderSecretModel:     decoderSecretModel,
		DecoderJWTHeaderModel:  decoderHeaderModel,
		DecoderJWTPayloadModel: decoderPayloadModel,
		HelpModel:              decoderHelpModel,
	}
}

type BubbleTeaModel struct {
	WindowSize tea.WindowSizeMsg

	SelectedView   View
	FocusedElement Element

	// UI Elements
	DecoderJWTModel        JWTTokenModel
	DecoderSecretModel     JWTSecretModel
	DecoderJWTHeaderModel  JWTHeaderModel
	DecoderJWTPayloadModel JWTPayloadModel
	DecodeResult           *JWTDecodeResult

	HelpModel help.Model
}

func (m BubbleTeaModel) Init() tea.Cmd {
	return nil
}

func (m BubbleTeaModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.WindowSize = msg

		headerHeight := 2
		footerHeight := 1

		styleTitle = styleTitle.Width((msg.Width / 2) - 2)
		styleTitleSelected = styleTitleSelected.Width((msg.Width / 2) - 2)

		// Available height is less than total height becayse we need to account
		// for header and footer.
		availableHeight := msg.Height - headerHeight - footerHeight - 5

		m.DecoderJWTModel.SetHeight((availableHeight / 2))
		m.DecoderJWTModel.SetWidth((msg.Width / 2))
		m.DecoderJWTModel.Focus()

		m.DecoderSecretModel.SetHeight((availableHeight / 2))
		m.DecoderSecretModel.SetWidth((msg.Width / 2))

		m.DecoderJWTPayloadModel.SetHeight((availableHeight / 2))
		m.DecoderJWTPayloadModel.SetWidth((msg.Width / 2))

		m.DecoderJWTHeaderModel.SetHeight((availableHeight / 2))
		m.DecoderJWTHeaderModel.SetWidth((msg.Width / 2))

		m.HelpModel.SetWidth(msg.Width)

	case tea.KeyMsg:
		switch msg.Key().String() {
		// Quit Program
		case KeyQuit, KeyQuitAlt:
			return m, tea.Quit
		case KeyFocusToken:
			m.FocusedElement = ElementDecoderJWTTextArea
			m.DecoderJWTModel.Focus()
			m.DecoderSecretModel.Blur()
			m.DecoderJWTHeaderModel.Blur()
			m.DecoderJWTPayloadModel.Blur()
		case KeyFocusSecret:
			m.FocusedElement = ElementDecoderSecretTextArea
			m.DecoderSecretModel.Focus()
			m.DecoderJWTModel.Blur()
			m.DecoderJWTHeaderModel.Blur()
			m.DecoderJWTPayloadModel.Blur()
		case KeyFocusHeader:
			m.FocusedElement = ElementDecoderHeaderTextArea
			m.DecoderJWTHeaderModel.Focus()
			m.DecoderJWTModel.Blur()
			m.DecoderSecretModel.Blur()
			m.DecoderJWTPayloadModel.Blur()
		case KeyFocusPayload:
			m.FocusedElement = ElementDecoderPayloadTextArea
			m.DecoderJWTPayloadModel.Focus()
			m.DecoderJWTModel.Blur()
			m.DecoderSecretModel.Blur()
			m.DecoderJWTHeaderModel.Blur()
		}
	}

	// Update JWT model
	m.DecoderJWTModel.Focused = m.FocusedElement == ElementDecoderJWTTextArea
	m.DecoderJWTModel, cmd = m.DecoderJWTModel.Update(msg)
	cmds = append(cmds, cmd)

	// Update secret model
	m.DecoderSecretModel.Focused = m.FocusedElement == ElementDecoderSecretTextArea
	m.DecoderSecretModel, cmd = m.DecoderSecretModel.Update(msg)
	cmds = append(cmds, cmd)

	// Update header model
	m.DecoderJWTHeaderModel.Focused = m.FocusedElement == ElementDecoderHeaderTextArea
	m.DecoderJWTHeaderModel, cmd = m.DecoderJWTHeaderModel.Update(msg)
	cmds = append(cmds, cmd)

	// Update payload model
	m.DecoderJWTPayloadModel.Focused = m.FocusedElement == ElementDecoderPayloadTextArea
	m.DecoderJWTPayloadModel, cmd = m.DecoderJWTPayloadModel.Update(msg)
	cmds = append(cmds, cmd)

	// Update header and payload models with decode results
	m.DecodeResult = nil
	token := m.DecoderJWTModel.GetToken()
	secret := m.DecoderSecretModel.GetSecret()

	if token != "" {
		m.DecodeResult = JWTDecodeToken(token, secret)

		// Update the result in the models
		m.DecoderJWTModel.Result = m.DecodeResult
		m.DecoderSecretModel.Result = m.DecodeResult

		if m.DecodeResult.Token != nil {
			m.DecoderJWTHeaderModel.SetData(m.DecodeResult.JsonMarshaledHeader())
			m.DecoderJWTPayloadModel.SetData(m.DecodeResult.JsonMarshaledClaims())
		} else {
			m.DecoderJWTHeaderModel.SetData("")
			m.DecoderJWTPayloadModel.SetData("")
		}
	}

	return m, tea.Batch(cmds...)
}

func (m BubbleTeaModel) View() tea.View {
	v := tea.View{
		AltScreen: true,
	}

	switch m.SelectedView {
	case ViewJWTDecoder:
		pane1 := lipgloss.JoinVertical(lipgloss.Left,
			m.DecoderJWTModel.View(),
			m.DecoderSecretModel.View(),
		)

		pane2 := lipgloss.JoinVertical(lipgloss.Left,
			m.DecoderJWTHeaderModel.View(),
			m.DecoderJWTPayloadModel.View(),
		)

		content := lipgloss.JoinHorizontal(lipgloss.Left,
			pane1,
			pane2,
		)

		decoderStyle, encoderStyle := styleInactiveScreen, styleInactiveScreen

		switch m.SelectedView {
		case ViewJWTDecoder:
			decoderStyle = styleActiveScreen
		case ViewJWTEncoder:
			encoderStyle = styleActiveScreen
		}

		header := styleHeader.Width(m.WindowSize.Width).
			Render(decoderStyle.Render(TitleDecoder) + styleInactiveScreen.Render(" | ") + encoderStyle.Render(TitleEncoder))

		footer := lipgloss.NewStyle().Padding(0, 1, 0, 1).MarginTop(1).Render(m.HelpModel.View(m))

		fullContent := lipgloss.JoinVertical(lipgloss.Left,
			header,
			content,
			footer,
		)

		v.SetContent(fullContent)
	}

	return v
}

func (m BubbleTeaModel) ShortHelp() []key.Binding {
	switch m.SelectedView {
	case ViewJWTDecoder:
		return []key.Binding{
			key.NewBinding(key.WithKeys(KeyQuit, KeyQuitAlt), key.WithHelp(KeyQuit, "Quit")),
			key.NewBinding(key.WithKeys(KeyFocusToken), key.WithHelp(KeyFocusToken, "Focus Token")),
			key.NewBinding(key.WithKeys(KeyFocusSecret), key.WithHelp(KeyFocusSecret, "Focus Secret")),
			key.NewBinding(key.WithKeys(KeyFocusHeader), key.WithHelp(KeyFocusHeader, "Focus Header")),
			key.NewBinding(key.WithKeys(KeyFocusPayload), key.WithHelp(KeyFocusPayload, "Focus Payload")),
		}
	}

	return []key.Binding{}
}

func (m BubbleTeaModel) FullHelp() [][]key.Binding {
	return [][]key.Binding{}
}
