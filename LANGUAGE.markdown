# The Language

An object oriented programming language, including functions, and entry points.

## General Overview

The aim of the language is to help with building and deploying projects. But also
help to follow-up on the tasks, providing abilities to deal with factors that
required by the tasks.

The aim is to help tools such as Puppet, Ansible etc to work after deployment
if needed.

Also configuration and connection for tools such as Docker, Kubernetes etc...

The language itself is not a pure programming language for general purpose, but
designed to optimized for it's goals.

### Why another build system?

Most building tools focused on programming language or specific tasks, while
everything that does not suite to such requirements find itself implemented in
different technologies, making it harder to maintain.

### Why new programming language?

Two main types of programming languages exists in the world - compiled vs interpreted.
While both types are good, they contains other issues that exists - they are
general purposed languages most of the times.

The current implementation provides a domain specific language for building applications.
The language helps the builder to focus on the build task rather then on writing
a general purpose language.

### Language goals

**The main goals are:**

- Providing minimal footprint of written code.
- Providing as much as possible tools for both building and deployment of projects.
- Provide agnostic project support, so the project programming should not be important.
- Provide reusable tools to-do build and deploy of a project for writing once,
  and use everywhere ;-).

**Non Goals:**

- Full featured programming language.
- Development for non open source operating systems support (by me).
-

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

Function is a means to create sub execution of specific rule, that *can* return
any value back.

It focused on a small code chunk that does something, rather then a collection
of a big logic.
It can use Object's instances, create them, call other functions, and use data types.

## Objects

Object is a means to hold a structure that should be created from many sub related
elements.



### Methods


### Members


## Namespace


## Import


## Export


## Entry Points


## Arrays


## Hashes


## Numeric


## Struct


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
