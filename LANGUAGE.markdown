# The Language

An object oriented programming language, including functions, and entry points.

## General Overview

The aim of the language is to help with building and deploying projects. But also
help to follow up on the tasks, providing abilities to deal with many factors that
are required by the tasks.

The language itself is not a pure programming language, for general purpose, but
designed to be optimized for it's goals.

### Why another build language?
The reason for that is, that most building tools, are very focused on programming
language, or specific type of tasks, while everything that does not suite such
requirement find itself implemented in different technology, making it hard to
maintain many technologies making the building and also deploying an avoided domain
for many developers.

### Why new programming language?
There are two main programming languages - compiled vs interpreted.
While both ways are good, they usually contain many other issues with them in order
to provide the best solution for what they build for, making it harder to use them
inside domain specific.

This language is a domain specific language, and as such, it helps to focus on the
tasks it was set to, while not providing more then needed.


### Language goals
The main goals are:

- Providing minimal footprint of written code.
- Providing as much as possible tools for both building and deployment of projects.
- Provide agnostic project support, so the project programming should not be important.
- Provide reusable and extendible tools to do build and deploy of a project.

Non Goals:

- Full programming language.
- Cross platform only open source Unix based OS will be supported by me, but
patches are more then welcome :)

##

## Functions


## Objects


## Namespace


## Import


## Export


## Entry Points


##


## Numeric


## String
The language strings are UTF8 by default, however, there is a literal string, that
stores the written chars as bytes, and can be any encoding that is needed.


## Reserved Words

| Word   | Meaning                   |
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
- ' - uninterpeted string
- " - inerpeted string
- \` - literal char (or string), converted to bytes

## Symbols

### General

- \# - single line comment
- ( - group start
- ) - group end
- , - argument seperator
- ; - end of action
- < - smaller then
- \> - bigger then
- = - equal
- <- assign
