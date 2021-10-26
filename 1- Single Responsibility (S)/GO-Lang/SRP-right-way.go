type Command struct {
    commandType string 
    args []string
}
type CommandFactory struct {
    ...
}
// Create decode and validate the command
func (cf CommandFactory) Create(data []byte) (*Command, error) {
    // decode command
    command, err := cf.Decode(data)
    if err != nil {
        return nil, err
    }
    // validate type
    switch cf.Type { 
       case Foo:
       case Bar:
       default:
          return nil, InvalidCommandType    
    }
    return command, nil
}
type CommandExecutor struct {
}
// Execute executes the command 
func (ce CommandExecutor) Execute(command *Command) ([]byte, error) {
   // validate input and execute 
   switch command.Type {
   case Foo: 
       if len(args) == 0 || len(args[0]) == 0 {
           return nil, InvalidInput
       }
       return ExecuteFoo(command)
   case Bar: // Bar doesn't take any input
       return ExecuteBar(command)
   }
}