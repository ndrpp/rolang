# Ro programming language

A programming language with keywords in Romanian.

## Lexical Conventions

Identifiers in Ro can be any string of Latin letters, Arabic-Indic digits, and 
underscores, not beginning with a digit and not being a reserved word. Identifiers
are used to name variables, table fields, and labels. 

The following keywords are reserved and cannot be used as names:

*daca*, *altfel*, *var*, *adev*, *fals*, *fn*, *ret*, *si*, *sau*, *increment*,
*decrement*

## Variables

Variable assignment:

```
var a = 5
```

Constants might follow in a later version.

## Data types

int, bool, string, array, object, nil

## Statements

If statements:

```
daca 5 < 10 {
    ret adev
} altfel {
    ret fals
}
```

For statements:

```
pentru user in users {
    print(user.name)
}
```

## Arithmetic Operations

- +: addition
- -: subtraction
- *: multiplication
- /: float division
- //: floor division
- %: modulo
- ^: exponentiation

## Bitwise Operators

Coming in a future version.

- &: bitwise AND
- |: bitwise OR
- ~: bitwise exclusive OR
- \>>: right shift
- <<: left shift

## Relational Operators

- ==: equality
- !=: inequality
- <: less than
- \>: greater than
- <=: less or equal
- \>=: greater or equal

## Logical Operators

- si: AND
- sau: OR

## Builtins

- concat: string concatenation
- lungime: string & array length

## Functions

```
var add = fn (x, y) {
    x + y
}

var result = add(5, 10)
```

## What could come next

1. closures
2. while loops, switch statements and other control flows
3. classes
3. language server and formatter 
4. some other funky stuff
