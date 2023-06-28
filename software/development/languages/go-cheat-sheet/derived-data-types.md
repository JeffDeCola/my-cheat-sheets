
```go
// ARRAY (Data Structure)
    // DECLARE TYPE
    var a [2]float32                                // var name [number]type
    // ASSIGN VALUE
    a[0] = 1.1                                      // name[number] = value
    a[1] = 2.0
    // DECLARE & ASSIGN (INITIALIZE)
    var b = [2]float32{1.1, 2.0}                    // Verbose - var name = [number]type{value, value, ...}
    c := [2]float32{1.1, 2.0}                       // Array Shorthand Assignment
    // PRINT
    fmt.Println(a, b, c)                            // [1.1 2] [1.1 2] [1.1 2]
// SLICE (Data Structure, Reference Type) (make)
    // DECLARE TYPE - NO SIZE
    var a []float64                                 // var name []type
    // ASSIGN VALUE - ADD LENGTH TO SLICE
    a = append(a, 5.7)                              // name = append(name, value, value, ...)
    a = append([]float64{1.1}, a...)                // Trick to prepend to slice
    // DECLARE TYPE - WITH SIZE (make)
    b := make([]string, 1, 25)                      // name := make([]type, length, capacity)
    // ASSIGN VALUE
    b[0] = "hello"                                  // name[index] = value
    // DECLARE & ASSIGN (INITIALIZE)
    var c = []float32{1.1, 2.0}                     // Verbose - var name = []type{value, value, ...}
    d := []float32{3.4, 4.5}                        // Array Shortcut Assignment
    // PRINT
    fmt.Println(a, b, c, d)                         // [1.1 5.7] [hello] [1.1 2] [3.4 4.5]
// MAP (Data Structure, Reference Type) (make)
    // DECLARE TYPES - THIS IS A NIL MAP - DON'T DO THIS
    var a map[string]int                            // var name map[type]type
    // DECLARE TYPES (make)
    var b = make(map[string]int)                    // var name make(map[type]type)
    c := make(map[string]int)                       // name := make(map[type]type)
    // ASSIGN KEY:VALUE
    b["Jill"] = 23                                  // name[key] = value
    b["Bob"] = 34
    b["Mark"] = 28
    c["Jill"], c["Bob"], c["Mark"] = 23, 34, 28
    // DECLARE & ASSIGN KEY:VALUE (INITIALIZE)
    var d = map[string]int{                         // Verbose - var name = map[type]type {key:value, ...}
        "Jill": 23,
        "Bob":  34,
        "Mark": 28,
    }
    e := map[string]int{                            // Array Shortcut Assignment
        "Jill": 23,
        "Bob":  34,
        "Mark": 28,
    }
    // PRINT
    fmt.Println(a, b, c, d, e)                      // map[Jill:23 Bob:34 Mark:28] (For all of the maps)
    // DELETE A KEY
    delete(e,"Jill")                                // Delete key "Jill"
// STRUCT (Data Structure)
    // CREATE STRUCT TYPE
    type Rect struct {                              // Define a struct
        w, h int
        }
    // DECLARE TYPE
    var r1 Rect                                     // Create a struct
    r2 := new(Rect)                                 // Returns Pointer
    // ASSIGN VALUE TO FIELDS
    r1.w = 2                                        // name.field = value
    r1.h = 4
    r2.w, r2.h = 3, 5
    // DECLARE & ASSIGN (INITIALIZE)
    var r3 Rect = Rect{2, 4}                        // Verbose (Don't use)
    var r4 = Rect{2, 4}                             // Type Inference - var name = structName{value, ....}
    r5 := Rect{w: 2, h: 4}                          // Shorthand Assignment
    r6 := Rect{2, 4}                                // Shorthand Assignment
    // PRINT
    fmt.Println(r1, *r2, r3, r4, r5, r6)            // {2 4} {3 5} {2 4} {2 4} {2 4} {2 4}
// POINTER (new)
    // DECLARE A POINTER TYPE AND ASSIGN
    a := new(int)                                   // Create int pointer type
    *a = 9                                          // "Contents of a is 9"
    fmt.Println("Content of pointer *a is", *a)     // 9
    // ASSIGN A POINTER TYPE AND ASSIGN (OVERKILL)
    var c *int                                      // YOU DON'T NEED THIS
    b := 33                                         // This is an int
    c = &b                                          // "Address of b" is in pointer a
    fmt.Println("Content of pointer *a is", *a)     // 9
    // ASSIGN A POINTER TO A TYPE INT (INFERRED) - DO THIS
    b := 33                                         // If we have a var int 5
    c := &b                                         // assign c the "address of" b
    fmt.Println("Contents of pointer *c is", *c)    // 33
    fmt.Println("Address of &b is", &b)             // 0x02
    fmt.Println("Contents of pointer c is", c)      // 0x02
// FUNCTION AS A TYPE
    // ASSIGN ANONYMOUS FUNCTION (func LITERAL) TO A VARIABLE
    a, b := 3, 9
    add := func() int {                             // Anonymous func as a type (no name)
        return a + b                                // returns int
    }
    fmt.Println(add())                              // 12
    a = 9
    fmt.Println(add())                              // 18
    // CLOSURE - RETURN A FUNCTION TO A FUNCTION
    func addThis(a, b int) func() int {
        return func() int {
            return a + b
        }
    }
    func main() {
        a, b := 3, 9
        add := addThis(a, b)                        // Think of the func add like a variable
        fmt.Println(add())                          // 12
        a = 9
        fmt.Println(add())                          // 12 <- NOTE THIS
    }
// INTERFACE (SEE OWN CHEAT SHEET)
// CHANNEL (make) (SEE OWN CHEAT SHEET)
```
