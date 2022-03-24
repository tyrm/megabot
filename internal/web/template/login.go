package template

// LoginTemplate contains the variables for the "login" template.
type LoginTemplate struct {
	Common

	BotImage string

	FormError    string
	FormEmail    string
	FormPassword string
}
