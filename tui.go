package main

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

func NewBubbleTeamModel() BubbleTeaModel {
	decoderJWTModel := NewJWTModel()
	decoderSecretModel := NewSecretModel()
	decoderHeaderModel := NewHeaderModel()
	decoderHeaderModel.Title = TitleDecodedHeader
	decoderPayloadModel := NewPayloadModel()
	decoderPayloadModel.Title = TitleDecodedPayload

	return BubbleTeaModel{
		SelectedView:        ViewJWTDecoder,
		FocusedElement:      ElementDecoderJWTTextArea,
		DecoderJWTModel:     decoderJWTModel,
		DecoderSecretModel:  decoderSecretModel,
		DecoderHeaderModel:  decoderHeaderModel,
		DecoderPayloadModel: decoderPayloadModel,
	}
}

type BubbleTeaModel struct {
	WindowSize tea.WindowSizeMsg

	SelectedView   View
	FocusedElement Element

	// UI Elements
	DecoderJWTModel     JWTModel
	DecoderSecretModel  SecretModel
	DecoderHeaderModel  HeaderModel
	DecoderPayloadModel PayloadModel
	DecodeResult        *JWTDecodeResult
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

		// Account for header height (header takes 1 line + margin)
		headerHeight := 2
		footerHeight := 1

		styleTitle = styleTitle.Width((msg.Width / 2) - 2)
		styleTitleSelected = styleTitleSelected.Width((msg.Width / 2) - 2)

		// Adjust available height by subtracting header height
		availableHeight := msg.Height - headerHeight - footerHeight - 3

		// Set dimensions for JWT model
		m.DecoderJWTModel.SetHeight((2 * availableHeight / 3))
		m.DecoderJWTModel.SetWidth((msg.Width / 2))
		m.DecoderJWTModel.Focus()

		// Set dimensions for secret model
		m.DecoderSecretModel.SetHeight((availableHeight / 3))
		m.DecoderSecretModel.SetWidth((msg.Width / 2))

		// Set dimensions for payload model
		m.DecoderPayloadModel.SetHeight((2 * availableHeight / 3))
		m.DecoderPayloadModel.SetWidth((msg.Width / 2))

		// Set dimensions for header model
		m.DecoderHeaderModel.SetHeight((availableHeight / 3))
		m.DecoderHeaderModel.SetWidth((msg.Width / 2))

	case tea.KeyMsg:
		switch msg.Key().String() {
		// Quit Program
		case KeyQuit, KeyQuitAlt:
			return m, tea.Quit
		case KeyFocusToken:
			m.FocusedElement = ElementDecoderJWTTextArea
			m.DecoderJWTModel.Focus()
			m.DecoderSecretModel.Blur()
		case KeyFocusSecret:
			m.FocusedElement = ElementDecoderSecretTextArea
			m.DecoderSecretModel.Focus()
			m.DecoderJWTModel.Blur()
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
			m.DecoderHeaderModel.SetData(m.DecodeResult.JsonMarshaledHeader())
			m.DecoderPayloadModel.SetData(m.DecodeResult.JsonMarshaledClaims())
		} else {
			m.DecoderHeaderModel.SetData("")
			m.DecoderPayloadModel.SetData("")
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
			m.DecoderHeaderModel.View(),
			m.DecoderPayloadModel.View(),
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

		fullContent := lipgloss.JoinVertical(lipgloss.Center,
			header,
			content,
		)

		v.SetContent(fullContent)
	}

	return v
}
