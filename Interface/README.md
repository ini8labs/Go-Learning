# Interface

## What is an interface

- Interface is a custom type that is used to specify a set of one or more method signatures and the interface is abstract, so you are not allowed to create an instance of the interface.

- But you are allowed to create a variable of an interface type and this variable can be assigned with a concrete type value that has the methods the interface requires.

- An interface type is defined as a set of method signatures. A value of interface type can hold any value that implements those methods.

- Interface types express generalizations about the behaviors of other types.

- It facilitate more flexible and adaptable functions because they are not tied to the details of one particular implementation.

- Interface is a type used to specify set of methods.

- Multiple types implement same interface

## Why Interface

- To acheive Polymorphism(Runtime), call is resolved at runtime.

- Reduce duplication of code and reduces size of code

- Reduce dependency between diffrent parts of code

- where argument of dynamic types are needed

- Futureability: Now if we add other object sphere in future, we just need to define area and volume function for sphere without any changes in old code, then pass sphere in slice argument in calculate().

## How to create an interface

- Interface can be created using a ```type``` keyword, followed by name of interface and ```interface``` keyword.

```text
type interface_name interface {
    Method signatures
}
```

- Example of interface

```text
package pack

type Shape interface {
    Area() float32
    Volume() float32
}
```

An interface "Shape" containing methods Area() and Volume() is created.

## Implementing an Interface

- There is no explicit keyword to implement an interface in Go.

- To implement an interface simply implement all the methods declared in the interface.

```text
package main

import (
    "fmt"
)

type Shape interface {
    Area() float32
    Volume() float32
}
type Cube struct {
    Side float32
}

type Cylinder struct {
    Height float32
    Radius float32
}

func (a Cube) Area() float32 {
    return 6 * a.Side * a.Side
}

func (a Cube) Volume() float32 {
    return a.Side * a.Side * a.Side
}

func (c Cuboid) Area() float32 {
    return 2 * (c.Length*c.Breadth + c.Breadth*c.Height + c.Height*c.Length)
}

func (c Cuboid) Volume() float32 {
    return c.Length * c.Breadth * c.Height
}

func Calculate(o []Shape) {
    for _, v := range o {
        area := v.Area()
        volume := v.Volume()
        fmt.Println(v)
        fmt.Println("Area of object ", area)
        fmt.Println("Volume of object ", volume)
    }

func main() {
    Cube := pack.Cube{Side: 10}

    Cuboid := pack.Cuboid{
        Length:  3,
        Breadth: 5,
        Height:  10,
    }

    shapes := []pack.Shape{Cube, Cuboid}
    pack.Calculate(shapes)
}
```

- In above example we have created ```Shape``` interface and struct type ```Cube``` and ```Cuboid```.

- Then we defined methods ```Area()``` and ```Volume()``` which belong to structs ```Cube``` and ```Cuboid```.

- Therefore, ```Cube``` and ```Cuboid``` implemented those methods.

- Since these methods are defined by ```Shape``` interface, ```Cube``` and ```Cuboid``` struct type implemented ```Shape``` interface.

- This shows how interfaces are implemented implicitly in Go.

- When a type implements an interface, a variable of that type can also be represented as type of an interface. This indicates **polymorphism**.

- As in example varialble of two types ```Cube``` and ```Cuboid``` are implementing same interface ```Shape```.

## Empty Interface

- When an interface has zero methods, it is called empty interface.

- ```type Shape interface{}``` create an empty interface named Shape.

- Since it has zero methods, all types implement this interface implicitly.

- Empty interface are used by code that handles values of unknown type.

### Example of empty interface

```text
package main

import "fmt"

func main() {
    var i interface{}

  a := "Welcome to Programiz"
  b := 20
  c := false

displayValue(a)
displayValue(a, b)
displayValue(a, b, c)


func display(i interface{}) {
    fmt.Printf(i)
}
```

- Here ```display()``` takes any number of arguments of type interface().
