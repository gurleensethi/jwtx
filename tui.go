package main

import (
	"image/color"

	"charm.land/bubbles/v2/textarea"
	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type View string
type Element string

var (
	ViewJWTEncoder View = "jwt_encoder"
	ViewJWTDecoder View = "jwt_decoder"

	ElementEncoderJWTTextArea    Element = "el_encoder_jwt_text_area"
	ElementEncoderSecretTextArea Element = "el_encoder_secret_text_area"

	styleTitle = lipgloss.NewStyle().
			MarginBottom(1)

	styleTitleSelected = styleTitle.
				Bold(true).
				Background(color.White).
				Foreground(color.Black)

	styleHeader = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#ffffff")).
			Background(lipgloss.Color("#db3fce")).
			Padding(1, 2).
			MarginBottom(1).
			Align(lipgloss.Center).
			Bold(true)

	styleActiveScreen = lipgloss.NewStyle().
				Bold(true).
				Underline(true).
				Background(lipgloss.Color("#db3fce"))

	styleInactiveScreen = lipgloss.NewStyle().
				Background(lipgloss.Color("#db3fce"))

	styleBox = lipgloss.NewStyle().Border(lipgloss.NormalBorder())
)

func NewBubbleTeamModel() BubbleTeaModel {
	decoderTokenTextArea := textarea.New()
	decoderTokenTextArea.Placeholder = "Enter the JSON Web Token (JWT) here..."
	decoderTokenTextArea.Prompt = ""
	decoderTokenTextArea.ShowLineNumbers = false

	decoderSecretTextArea := textarea.New()
	decoderSecretTextArea.Placeholder = "Enter Secret"
	decoderSecretTextArea.Prompt = ""
	decoderSecretTextArea.ShowLineNumbers = false

	decoderHeaderViewport := viewport.New()
	decoderPayloadViewport := viewport.New()

	return BubbleTeaModel{
		SelectedView:         ViewJWTDecoder,
		SelectedElement:        ElementEncoderJWTTextArea,
		DecoderTokenTextArea:   decoderTokenTextArea,
		DecoderSecretTextArea:  decoderSecretTextArea,
		DecoderHeaderViewport:  decoderHeaderViewport,
		DecoderPayloadViewport: decoderPayloadViewport,
	}
}

type BubbleTeaModel struct {
	WindowSize tea.WindowSizeMsg

	SelectedView  View
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

		// Account for header height (header takes 1 line + margin)
		headerHeight := 2
		heightDeduction := 6
		widthDeduction := 2

		styleTitle = styleTitle.Width((msg.Width / 2) - widthDeduction)
		styleTitleSelected = styleTitleSelected.Width((msg.Width / 2) - widthDeduction)

		// Adjust available height by subtracting header height
		availableHeight := msg.Height - headerHeight

		m.DecoderTokenTextArea.SetHeight((2 * availableHeight / 3) - heightDeduction)
		m.DecoderTokenTextArea.SetWidth((msg.Width / 2) - widthDeduction)
		m.DecoderTokenTextArea.Focus()

		m.DecoderSecretTextArea.SetHeight((availableHeight / 3) - heightDeduction)
		m.DecoderSecretTextArea.SetWidth((msg.Width / 2) - widthDeduction)

		m.DecoderPayloadViewport.SetHeight((2 * availableHeight / 3) - heightDeduction)
		m.DecoderPayloadViewport.SetWidth((msg.Width / 2) - widthDeduction)

		m.DecoderHeaderViewport.SetHeight((availableHeight / 3) - heightDeduction)
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

	switch m.SelectedView {
	case ViewJWTDecoder:
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

		fullContent := lipgloss.JoinVertical(lipgloss.Center,
			m.renderHeader(),
			content,
		)

		v.SetContent(fullContent)
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

func (m BubbleTeaModel) renderHeader() string {
	decoderStyle, encoderStyle := styleInactiveScreen, styleInactiveScreen

	switch m.SelectedView {
	case ViewJWTDecoder:
		decoderStyle = styleActiveScreen
	case ViewJWTEncoder:
		encoderStyle = styleActiveScreen
	}

	return styleHeader.Width(m.WindowSize.Width).
		Render(decoderStyle.Render("JWT Decoder") + encoderStyle.Render(" | JWT Encoder"))
}
