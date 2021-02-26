## PHP Example Explanation

Still building from the previous `ShapeInterface` example, you will need to support the new `three-dimensional` shapes of Cuboid and Spheroid, and these shapes will need to also calculate volume.

Let’s consider what would happen if you were to modify the `ShapeInterface` to add another contract:

```
interface ShapeInterface
{
    public function area();

    public function volume();
}
```

Now, any shape you create must implement the volume method, but you know that squares are flat shapes and that they do not have volumes, so this interface would force the Square class to implement a method that it has no use of.

This would violate the interface segregation principle. Instead, you could create another interface called `ThreeDimensionalShapeInterface` that has the volume contract and three-dimensional shapes can implement this interface:

```
interface ShapeInterface
{
    public function area();
}

interface ThreeDimensionalShapeInterface
{
    public function volume();
}

class Cuboid implements ShapeInterface, ThreeDimensionalShapeInterface
{
    public function area()
    {
        // calculate the surface area of the cuboid
    }

    public function volume()
    {
        // calculate the volume of the cuboid
    }
}
```

This is a much better approach, but a pitfall to watch out for is when type-hinting these interfaces. Instead of using a ShapeInterface or a `ThreeDimensionalShapeInterface`, you can create another interface, maybe `ManageShapeInterface`, and implement it on both the flat and three-dimensional shapes.

This way, you can have a single API for managing the shapes:

```
interface ManageShapeInterface
{
    public function calculate();
}

class Square implements ShapeInterface, ManageShapeInterface
{
    public function area()
    {
        // calculate the area of the square
    }

    public function calculate()
    {
        return $this->area();
    }
}

class Cuboid implements ShapeInterface, ThreeDimensionalShapeInterface, ManageShapeInterface
{
    public function area()
    {
        // calculate the surface area of the cuboid
    }

    public function volume()
    {
        // calculate the volume of the cuboid
    }

    public function calculate()
    {
        return $this->area();
    }
}
```

Now in `AreaCalculator` class, you can replace the call to the area method with calculate and also check if the object is an instance of the `ManageShapeInterface` and not the `ShapeInterface`.

That satisfies the `interface segregation principle`.
