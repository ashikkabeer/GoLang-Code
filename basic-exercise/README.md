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

