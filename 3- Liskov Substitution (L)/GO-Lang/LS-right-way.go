type Command interface {
	Execute() ([]byte, error)
}
type CommandWithInput interface {
	Command
	ValidateInput() bool
}