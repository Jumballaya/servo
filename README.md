# Servo

Servo is a simple interpreted scripting language with syntax similar to JavaScript. The language was designed to build simple HTTP APIs with a builtin 'app' object with an API similar to Express.js.

This language was built following the book [Writing an Interpreter in Go by Thorsten Ball](https://interpreterbook.com/).

## Installation

`go install servo`

## Usage

#### REPL

run the command `servo`

#### Scripts

run the command `servo path/to/file.sv`

#### Docker

I have no official image up yet. Build your own image from the Dockerfile in this repository.

## Documentation

None made so far

## TODO List

  * ~~Comments using #~~
  * ~~Single quotes for strings~~
  * ~~String infix operators~~
  * ~~Integer/Boolean infix operators~~
  * ~~Fix comment parsing so quotes and special characters work~~
  * ~~Fix string parsing to include the single quote~~
  * ~~Fix string parsing to escape the special characters like \n \t etc.~~
  * ~~Add import keyword to import specific identifier objects from a file~~
  * ~~Get null working~~
  * ~~Bitwise operators >>, <<, |, &, &^~~
  * ~~Automatic semicolon insertion~~
  * ~~Increment/Decremnt by operators += , -= , \*= , \=~~
  * ~~Power operator ^~~
  * Add simple classes with fields and methods
    - Dot notation `foo.bar`
  * Add a standard library that can be imported into any file.
    - ~~Basic support~~
    - Standard lib files can be imported by itself (aka `import map from 'Array';`) and without extension
    - Any other files have to be a relative path or absolute path to the file to import (aka `import func from './module.svo';`) and must have the file extension
  * Implement bytes
  * Implement floats
  * Rewrite lexer/parser to use runes
  * Multiline Comments /\* \*/
  * String escaping e.g. `\\b` or `"\"hello\""`
  * Hex digits e.g. `0xfff`
  * Wrapper for Go's HTTP functions
    - Routing string parser
    - Route function
    - Route middleware
    - Static files
    - Forms
  * Templates
  * Documentation and examples
  * Dot notation for hashes
  * Add line and column location to tokens (for better debugging)
  * Add for/while loops
  * Add try/catch
  * Add `import './file.svo' as file` syntax to import
  * Add `import func from './example.svo' as function` syntax to import
  * Change import so it builds the AST during the parsing stage rather than evaluation
