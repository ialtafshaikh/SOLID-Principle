type Command struct {
	commandType string 
	args []string
 }
 func (c Command) Decode(data []byte) {
	// decodes and initialise
 }
 func (c Command) ValidateCommandType() bool {
	// validates command type
 }
 func (c Command) ValidateArgs() bool {
	// validate provided args as if input
 }
 func (c Command) Execute() {
	// Executes seperate types of commands 
 }