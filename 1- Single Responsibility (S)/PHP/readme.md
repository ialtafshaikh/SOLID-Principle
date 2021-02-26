## PHP Example Explanation

For example, consider an application that takes a collection of shapes—circles, and squares—and calculates the sum of the area of all the shapes in the collection.

First, create the shape classes and have the constructors set up the required parameters.

For squares, you will need to know the length of a side:

```
class Square
{
public $length;

    public function construct($length)
    {
        $this->length = $length;
    }

}
```

For circles, you will need to know the radius:

```
class Circle
{
public $radius;

    public function construct($radius)
    {
        $this->radius = $radius;
    }

}
```

Next, create the AreaCalculator class and then write up the logic to sum up the areas of all provided shapes. The area of a square is calculated by length squared. The area of a circle is calculated by pi times radius squared.

```
class AreaCalculator
{
protected $shapes;

    public function __construct($shapes = [])
    {
        $this->shapes = $shapes;
    }

    public function sum()
    {
        foreach ($this->shapes as $shape) {
            if (is_a($shape, 'Square')) {
                $area[] = pow($shape->length, 2);
            } elseif (is_a($shape, 'Circle')) {
                $area[] = pi() * pow($shape->radius, 2);
            }
        }

        return array_sum($area);
    }

    public function output()
    {
        return implode('', [
          '',
              'Sum of the areas of provided shapes: ',
              $this->sum(),
          '',
      ]);
    }

}
```

To use the AreaCalculator class, you will need to instantiate the class and pass in an array of shapes and display the output at the bottom of the page.

Here is an example with a collection of three shapes:

- a circle with a radius of 2
- a square with a length of 5
- a second square with a length of 6

```
  $shapes = [
  new Circle(2),
  new Square(5),
  new Square(6),
  ];

$areas = new AreaCalculator($shapes);

echo $areas->output();
```

The problem with the output method is that the AreaCalculator handles the logic to output the data.

Consider a scenario where the output should be converted to another format like JSON.

All of the logic would be handled by the AreaCalculator class. This would violate the single-responsibility principle. The AreaCalculator class should only be concerned with the sum of the areas of provided shapes. It should not care whether the user wants JSON or HTML.

To address this, you can create a separate SumCalculatorOutputter class and use that new class to handle the logic you need to output the data to the user:

```
class SumCalculatorOutputter
{
protected $calculator;

    public function __constructor(AreaCalculator $calculator)
    {
        $this->calculator = $calculator;
    }

    public function JSON()
    {
        $data = [
          'sum' => $this->calculator->sum(),
      ];

        return json_encode($data);
    }

    public function HTML()
    {
        return implode('', [
          '',
              'Sum of the areas of provided shapes: ',
              $this->calculator->sum(),
          '',
      ]);
    }

}
```

The SumCalculatorOutputter class would work like this:

```
$shapes = [
new Circle(2),
new Square(5),
new Square(6),
];

$areas = new AreaCalculator($shapes);
$output = new SumCalculatorOutputter($areas);

echo $output->JSON();
echo $output->HTML();
```

Now, the logic you need to output the data to the user is handled by the `SumCalculatorOutputter` class.

That satisfies the single-responsibility principle.
