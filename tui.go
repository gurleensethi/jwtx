package main

import (
	"image/color"

	"charm.land/bubbles/v2/textarea"
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
			Padding(1).
			MarginBottom(1)

	styleTitleSelected = styleTitle.
				Bold(true).
				Background(color.White).
				Foreground(color.Black)

	styleBox = lipgloss.NewStyle().Border(lipgloss.NormalBorder())
)

func NewBubbleTeamModel() BubbleTeaModel {
	encoderTokenTextArea := textarea.New()
	encoderTokenTextArea.Placeholder = "Enter the JSON Web Token (JWT) here..."
	encoderTokenTextArea.Prompt = ""

	encoderSecretTextArea := textarea.New()
	encoderSecretTextArea.Placeholder = "Enter Secret"
	encoderSecretTextArea.Prompt = ""

	return BubbleTeaModel{
		SelectedScreen:        ScreenJWTEncoder,
		SelectedElement:       ElementEncoderJWTTextArea,
		EncoderTokenTextArea:  encoderTokenTextArea,
		EncoderSecretTextArea: encoderSecretTextArea,
	}
}

type BubbleTeaModel struct {
	WindowSize tea.WindowSizeMsg

	SelectedScreen  Screen
	SelectedElement Element

	// Encoder
	EncoderTokenTextArea  textarea.Model
	EncoderSecretTextArea textarea.Model
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
		widthDeduction := 4

		styleTitle = styleTitle.Width((msg.Width / 2) - widthDeduction)
		styleTitleSelected = styleTitleSelected.Width((msg.Width / 2) - widthDeduction)

		m.EncoderTokenTextArea.SetHeight((msg.Height / 2) - heightDeduction)
		m.EncoderTokenTextArea.SetWidth((msg.Width / 2) - widthDeduction)
		m.EncoderTokenTextArea.Focus()

		m.EncoderSecretTextArea.SetHeight((msg.Height / 2) - heightDeduction)
		m.EncoderSecretTextArea.SetWidth((msg.Width / 2) - widthDeduction)

	case tea.KeyMsg:
		switch msg.Key().String() {
		// Quit Program
		case "ctrl+c", "ctrl+q":
			return m, tea.Quit
		case "ctrl+1":
			m.SelectedElement = ElementEncoderJWTTextArea
			m.EncoderSecretTextArea.Blur()
			m.EncoderTokenTextArea.Focus()
		case "ctrl+2":
			m.SelectedElement = ElementEncoderSecretTextArea
			m.EncoderTokenTextArea.Blur()
			m.EncoderSecretTextArea.Focus()
		}
	}

	m.EncoderTokenTextArea, cmd = m.EncoderTokenTextArea.Update(msg)
	cmds = append(cmds, cmd)

	m.EncoderSecretTextArea, cmd = m.EncoderSecretTextArea.Update(msg)
	cmds = append(cmds, cmd)

	return m, nil
}

func (m BubbleTeaModel) View() tea.View {
	v := tea.View{
		AltScreen: true,
	}

	switch m.SelectedScreen {
	case ScreenJWTEncoder:
		pane1 := lipgloss.JoinVertical(lipgloss.Left,
			m.renderJsonWebTokenBox(),
			m.renderSecretBox(),
		)

		pane2 := lipgloss.JoinVertical(lipgloss.Left, "")

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
			m.EncoderTokenTextArea.View(),
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
			m.EncoderSecretTextArea.View(),
		),
	)
}
