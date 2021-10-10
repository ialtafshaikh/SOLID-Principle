“A class should have one, and only one, reason to change”
In other words, every component of your code (in general a class, but also a function) should have one and only one responsibility. As a consequence of that, there should be only a reason to change it.
Too often you see a piece of code that takes care of an entire process all at once. I.e., A function that loads data, modifies and, plots them, all before returning its result.
Let’s take a simpler example, where we have a list of number L = [n1, n2, …, nx] and we compute some mathematical functions to this list. For example, compute the mean, median, etc.

A bad approach would be to have a single function doing all the work:

```
import numpy as np

def math_operations(list_):
    # Compute Average
    print(f"the mean is {np.mean(list_)}")
    # Compute Max
    print(f"the max is {np.max(list_)}") 

math_operations(list_ = [1,2,3,4,5])
# the mean is 3.0
# the max is 5

```

The first thing we should do, to make this more SRP compliant, is to split the function math_operations into atomic functions! Thus, when a function’s responsibility cannot be divided into more subparts.
The second step is to make a single function (or class), generically named, “main”. This will call all the other functions one-by-one in a step-to-step process.


```
def get_mean(list_):
    '''Compute Max'''
    print(f"the mean is {np.mean(list_)}") 

def get_max(list_):
    '''Compute Max'''
    print(f"the max is {np.max(list_)}") 

def main(list_): 
    # Compute Average
    get_mean(list_)
    # Compute Max
    get_max(list_)

main([1,2,3,4,5])
# the mean is 3.0
# the max is 5
view raw
```

Now, you would only have one single reason to change each function connected with “main”.
The result of this simple action is that now:
It is easier to localize errors. Any error in execution will point out to a smaller section of your code, accelerating your debug phase.
Any part of the code is reusable in other section of your code.
Moreover and, often overlooked, is that it is easier to create testing for each function of your code. Side note on testing: You should write tests before you actually write the script. But, this is often ignored in favour of creating some nice result to be shown to the stakeholders instead.
This is already a much bigger improvement with respect to the first code example. But, having created a “main” and calling functions with single responsibility is not the full fulfilment of the SR principle.
Indeed, our “main” has many reasons to be changed. The class is actually fragile and hard to maintain.
To solve that, let’s introduce the next principle: