package main

import (
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

	return BubbleTeaModel{
		SelectedView:           ViewJWTDecoder,
		FocusedElement:         ElementDecoderJWTTextArea,
		DecoderJWTModel:        decoderJWTModel,
		DecoderSecretModel:     decoderSecretModel,
		DecoderJWTHeaderModel:  decoderHeaderModel,
		DecoderJWTPayloadModel: decoderPayloadModel,
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
		m.DecoderJWTPayloadModel.SetHeight((2 * availableHeight / 3))
		m.DecoderJWTPayloadModel.SetWidth((msg.Width / 2))

		// Set dimensions for header model
		m.DecoderJWTHeaderModel.SetHeight((availableHeight / 3))
		m.DecoderJWTHeaderModel.SetWidth((msg.Width / 2))

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

		fullContent := lipgloss.JoinVertical(lipgloss.Center,
			header,
			content,
		)

		v.SetContent(fullContent)
	}

	return v
}
