“Functions that use pointers or references to base classes must be able to use objects of derived classes without knowing it”
Alternatively, this can be expressed as “Derived classes must be substitutable for their base classes”.

In (maybe) simpler words, if a subclass redefines a function also present in the parent class, a client-user should not be noticing any difference in behaviour, and it is a substitute for the base class.

For example, if you are using a function and your colleague change the base class, you should not notice any difference in the function that you are using.
Among all the SOLID principle, this is the most abstruse to understand and to explain. For this principle, there is no standard “template-like” solution where it must be applied, and it is hard to offer a “standard example” to showcase.

In the most simplistic way, I can put it, this principle can be summarised by saying:
If in a subclass, you redefine a function that is also present in the base class, the two functions ought to have the same behaviour. This, though, does not mean that they must be mandatorily equal, but that the user, should expect that the same type of result, given the same input.
In the example ocp.py, the “operation” method is present in the subclasses and in the base class, and an end-user should expect the same behaviour from the two.
The result of this principle is that we’d write our code in a consistent manner and, the end-user will need to learn how our code works, only one.


Addendum:
(You can skip to the next principle).
A consequence of LSP is that: the new redefined function in the sub-class should be valid and be possibly used wherever the same function in the parent class is used.
This is not, typically the case, indeed usually we, human, think in terms of set theory. Having a class that define a concept and sub-classes that expand the first with an exception or different behaviour.
For example, the sub-class “Platypus”, of the base class “Mammals”, would have the exception that these mammals lay eggs. The LSP, tell us that it would create a function called “give_birth”, this function will have different behaviour for the sub-class Platypus and the sub-class Dog. Therefore, we should have had a more abstract base class than Mammals that accommodate this.
If this sounds very confusing, do not worry, the application of this latter aspect of the LSP is rarely fully implemented, and it rarely leaves the theoretical textbooks.
