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
  * String infix operators
    - ==
    - !=
  * Integer/Boolean infix operators
    - \>=
    - <=
  * Dot notation for hashes
  * ~~Require function for binding an evaluated file to an identifier~~
    - Still needs tests
  * Add import keyword to import specific identifier objects from a file
  * Wrapper for Go's HTTP functions
    - Routing string parser
    - Route function
    - Route middleware
    - Static files
    - Forms
  * JSON support
  * Templates
  * Documentation and examples
