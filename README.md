# ws_substring - MySQL UDF for substring.

### about

This is MySQL User Defined Function written by cgo.  
Substring function.  

|arg|explain|
|---|-------|
|arg1|`input` input string|
|arg2|`start` start position (0 is first)|
|arg3|`length` substring length|

### why substring?

MySQL has already substring function.  
But, when I tried to substring a huge text via JDBC or ADO.NET,  
`Incorrect string value:` error was happened.  
I don't understarnd why this error is happened, so I try to create this udf.  

### how to install

    $ ./build.sh

(notice)  

* require root privilege

### example

(simple1)  

    MariaDB [(none)]> select ws_substring('aaabbbccc', 3, 3);
    +---------------------------------+
    | ws_substring('aaabbbccc', 3, 3) |
    +---------------------------------+
    | bbb                             |
    +---------------------------------+

(simple2)  

    MariaDB [(none)]> select ws_substring('𠮷野家で𠮷野がご飯をたべる', 3, 3);
    +-----------------------------------------------------------+
    | ws_substring('?野家で?野がご飯をたべる', 3, 3)            |
    +-----------------------------------------------------------+
    | で𠮷野                                                    |
    +-----------------------------------------------------------+


