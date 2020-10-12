# Lexer
These are notes and a working example of a lexer written in Go. This was developed by going through the **Lexing** portion of [Writing An Interpreter In Go](https://interpreterbook.com/).

## Notes
- A **lexer** takes source code as input and output the tokens that represent it. It is also known as **lexical analysis**, **lexing**, or **tokenization**.
  - Lexing is the first step in transforming text written by humans to a computer-readable format that is easy to work with.
    ```
    Source code:
    let x = 5 + 5;

    Lexical output:
    [
        LET,
        IDENTIFIER("x"),
        EQUAL_SIGN,
        INTEGER(5),
        PLUS_SIGN,
        INTEGER(5),
        SEMICOLON
    ]
    ```
- **Tokens** are small, categorizable data structures that are fed to a **parser**, which handles the second transformation step, turning our tokens into an **Abstract Syntax Tree** (AST).
  - A token is a string with an assigned and identifiable meaning. It is structured as a pair consisting of a token **name** and an optional token **value**.
  - The token name is a category of lexical unit. Common token names are
    - **identifier:** names the programmar chooses
    - **keyword:** names already in the programming language
    - **separator:** puncuation characters and paired-delimiters
    - **operator:** symbols that operate on arguments and produce results
    - **literal:** numeric, logical, textual, reference literals
    - **comment:** line, block
