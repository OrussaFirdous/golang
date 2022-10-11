NAMING
---------------------

The names of Go functions, variables, constants, types, statement labels, and packages follow a
begins with a letter or an underscore and may have any number of additional letters, digits, and
underscores.

If an entity is declared within a function, it is local to that function. If declared outside of a
function, however, it is visible in all files of the package to which it belongs.

The case of the first letter of a name determines its visibility across package bound aries. If the
name begins with an upper-case letter, it is exported, which means that it is visible and accessible
outside of its own package and may be referred to by other parts of the program, as with Printf in the
fmt package. 

Package names themselves are always in lowercase.

Generally, the larger the scope of a name, the longer and more meaningful it should be.

Go programmers use ‘camel case’ when forming names by combining words; that is, interior capital letters\
are preferred over interior underscores. 

DECLARATION
--------------

There are four major kinds of declarations: var, const, type, and func.

The name of each package-level entity is visible not only throughout the source file that contains its
declaration, but throughout all the files of the package. By contrast, local declarations are visible only
within the function in which they are declared and perhaps only within a small part of it.

A function declaration has a name, a list of parameters (the variables whose values are provided by the 
function’s callers), an optional list of results, and the function body, which contains the statements that
define what the function does.
