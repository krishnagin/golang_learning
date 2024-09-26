# golang_learning
Golang Learning Project

Following the bootcamp 

https://one2n.io/go-bootcamp/go-projects


# Makefile

Make file for managing the common scripts being run for building, running and testing the project.

Make file resembles like package.json script section from JS projects.

# Topics

Go code is organized as packages. Each file starts with a package name - that can be used while importing into other part of the program

Group the imports under one import section line by line without any commas or symbols (factored imports appears to be a neat way)

Exports from a package starts with Captial Letter like P in Pi

Functions are basic unit of programming -> inputs have types after the variable name, return type is specified at the right side as well (It feels weird to see this syntax coming from a Java background)

Functions can return any number of values - very interesting feature coming from Java. (Multi return values are wrapped in ())

Named return values - skip

var is the keyword to declare and initialize variables, more shorthand options available.

short assignment syntax  := doesn't require var keyword and need a value for initalization. type can be referred from value 