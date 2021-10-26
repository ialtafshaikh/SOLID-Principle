## Interface segregation principle

“Many client specific interfaces are better than one general purpose interface” — Robert C. Martin

In an object-oriented class-based language, it states that if a class provides methods to multiple clients, then rather than having a generic interface loaded with 
all methods, provide a separate interface for each client and implement all of them in the class.

Say, client C1 uses F1 method, C2 uses F2 method. Interface I provides F1and F2. class A implements interface I. The problems with the generalised interface is that:
Changes in client C1‘s methods can cause changes in C2’s method

A new class B implements interface I but B only get used by the client C2. Which promotes unimplemented methods in B.
The interface segregation principle suggests segregating the interface I into IC1 and IC2, so thatIC1 is responsible for the client C1 and IC2 is responsible for the client C2 .

In Golang interfaces are satisfied implicitly, rather than explicitly, which makes it easier to extend a class behaviour by implementing multiple interface based on needs. It also encourages to the design of small and reusable interfaces.

```
type I1 interface { // consumed by C1
    M1()
    M2()
    M3()       
}
type I2 interface { // consumed by C2 and C3
    M3()       
    M4()
}
```

These may lead to several small interfaces and some clients need to use a type that implements a set of interfaces among all. In Golang aggregates of the interface is particularly useful to define an aggregate interface with a set of interfaces. But breaking down interfaces can be tricky.

Robert C. Martin in his Design Principles and Design Patterns paper mentioned:
As with all principles, care must be taken not to overdo it. The specter of a class with hundreds of different interfaces, some segregated by client and other segregated by version, would be frightening indeed.

The rule to follow is that each interface must be defined in a way so that it provides the exact and full set of functionalities needed by at least one of the client. This also means that there is no need to break an interface if consumed by only one client.

```
type I1 interface { // consumed by C1
    M1()
    M2()
    M3()       // also defined in I2
}
type I2 interface { // consumed by C2 and C3
    M3()       // here M3 not in a separate interface as no client
               // use an interface with only M3()
    M4()
}
type I3 interface { // consumed by C4
    M5()       // similarly M5() only used along with I1 and I2
               // thus not needed to have it in a separate interface
    I1
    I2
}
```

In our previous example, we separated the Command interface into two interfaces:

```
type Command interface {
     Execute() ([]byte, error)
}
type CommandWithInput interface {
     Command
     ValidateInput() bool
}
```

Although we have only one client CommandExecutor that consumes it. Therefore it might not be the best idea to break it into two. Alternatively, we could have added a method NeedInput() that returns either true and false. This way we also make the Contract complete.

```
type Command interface {
     Execute() ([]byte, error)
     HasInput() bool
     ValidateInput() bool
}
```

and change CommandExecutor as

```
func (c CommandExecutor) Execute(command Command) {
     if !command.HasInput() || command.ValidateInput() {
          command.Execute()
     }
}

```