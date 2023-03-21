# Unit testing

## 3 sections of a Unit test

- **Arrainge**
- **Act**
- **Assert**

package adder

    func Test_addNumbers(t *testing.T) {
    expecetd:= 5.0

    got := addNumbers(1,4)        
    if got!=expected {
                t.Error("expected %.1f got %.1f", expected, got)
        }
    }

## Arrainge

- Define all the things needed in order to run the code.

- Usually also define the expected outcome of the test.

    package adder

    func Test_addNumbers(t *testing.T) {
        expecetd:= 5.0

        got := addNumbers(1,4)
        
        if got!=expected {
            t.Error("expected %.1f got %.1f", expected, got)
        }
    }

## Act

- Call the functionality to be tested.

## Assert

- Declare that the actual output of the function is the same as the expected.
