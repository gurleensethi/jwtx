package main

import (
	"image/color"

	"charm.land/bubbles/v2/textarea"
	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type Screen string
type Element string

var (
	ScreenJWTEncoder Screen = "jwt_encoder"
	ScreenJWTDecoder Screen = "jwt_decoder"

	ElementEncoderJWTTextArea    Element = "el_encoder_jwt_text_area"
	ElementEncoderSecretTextArea Element = "el_encoder_secret_text_area"

	styleTitle = lipgloss.NewStyle().
			MarginBottom(1)

	styleTitleSelected = styleTitle.
				Bold(true).
				Background(color.White).
				Foreground(color.Black)

	styleBox = lipgloss.NewStyle().Border(lipgloss.NormalBorder())
)

func NewBubbleTeamModel() BubbleTeaModel {
	decoderTokenTextArea := textarea.New()
	decoderTokenTextArea.Placeholder = "Enter the JSON Web Token (JWT) here..."
	decoderTokenTextArea.Prompt = ""

	decoderSecretTextArea := textarea.New()
	decoderSecretTextArea.Placeholder = "Enter Secret"
	decoderSecretTextArea.Prompt = ""

	decoderHeaderViewport := viewport.New()
	decoderPayloadViewport := viewport.New()

	return BubbleTeaModel{
		SelectedScreen:         ScreenJWTDecoder,
		SelectedElement:        ElementEncoderJWTTextArea,
		DecoderTokenTextArea:   decoderTokenTextArea,
		DecoderSecretTextArea:  decoderSecretTextArea,
		DecoderHeaderViewport:  decoderHeaderViewport,
		DecoderPayloadViewport: decoderPayloadViewport,
	}
}

type BubbleTeaModel struct {
	WindowSize tea.WindowSizeMsg

	SelectedScreen  Screen
	SelectedElement Element

	// Decoding
	DecoderTokenTextArea   textarea.Model
	DecoderSecretTextArea  textarea.Model
	DecoderHeaderViewport  viewport.Model
	DecoderPayloadViewport viewport.Model
	DecodeResult           JWTDecodeResult
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

		heightDeduction := 6
		widthDeduction := 2

		styleTitle = styleTitle.Width((msg.Width / 2) - widthDeduction)
		styleTitleSelected = styleTitleSelected.Width((msg.Width / 2) - widthDeduction)

		m.DecoderTokenTextArea.SetHeight((2 * msg.Height / 3) - heightDeduction)
		m.DecoderTokenTextArea.SetWidth((msg.Width / 2) - widthDeduction)
		m.DecoderTokenTextArea.Focus()

		m.DecoderSecretTextArea.SetHeight((msg.Height / 3) - heightDeduction)
		m.DecoderSecretTextArea.SetWidth((msg.Width / 2) - widthDeduction)

		m.DecoderPayloadViewport.SetHeight((2 * msg.Height / 3) - heightDeduction)
		m.DecoderPayloadViewport.SetWidth((msg.Width / 2) - widthDeduction)

		m.DecoderHeaderViewport.SetHeight((msg.Height / 3) - heightDeduction)
		m.DecoderHeaderViewport.SetWidth((msg.Width / 2) - widthDeduction)

	case tea.KeyMsg:
		switch msg.Key().String() {
		// Quit Program
		case "ctrl+c", "ctrl+q":
			return m, tea.Quit
		case "ctrl+1":
			m.SelectedElement = ElementEncoderJWTTextArea
			m.DecoderSecretTextArea.Blur()
			m.DecoderTokenTextArea.Focus()
		case "ctrl+2":
			m.SelectedElement = ElementEncoderSecretTextArea
			m.DecoderTokenTextArea.Blur()
			m.DecoderSecretTextArea.Focus()
		}
	}

	m.DecoderTokenTextArea, cmd = m.DecoderTokenTextArea.Update(msg)
	cmds = append(cmds, cmd)

	m.DecoderSecretTextArea, cmd = m.DecoderSecretTextArea.Update(msg)
	cmds = append(cmds, cmd)

	m.DecodeResult = JWTDecodeToken(m.DecoderTokenTextArea.Value(), m.DecoderSecretTextArea.Value())

	if m.DecodeResult.Token != nil {
		m.DecoderHeaderViewport.SetContent(m.DecodeResult.JsonMarshedHeader())
		m.DecoderPayloadViewport.SetContent(m.DecodeResult.JsonMarshledClaims())
	} else {
		m.DecoderHeaderViewport.SetContent("")
		m.DecoderPayloadViewport.SetContent("")
	}

	return m, nil
}

func (m BubbleTeaModel) View() tea.View {
	v := tea.View{
		AltScreen: true,
	}

	switch m.SelectedScreen {
	case ScreenJWTDecoder:
		pane1 := lipgloss.JoinVertical(lipgloss.Left,
			m.renderJsonWebTokenBox(),
			m.renderSecretBox(),
		)

		pane2 := lipgloss.JoinVertical(lipgloss.Left,
			m.renderPayloadBox(),
			m.renderHeaderBox(),
		)

		content := lipgloss.JoinHorizontal(lipgloss.Left,
			pane1,
			pane2,
		)

		v.SetContent(content)
	}

	return v
}

func (m BubbleTeaModel) renderJsonWebTokenBox() string {
	title := styleTitle
	if m.SelectedElement == ElementEncoderJWTTextArea {
		title = styleTitleSelected
	}

	return styleBox.Render(
		lipgloss.JoinVertical(lipgloss.Left,
			title.Render("JSON WEB TOKEN [ctrl+1]"),
			m.DecoderTokenTextArea.View(),
		),
	)
}

func (m BubbleTeaModel) renderSecretBox() string {
	title := styleTitle
	if m.SelectedElement == ElementEncoderSecretTextArea {
		title = styleTitleSelected
	}

	return styleBox.Render(
		lipgloss.JoinVertical(lipgloss.Left,
			title.Render("SECRET [ctrl+2]"),
			m.DecoderSecretTextArea.View(),
		),
	)
}

func (m BubbleTeaModel) renderHeaderBox() string {
	return styleBox.Render(
		lipgloss.JoinVertical(lipgloss.Left,
			styleTitle.Render("DECODED HEADER"),
			m.DecoderHeaderViewport.View(),
		),
	)
}

func (m BubbleTeaModel) renderPayloadBox() string {
	return styleBox.Render(
		lipgloss.JoinVertical(lipgloss.Left,
			styleTitle.Render("DECODED PAYLOAD"),
			m.DecoderPayloadViewport.View(),
		),
	)
}
