# Introduction

The following project, is aim to re-think how to build projects.
Up until now, most projects uses GNU Make, cmake etc...

The Project goals to create a programming language for projects that will allow to do

- Define entry points
- Share information between points
- Create Pre and Post processing
- Reusable of code
- Easy to read and maintain
- Easy to expand and to be changed based on requirements
- Reusable instructions and deployment
- Support multi-threading processing
- Helping developers to develop and deploy

__**Important**__
Current status of this project is ***still on design mode***, so nothing here is
stable, and everything might change. I check and test stuff to see what works
and what does not.

# FAQ

Q. **Why a new language?**
A. I'm creating a language that is dedicated for generating and  deploying projects,
  but that helps to developers to do so, without working hard and guess how stuff will work.

Q. **Why not to use Python/Ruby/Lua/Language in order to do so?**
A. They all have additional requirements that should be added.
The aim of this project is to be stand-alone. Only code that the developer writes
are the actual dependencies.

Q. **We can extend the language, so there are dependencies**
A. Yes and no.

An expansion of the current build tool, are the same as reusing code, and they
are not part of the stdlib or extlib that arrives with the build system language.

# License

The Application is released under MIT license

