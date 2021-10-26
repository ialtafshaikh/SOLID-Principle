## Open/Closed principle

“A module should be open for extensions, but closed for modification” — Robert C. Martin


Open Closed is considered one of the most important principles in an object-oriented class-based language. The concept suggests that modules should be written in a way so that we can add new modules or new functionalities in a module without requiring existing modules to be modified.

Let’s assume we have an abstract class S, which provide a common method F() for the derived types A and B. Class S would be considered as closed for an extension if the method F() need to be aware of the existence of the derived classes. This means the addition of a newly derived class say C would need F() to be changed, making 

F() open for modification.

One of the solutions is to make F() work on a defined interface rather than handling the subtypes. Say interface I define the necessary abstract method and need to be implemented by the subtypes A,B and C. The interface I can have many subtypes, so it’s open for extension. And F() is implemented separately to work on interface I, so that it’s closed for modification.

In Golang there is no concept of generalization. Reusability is available as a form of embedding. Although a similar pattern could be seen in practice. Let’s take the example of the CommandExecutor, which is responsible for executing Commands. The Execute() and ValidateInput() methods need to handle each command separately. 

So every time a new command is added Execute() implementation needs to change.

Here we can use a Command interface with Execute() and ValidateInput() method.