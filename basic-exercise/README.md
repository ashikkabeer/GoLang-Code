# Go Basics
collection of exercises is a showcase of my journey as I dive into the fundamentals of the Go programming language (Golang).

## Points
### var vs :=
- | Feature                                           | Global Scope                                   | Function Scope                                 |
  |---------------------------------------------------|------------------------------------------------|------------------------------------------------|
  | Declaration using `var`                            | ✔️ Can be used inside and outside of functions | ✔️ Can only be used inside functions            |
  | Separation of declaration and assignment          | ✔️ Variable declaration and value assignment can be done separately | ❌ Variable declaration and value assignment must be done in the same line |

### Formatting Verbs
- **%v::** Prints the value in the default format.
- **%#v:** Prints the value in Go-syntax format, which can be used to reproduce the value.
- **%T:**  Prints the type of the value.
- **%%:**  Prints the % sign itself.
- For more formatting verbs and details, you can refer to the [W3Schools Go Formatting Verbs](https://www.w3schools.com/go/go_formatting_verbs.php) page. 

### Types of Basic DataTypes
| Data Type      | Description                                | Example                                       |
|----------------|--------------------------------------------|-----------------------------------------------|
| `int`          | Default integer type (platform-dependent)  | `var age int = 25`                            |
| `int8`         | 8-bit signed integer                       | `var smallNumber int8 = 127`                  |
| `int16`        | 16-bit signed integer                      | `var mediumNumber int16 = 32767`              |
| `int32`        | 32-bit signed integer                      | `var standardNumber int32 = 2147483647`       |
| `int64`        | 64-bit signed integer                      | `var bigNumber int64 = 9223372036854775807`  |
| `uint`, `uint8`, `uint16`, `uint32`, `uint64` | Unsigned integers    | `var count uint = 42`                        |
| `uint8`        | 8-bit unsigned integer                     | `var smallPositiveNumber uint8 = 255`         |
| `uint16`       | 16-bit unsigned integer                    | `var mediumPositiveNumber uint16 = 65535`    |
| `uint32`       | 32-bit unsigned integer                    | `var standardPositiveNumber uint32 = 4294967295` |
| `uint64`       | 64-bit unsigned integer                    | `var bigPositiveNumber uint64 = 18446744073709551615` |
| `bool`         | Boolean type                               | `var isTrue bool = true`                     |
| `float32`      | 32-bit floating-point number               | `var pi float32 = 3.14`                      |
| `float64`      | 64-bit floating-point number               | `var piPrecise float64 = 3.14159265359`      |
| `string`       | String type                                | `var message string = "Hello, Go!"`          |

