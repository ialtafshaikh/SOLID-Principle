## Single responsibility principle
“Do one thing and do it well” — McIlroy (Unix philosophy)

The single responsibility principle suggests that two separate aspects of a problem need to be handled by a different module. In other word, changes in a module should be originated from only one reason.

In Object-oriented language, if you have more than one responsibilities embedded into a single class, the internal logics become highly coupled, which makes the class less responsive towards changes. Similarly, if you have two separate classes say class A and class B, and if the consumer of class A needs to know about class B, then A and B are considered highly coupled. Single responsibility principle aims to maintain a good level of Coupling that also maintains a good level of Cohesion.

Let’s take an example in Golang. Say we have a module Command in a command-driven system. The Command module decode, validate and finally execute the incoming commands.

```
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
```

In this case, changes on how a Command gets decoded and validated and how a command gets executed will directly affect theCommand module. Hence the module performs multiple responsibilities and highly coupled. As per single responsibility principle the Decode() and Validate() is a separate concern than Execute(), and should be handle in separate modules.

We can introduce the CommandFactory module that parses, validates and initializes a command, where the CommandExecutor module executes the command. Now CommandFactory and CommandExecutor are loosely coupled via Command module. Also notice how we separated the validation of command type and input to the corresponding module.

```
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
```