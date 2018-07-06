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
  * ~~Require function for binding an evaluated file to an identifier~~
    - Still needs tests
  * ~~Fix comment parsing so quotes and special characters work~~
  * ~~Fix string parsing to include the single quote~~
  * Fix string parsing to escape the special characters like \n \t etc.
  * Add import keyword to import specific identifier objects from a file
  * Add a standard library that can be imported into any file.
    - Standard lib files can be imported by itself (aka `import map from 'Array';` or `import 'Array'`) and without extension
    - Any other files have to be a relative path or absolute path to the file to import (aka `import func from './module.svo';` or `import './module.svo'`) and must have the file extension
    - Once the import keyword is tested and working, remove the require function
  * Wrapper for Go's HTTP functions
    - Routing string parser
    - Route function
    - Route middleware
    - Static files
    - Forms
  * JSON support
  * Templates
  * Documentation and examples
  * Dot notation for hashes
