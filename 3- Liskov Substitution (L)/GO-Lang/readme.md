## Liskov substitution principle

“Derived methods should expect no more and provide no less” — Robert C. Martin

In an object-oriented class-based language, the concept of the Liskov substitution principle is that a user of a base class should be able to function properly of all derived classes.

This means if client C uses a class A. And B is a class derived from class A. Then functionalities within client C that depend on the methods of class A should work as it is with the same method of class B. Class B should provide no special method for client C, neither it should leave any method unimplemented. But in practice, it often happens that we come to a situation where the client needs to handle the base class and subclass separately.

Liskov substitution principle suggests that the client and the derived classes should interact via a Contract that defines the client’s intention.
The functionalities within client C, that depend on the methods of class A should be substituted via an abstract base class T with A and B being concrete subtypes. Class T becomes the Contract, and client C consumes the Contract methods.

As we discussed earlier, in Golang there is no class-based inheritance. Instead, Golang provides a more powerful approach towards polymorphism via Interfaces and Struct Embedding. Unlike class-based language, Go polymorphism involves creating many different data types that satisfy a common interface.

In the above example client C need to consume an interface T (the Contract) so that multiple concrete types A and B can be passed. Good thing is that one not need to be aware of all the contracts at the time of defining a type. As in Go, interfaces are satisfied implicitly, rather than explicitly.

> Well designed interfaces are more likely to be small interfaces; the prevailing idiom is an interface contains only a single method. It follows logically that small interfaces lead to simple implementations, because it is hard to do otherwise. Which leads to packages comprised of simple implementations connected by common behaviour.

Designing simple interfaces has been the core fundaments of Golang ecosystem. Such interfaces example includes
> io.Reader
> error
> context.Context

These interface designed in a way so that the implementations are substitutable without any special handling as they fulfil the same contract.

In our earlier example, we provided theCommand interface. Which is fairly simple, but is it good enough?
BarCommand doesn’t have any input. For the same reason, ValidateInput() always return False. Now the client CommandExecutor will fail as it expects ValidateInput() to work. Here BarCommand provides less than what is expected.

Alternatively, we can separate the interface into Command and CommandWithInput as