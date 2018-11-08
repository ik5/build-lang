# The Language

An object oriented programming language, including functions, and entry points.

## General Overview

The aim of the language is to help with building and deploying projects. But also
help to follow-up on the tasks, providing abilities to deal with factors that
required by the tasks.

The language itself is not a pure programming language for general purpose, but
designed to optimized for it's goals.

### Why another build language?

Most building tools focused on programming language or specific tasks, while
everything that does not suite such requirements find itself implemented in
different technology, making it hard to maintain these technologies.

### Why new programming language?

Two main types of programming languages exists in the world - compiled vs interpreted.
While both ways are good, they contains other issues that exists - they are
general purposed languages most of the times.

The current implementation provides a domain specific language for building applications.
The language helps the builder to focus on the build task rather then on writing
a general purpose language.

### Language goals

The main goals are:

- Providing minimal footprint of written code.
- Providing as much as possible tools for both building and deployment of projects.
- Provide agnostic project support, so the project programming should not be important.
- Provide reusable tools to-do build and deploy of a project.

Non Goals:

- Full programming language.
- Development for non open source operating systems support.

# The language logic

When breaking the build process down, there are key actions that constantly exists:

  1. Functionality (e.g. how-to copy a file, how-to create user etc...).
  1. Action (e.g. copy a file, create user etc...).
  1. Validation (e.g. does a package exists, what is the OS etc...).
  1. Group of actions and validations.
  1. Build rules.

The language itself accommodate the above steps in a manner that is helpful for
the build process.


# The language itself

## Functions


## Objects


## Namespace


## Import


## Export


## Entry Points


## Arrays


## Hashes


## Numeric


## Modules


## String
There are two types of strings with the language:
  1. `UTF-8` - Default encoding of the string.
  2. Literal string - A string that stored as a chain (array) of bytes with
     a prefix of length and charset identifier.

## Reserved Words

| Word   | Meaning                   |
| ------ | ------------------------- |
| begin  | block of code starts here |
| end    | block of code ends here   |
| const  | create a constant value   |
| var    | create a local variable   |
| func   | create a function         |
| return | return a value            |


## Operators

### Arithmetic

- \+ - add numbers
- \- - subtract numbers
- \* - multiple numbers
- / - divide numbers
- % - reminder on numbers
- ** - Exponentiation

### String

- .. - concat string
- ' - uninterpreted string
- " - interpreted string
- \` - literal char (or string), converted to bytes

## Symbols

### General

- \# - single line comment
- ( - group start
- ) - group end
- , - argument seperator
- ; - end of action
- \< - smaller then
- \> - bigger then
- = - equal
- \<- assign

# Syntax Examples
