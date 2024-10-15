# Ro programming language

A programming language with keywords in Romanian.

## Lexical Conventions

Identifiers in Ro can be any string of Latin letters, Arabic-Indic digits, and 
underscores, not beginning with a digit and not being a reserved word. Identifiers
are used to name variables, table fields, and labels. 

The following keywords are reserved and cannot be used as names:

*daca*, *altfel*, *var*, *adevarat*, *fals*, *fn*, *returneaza*, *si*, *sau*, *incrementeaza*,
*decrementeaza*, *pentru {} in*, *afiseaza*

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
    returneaza adevarat
} altfel {
    returneaza fals
}
```

For statements:

```
pentru user in users {
    afiseaza(user.name)
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

- și: AND
- sau: OR

## Builtins

- concatenează: string concatenation
- lungime: string & array length

## Functions

```
var add = fn (x, y) {
    x + y
}

var result = add(5, 10)
```

## What could come next

- [ ] float numbers
- [ ] closures
- [ ] while loops, switch statements and other control flows
- [ ] classes
- [ ] language server and formatter 
- [ ] some other funky stuff
