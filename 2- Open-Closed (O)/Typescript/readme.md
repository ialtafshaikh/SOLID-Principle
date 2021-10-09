 Software entities should be open for extension but closed for modification.

The risk of changing an existing class is that you will introduce  an inadvertent change in behaviour. The solution is create another class that overrides the behaviour of  the original class. By following the OCP, a component is more likely to contain maintainable and re-usable code.

Example – right way
The CreditCard class describes a method to calculate the monthlyDiscount(). The monthlyDiscount() depends on the type of Card, which can be : Silver or Gold. To change the monthly discount calc, you should create another class which overrides the monthlyDiscount() Method.

The solution is to create two new classes: one for each type of card.
