type Command interface {
	Execute() ([]byte, error)
	ValidateInput() bool
}
type CommandExecutor struct {
}
func (c CommandExecutor) Execute(command Command) {
	if command.ValidateInput() {
		 command.Execute()
	}
}
type FooCommand struct {
	args []string // need args
}
func (c FooCommand) ValidateInput() {
   // validate args 
   if len(args) >= 1 && len(args[0]) > 0 {
	  return true
   }
   return false
}
func (c FooCommand) Execute() ([]byte, error) {
   ...
}
type BarCommand struct {
}
func (c BarCommand) ValidateInput() {
   // does nothing 
   return false
}
func (c BarCommand) Execute() ([]byte, error) {
   ...
}
