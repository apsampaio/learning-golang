<h2 align="center">
    Learning GoLang 📚
</h2>

<p align="center">
  <a target="blank"><img src="./.github/gopher.gif" width="100" alt="Gopher" /></a>
</p>

<p align="center">
    <img alt="BUILT WITH" src="https://forthebadge.com/images/badges/built-with-love.svg" />
    <img alt="GO" src="https://forthebadge.com/images/badges/made-with-go.svg" />
</p>

## 🤔 Why use Go?

- Like C, but with garbage collection, memory safety, and special mechanisms for concurrency
- Pointers but no pointer arithmetic
- No header files
- Simple, clean syntax
- Very fast native compilation (about as quick to edit code and restart as a dynamic language)
- Easy-to-distribute executables
- No implicit type coercions
- Simple built-in package system
- Simple tools
- Inferred types on variable declarations
- Slices and maps (feel like arrays and hashmaps in dynamic languages)
- Explicit error handling with error values and multiple return
- Interface-based polymorphism
- Interfaces implicitly implemented (allowing post-hoc interfaces for imported types)
- Goroutines (runtime managed lightweight threads)
- Channels for coordinating goroutines and sharing data between them (based on the theory of [CSP](https://en.wikipedia.org/wiki/Communicating_sequential_processes))

## 🧶 Base types

```go
int8           // 8-bit signed int
int16          // 16-bit signed int
int32          // 32-bit signed int
int64          // 64-bit signed int

uint8          // 8-bit unsigned int
uint16         // 16-bit unsigned int
uint32         // 32-bit unsigned int
uint64         // 64-bit unsigned int

float32       // 32-bit float
float64       // 64-bit float

complex64      // two 32-bit floats
complex128     // two 64-bit floats

int        // 32- or 64-bit signed int (depends upon compilation target)
uint       // 32- or 64-bit unsigned int (depends upon compiliation target)

uintptr    // unsigned int large enough to store an address on compilation target

string     // a string value is an address referencing UTF-8 data elsewhere in memory
bool

byte    // alias for uint8
rune    // alias for int 32 (used for representing Unicode code points)
```

## 🔮 Semi-colon insertion

The semi-colons used in most C-like syntaxes are generally left implicit in Go. Semi-colons are implicit at the end of any line ending with:

- a number, string, or boolean literal
- an identifier
- the reserved words _`break continue fallthrough return`_
- the operators _`++ --`_
- the end delimiters _`} ] )`_

## 🎉 Variable declarations and type inference

The full syntax for declaring a variable:

```
var NAME TYPE;                   // declare without initialization
var NAME TYPE = EXPRESSION;      // declare with initialization
```
