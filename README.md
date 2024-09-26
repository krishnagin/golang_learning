# golang_learning
Golang Learning Project

Following the bootcamp 

https://one2n.io/go-bootcamp/go-projects


# Makefile

Make file for managing the common scripts being run for building, running and testing the project.

Make file resembles like package.json script section from JS projects.

# Topics

## Basic syntax and data types

Go code is organized as packages. Each file starts with a package name - that can be used while importing into other part of the program

Group the imports under one import section line by line without any commas or symbols (factored imports appears to be a neat way)

Exports from a package starts with Captial Letter like P in Pi

Functions are basic unit of programming -> inputs have types after the variable name, return type is specified at the right side as well (It feels weird to see this syntax coming from a Java background)

Functions can return any number of values - very interesting feature coming from Java. (Multi return values are wrapped in ())

Named return values - skip

var is the keyword to declare and initialize variables, more shorthand options available.

short assignment syntax  := doesn't require var keyword and need a value for initalization. type can be referred from value 

built-in data types are bool, string, int (with bit sizes), byte, rune (is it dune, don't know what it is), float (With bit sizes) and complex (with bit sizes) - If no value initialized while declaring, default values will be assigned based on type

type converison has explicit function calls - unlike type casting operations in Java

constants are declared using const keyword with = assignment

for is while statement in Go. for and if loops don't need braces

for {} is an infinite loop

if statement can have assignment operations like in for

switch statement doesn't need break under each case

## Defer (More to read on this - Unclear at this moment)
defer is a blocking primitive - blocks function execution until the surrounding function returns (what is the use case? mostly for async communication?)

## Pointer
pointer - memory of the value, I think it's useful to manipulate the state of the program values

&p - generates the address value of the variable p

*p - dereferncing - accessing the value through the pointer

pointer p prints something like this 0xc00000a1a0 - memory location in the program execution context

## Struct

struct keyword helps to create custom data types - combination of built-in types to define characteristics of program data

## Arrays

Arrays fixed size n of values of type T

slices are dynamic size 
    - a subset of the underlying array 
    - Literally subset of array 
    - Not a copy, so any changes to the slice affects actual array
    - make provides way to create slices
    - iterate over slice using range keyword (get 2 values 1- index and 2- copy of the value at index)
    - omitting values in multi return type can be done using _ hyphen

## Maps

Typical key value pair data structure
map[key]value

map[string]string

map["key"] = "value"

## Functions

functions are first-class entities, can be passed around like other data types

function closures - skip for now

