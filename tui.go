package main

import (
	"encoding/json"

	"github.com/golang-jwt/jwt/v5"

	"charm.land/bubbles/v2/help"
	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

func NewBubbleTeamModel() BubbleTeaModel {
	decoderJWTModel := NewJWTTokenModel()
	decoderJWTModel.ElementID = ElementDecoderJWTTextArea
	decoderSecretModel := NewJWTSecretModel()
	decoderSecretModel.ElementID = ElementDecoderSecretTextArea
	decoderHeaderModel := NewJWTHeaderModel()
	decoderHeaderModel.ElementID = ElementDecoderHeaderTextArea
	decoderHeaderModel.Title = TitleDecodedHeader
	decoderPayloadModel := NewJWTPayloadModel()
	decoderPayloadModel.ElementID = ElementDecoderPayloadTextArea
	decoderPayloadModel.Title = TitleDecodedPayload

	encoderJWTModel := NewJWTTokenModel()
	encoderJWTModel.ElementID = ElementEncoderJWTTextArea
	encoderJWTModel.SetEditingMode(false) // In encoder view, JWT is output
	encoderSecretModel := NewJWTSecretModel()
	encoderSecretModel.ElementID = ElementEncoderSecretTextArea
	encoderHeaderModel := NewJWTHeaderModel()
	encoderHeaderModel.ElementID = ElementEncoderHeaderTextArea
	encoderHeaderModel.Title = TitleEncoderHeader
	encoderHeaderModel.SetEditingMode(true) // In encoder view, header is input
	encoderPayloadModel := NewJWTPayloadModel()
	encoderPayloadModel.ElementID = ElementEncoderPayloadTextArea
	encoderPayloadModel.Title = TitleEncoderPayload
	encoderPayloadModel.SetEditingMode(true) // In encoder view, payload is input

	decoderHelpModel := help.New()

	return BubbleTeaModel{
		SelectedView:           ViewJWTDecoder,
		FocusedElement:         ElementDecoderJWTTextArea,
		DecoderJWTModel:        decoderJWTModel,
		DecoderSecretModel:     decoderSecretModel,
		DecoderJWTHeaderModel:  decoderHeaderModel,
		DecoderJWTPayloadModel: decoderPayloadModel,
		EncoderJWTModel:        encoderJWTModel,
		EncoderSecretModel:     encoderSecretModel,
		EncoderJWTHeaderModel:  encoderHeaderModel,
		EncoderJWTPayloadModel: encoderPayloadModel,
		EncodeResult:           nil,
		HelpModel:              decoderHelpModel,
	}
}

type BubbleTeaModel struct {
	WindowSize tea.WindowSizeMsg

	SelectedView   View
	FocusedElement Element

	DecoderJWTModel        JWTTokenModel
	DecoderSecretModel     JWTSecretModel
	DecoderJWTHeaderModel  JWTHeaderModel
	DecoderJWTPayloadModel JWTPayloadModel
	DecodeResult           *JWTDecodeResult

	EncoderJWTModel        JWTTokenModel
	EncoderSecretModel     JWTSecretModel
	EncoderJWTHeaderModel  JWTHeaderModel
	EncoderJWTPayloadModel JWTPayloadModel
	EncodeResult           *JWTEncodeResult

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

		availableHeight := msg.Height - headerHeight - footerHeight - 5

		m.DecoderJWTModel.SetHeight((availableHeight / 2))
		m.DecoderJWTModel.SetWidth((msg.Width / 2))

		m.DecoderSecretModel.SetHeight((availableHeight / 2))
		m.DecoderSecretModel.SetWidth((msg.Width / 2))

		m.DecoderJWTPayloadModel.SetHeight((availableHeight / 2))
		m.DecoderJWTPayloadModel.SetWidth((msg.Width / 2))

		m.DecoderJWTHeaderModel.SetHeight((availableHeight / 2))
		m.DecoderJWTHeaderModel.SetWidth((msg.Width / 2))

		m.EncoderJWTHeaderModel.SetHeight((availableHeight / 2))
		m.EncoderJWTHeaderModel.SetWidth((msg.Width / 2))

		m.EncoderJWTPayloadModel.SetHeight((availableHeight / 2))
		m.EncoderJWTPayloadModel.SetWidth((msg.Width / 2))

		m.EncoderSecretModel.SetHeight((availableHeight / 2))
		m.EncoderSecretModel.SetWidth((msg.Width / 2))

		m.EncoderJWTModel.SetHeight((availableHeight / 2))
		m.EncoderJWTModel.SetWidth((msg.Width / 2))

		m.HelpModel.SetWidth(msg.Width)

		return m, FocusElementCmd(ElementDecoderJWTTextArea)

	case tea.KeyMsg:
		keyStr := msg.String()
		switch keyStr {
		case KeyQuit, KeyQuitAlt:
			return m, tea.Quit
		case KeySwitchView:
			if m.SelectedView == ViewJWTDecoder {
				m.SelectedView = ViewJWTEncoder
				m.FocusedElement = ElementEncoderHeaderTextArea
				return m, FocusElementCmd(m.FocusedElement)
			} else {
				m.SelectedView = ViewJWTDecoder
				m.FocusedElement = ElementDecoderJWTTextArea
				return m, FocusElementCmd(m.FocusedElement)
			}
		}

		switch m.SelectedView {
		case ViewJWTDecoder:
			switch keyStr {
			case KeyFocusToken:
				m.FocusedElement = ElementDecoderJWTTextArea
				return m, FocusElementCmd(m.FocusedElement)
			case KeyFocusSecret:
				m.FocusedElement = ElementDecoderSecretTextArea
				return m, FocusElementCmd(m.FocusedElement)
			case KeyFocusHeader:
				m.FocusedElement = ElementDecoderHeaderTextArea
				return m, FocusElementCmd(m.FocusedElement)
			case KeyFocusPayload:
				m.FocusedElement = ElementDecoderPayloadTextArea
				return m, FocusElementCmd(m.FocusedElement)
			}
		case ViewJWTEncoder:
			switch keyStr {
			case KeyFocusHeader:
				m.FocusedElement = ElementEncoderHeaderTextArea
				return m, FocusElementCmd(m.FocusedElement)
			case KeyFocusPayload:
				m.FocusedElement = ElementEncoderPayloadTextArea
				return m, FocusElementCmd(m.FocusedElement)
			case KeyFocusSecret:
				m.FocusedElement = ElementEncoderSecretTextArea
				return m, FocusElementCmd(m.FocusedElement)
			case KeyFocusToken:
				m.FocusedElement = ElementEncoderJWTTextArea
				return m, FocusElementCmd(m.FocusedElement)
			}
		}
	}

	switch m.SelectedView {
	case ViewJWTDecoder:
		m.DecoderJWTModel, cmd = m.DecoderJWTModel.Update(msg)
		cmds = append(cmds, cmd)

		m.DecoderSecretModel, cmd = m.DecoderSecretModel.Update(msg)
		cmds = append(cmds, cmd)

		m.DecoderJWTHeaderModel, cmd = m.DecoderJWTHeaderModel.Update(msg)
		cmds = append(cmds, cmd)

		m.DecoderJWTPayloadModel, cmd = m.DecoderJWTPayloadModel.Update(msg)
		cmds = append(cmds, cmd)

		m.DecodeResult = nil
		token := m.DecoderJWTModel.GetToken()
		secret := m.DecoderSecretModel.GetSecret()

		if token != "" {
			m.DecodeResult = JWTDecodeToken(token, secret)

			if m.DecodeResult != nil {
				if !m.DecodeResult.IsTokenValid {
					m.DecoderJWTModel.SetError(StatusInvalidToken)
				} else {
					m.DecoderJWTModel.SetError("") // Clear error if valid
				}

				if !m.DecodeResult.IsSignatureValid {
					m.DecoderSecretModel.SetError(StatusSignatureVerificationFailed)
				} else {
					m.DecoderSecretModel.SetError("") // Clear error if valid
				}
			}

			if m.DecodeResult.Token != nil {
				m.DecoderJWTHeaderModel.SetData(m.DecodeResult.JsonMarshaledHeader())
				m.DecoderJWTPayloadModel.SetData(m.DecodeResult.JsonMarshaledClaims())
			} else {
				m.DecoderJWTHeaderModel.SetData("")
				m.DecoderJWTPayloadModel.SetData("")
			}
		} else {
			m.DecoderJWTModel.SetError("")
			m.DecoderSecretModel.SetError("")
		}
	case ViewJWTEncoder:
		m.EncoderJWTHeaderModel, cmd = m.EncoderJWTHeaderModel.Update(msg)
		cmds = append(cmds, cmd)

		m.EncoderJWTPayloadModel, cmd = m.EncoderJWTPayloadModel.Update(msg)
		cmds = append(cmds, cmd)

		m.EncoderSecretModel, cmd = m.EncoderSecretModel.Update(msg)
		cmds = append(cmds, cmd)

		m.EncoderJWTModel, cmd = m.EncoderJWTModel.Update(msg)
		cmds = append(cmds, cmd)

		var headerStr, payloadStr string

		headerStr = m.EncoderJWTHeaderModel.GetData()
		payloadStr = m.EncoderJWTPayloadModel.GetData()

		secretStr := m.EncoderSecretModel.GetSecret()

		var headerError, payloadError string

		var header map[string]interface{}
		var claims jwt.MapClaims

		if headerStr != "" {
			if err := json.Unmarshal([]byte(headerStr), &header); err != nil {
				headerError = "Invalid header JSON: " + err.Error()
			}
		}

		if payloadStr != "" {
			if err := json.Unmarshal([]byte(payloadStr), &claims); err != nil {
				payloadError = "Invalid payload JSON: " + err.Error()
			}
		}

		m.EncoderJWTHeaderModel.SetError(headerError)
		m.EncoderJWTPayloadModel.SetError(payloadError)

		if (headerStr != "" && headerError == "") && (payloadStr != "" && payloadError == "") {
			m.EncodeResult = JWTEncodeToken(header, claims, secretStr)
			m.EncoderJWTModel.SetToken(m.EncodeResult.Token)
		} else {
			m.EncoderJWTModel.SetToken("")
			m.EncodeResult = &JWTEncodeResult{
				Token:        "",
				HeaderError:  headerError,
				PayloadError: payloadError,
				SigningError: "",
			}
		}
	}

	return m, tea.Batch(cmds...)
}

func (m BubbleTeaModel) View() tea.View {
	v := tea.View{
		AltScreen: true,
	}

	var content string

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

		content = lipgloss.JoinHorizontal(lipgloss.Left,
			pane1,
			pane2,
		)
	case ViewJWTEncoder:
		pane1 := lipgloss.JoinVertical(lipgloss.Left,
			m.EncoderJWTHeaderModel.View(),
			m.EncoderJWTPayloadModel.View(),
		)

		pane2 := lipgloss.JoinVertical(lipgloss.Left,
			m.EncoderSecretModel.View(),
			m.EncoderJWTModel.View(),
		)

		content = lipgloss.JoinHorizontal(lipgloss.Left,
			pane1,
			pane2,
		)
	}

	decoderStyle, encoderStyle := styleInactiveScreen, styleInactiveScreen

	switch m.SelectedView {
	case ViewJWTDecoder:
		decoderStyle = styleActiveScreen
	case ViewJWTEncoder:
		encoderStyle = styleActiveScreen
	}

	header := styleHeader.Width(m.WindowSize.Width).
		Render(decoderStyle.Render(TitleDecoder) + styleInactiveScreen.Render(" | ") + encoderStyle.Render(TitleEncoder) + styleInactiveScreen.Render(" (switch: ctrl+\\)"))

	footer := lipgloss.NewStyle().Padding(0, 1, 0, 1).MarginTop(1).Render(m.HelpModel.View(m))

	fullContent := lipgloss.JoinVertical(lipgloss.Left,
		header,
		content,
		footer,
	)

	v.SetContent(fullContent)

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
			key.NewBinding(key.WithKeys(KeySwitchView), key.WithHelp(KeySwitchView, "Switch to Encoder")),
		}
	case ViewJWTEncoder:
		return []key.Binding{
			key.NewBinding(key.WithKeys(KeyQuit, KeyQuitAlt), key.WithHelp(KeyQuit, "Quit")),
			key.NewBinding(key.WithKeys(KeyFocusHeader), key.WithHelp(KeyFocusHeader, "Focus Header")),
			key.NewBinding(key.WithKeys(KeyFocusPayload), key.WithHelp(KeyFocusPayload, "Focus Payload")),
			key.NewBinding(key.WithKeys(KeyFocusSecret), key.WithHelp(KeyFocusSecret, "Focus Secret")),
			key.NewBinding(key.WithKeys(KeyFocusToken), key.WithHelp(KeyFocusToken, "Focus JWT")),
			key.NewBinding(key.WithKeys(KeySwitchView), key.WithHelp(KeySwitchView, "Switch to Decoder")),
		}
	}

	return []key.Binding{}
}

func (m BubbleTeaModel) FullHelp() [][]key.Binding {
	return [][]key.Binding{}
}
