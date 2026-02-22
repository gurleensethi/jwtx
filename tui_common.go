package main

import (
	"image/color"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type View string
type Element string

const (
	ViewJWTEncoder View = "jwt_encoder"
	ViewJWTDecoder View = "jwt_decoder"

	ElementDecoderJWTTextArea     Element = "decoder-jwt-token"
	ElementDecoderSecretTextArea  Element = "decoder-secret-text-area"
	ElementDecoderHeaderTextArea  Element = "decoder-header-text-area"
	ElementDecoderPayloadTextArea Element = "decoder-payload-text-area"
	ElementEncoderHeaderTextArea  Element = "encoder-header-text-area"
	ElementEncoderPayloadTextArea Element = "encoder-payload-text-area"
	ElementEncoderSecretTextArea  Element = "encoder-secret-text-area"
	ElementEncoderJWTTextArea     Element = "encoder-jwt-token"

	KeyQuit         = "ctrl+c"
	KeyQuitAlt      = "ctrl+q"
	KeyFocusToken   = "ctrl+j"
	KeyFocusSecret  = "ctrl+s"
	KeyFocusHeader  = "ctrl+h"
	KeyFocusPayload = "ctrl+p"
	KeySwitchView   = "ctrl+\\"

	StatusValidJWT                    = "Valid JWT"
	StatusInvalidToken                = "Invalid token"
	StatusSignatureVerified           = "Signature Verified"
	StatusSignatureVerificationFailed = "Signature verification failed"

	PlaceholderJWT    = "Enter the JSON Web Token (JWT) here..."
	PlaceholderSecret = "Enter Secret"

	TitleJWTToken       = "JSON WEB TOKEN (ctrl+j)"
	TitleSecret         = "SECRET (ctrl+s)"
	TitleDecodedHeader  = "DECODED HEADER (ctrl+h)"
	TitleDecodedPayload = "DECODED PAYLOAD (ctrl+p)"
	TitleEncoderHeader  = "HEADER (ctrl+h)"
	TitleEncoderPayload = "PAYLOAD (ctrl+p)"
	TitleDecoder        = "JWT Decoder"
	TitleEncoder        = "JWT Encoder"
)

var (
	// List of all elements
	Elements = []Element{
		ElementDecoderJWTTextArea,
		ElementDecoderSecretTextArea,
		ElementDecoderHeaderTextArea,
		ElementDecoderPayloadTextArea,
		ElementEncoderHeaderTextArea,
		ElementEncoderPayloadTextArea,
		ElementEncoderSecretTextArea,
		ElementEncoderJWTTextArea,
	}

	styleTitle = lipgloss.NewStyle().
			MarginBottom(1)

	styleTitleSelected = styleTitle.
				Bold(true).
				Background(color.White).
				Foreground(color.Black)

	styleHeader = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#ffffff")).
			Background(lipgloss.Color("#db3fce")).
			Padding(1, 1).
			MarginBottom(1).
			Align(lipgloss.Center).
			Bold(true)

	styleActiveScreen = lipgloss.NewStyle().
				Underline(true).
				Background(lipgloss.Color("#db3fce"))

	styleInactiveScreen = lipgloss.NewStyle().
				Background(lipgloss.Color("#db3fce"))

	styleBox = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#777777"))

	styleBoxActive = styleBox.Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#ffffff"))

	styleStatus = lipgloss.NewStyle().
			Padding(0, 2, 0, 2)

	styleStatusError = styleStatus.
				Foreground(lipgloss.Color("#ffffff")).
				Background(lipgloss.Color("#ca4a00"))

	styleStatusSuccess = styleStatus.
				Foreground(lipgloss.Color("#ffffff")).
				Background(lipgloss.Color("#008202"))

	styleFooter = lipgloss.NewStyle().
			Padding(0, 2, 0, 2).
			MarginTop(2)

	styleFooterError = styleFooter.
				Foreground(lipgloss.Color("#ffffff")).
				Background(lipgloss.Color("#ca4a00"))

	styleFooterSuccess = styleFooter.
				Foreground(lipgloss.Color("#ffffff")).
				Background(lipgloss.Color("#008202"))
)

type FocusElementMsg struct {
	Element Element
}

func FocusElementCmd(element Element) tea.Cmd {
	return func() tea.Msg {
		return FocusElementMsg{Element: element}
	}
}
