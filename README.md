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

  1. Comments: # and //
  2. Single quotes for strings
  3. Dot notation for hashes
  4. Require function for binding an evaluated file to an identifier
  5. Wrappr for Go's HTTP functions
    * Routing string parser
    * Route function
    * Route middleware
    * Static files
    * Forms
  6. JSON support
  7. Templates
  8. Documentation and examples
