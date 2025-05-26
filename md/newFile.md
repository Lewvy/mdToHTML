# Writing an Interpreter in Go – Checklist

> Goal: Build a working interpreter for the Monkey programming language
> Duration: ~3 weeks (adjust as needed)

## Week 1 – Lexer & Parser Basics

### Chapter 1: Introduction
- [ ] Set up Go project structure
- [ ] Understand interpreter architecture overview

### Chapter 2: Lexing
- [ ] Implement `token` package (token types + constants)
- [ ] Implement `lexer` to scan input and return tokens
- [ ] Write tests for tokenization (e.g., `let five = 5;`)
- [ ] Run lexer against basic input and inspect token stream

### Chapter 3: Parsing - Part 1
- [ ] Define `ast` node structs (e.g., Program, Statement, Expression)
- [ ] Create parser that builds AST from tokens
- [ ] Parse `let` statements into AST
- [ ] Write tests for parser output

### Chapter 4: Parsing - Part 2
- [ ] Handle return statements and expressions
- [ ] Implement prefix/infix parsing (Pratt parsing)
- [ ] Parse integers, identifiers, and arithmetic
- [ ] Confirm AST output with test input

## Week 2 – Evaluator Core

### Chapter 5: Evaluator - Part 1
- [ ] Implement `object` system (Integer, Boolean, Null)
- [ ] Evaluate literals and basic expressions
- [ ] Add support for prefix and infix operations
- [ ] Write evaluation tests

### Chapter 6: Evaluator - Part 2
- [ ] Add support for `if-else` conditionals
- [ ] Implement return values
- [ ] Handle error objects in evaluation
- [ ] Write more complex evaluation tests

### Chapter 7: Environment
- [ ] Implement `Environment` (variable storage)
- [ ] Support `let` bindings and identifier lookups
- [ ] Write scope and environment tests

## Week 3 – Functions, Closures, REPL

### Chapter 8: Functions
- [ ] Implement function object + evaluation
- [ ] Parse and evaluate function definitions
- [ ] Implement function calls with arguments

### Chapter 9: Closures
- [ ] Add support for closures (functions returning functions)
- [ ] Test closures capturing environments

### Chapter 10: REPL
- [ ] Build interactive REPL shell
- [ ] Evaluate user input line-by-line
- [ ] Test interpreter interactively

## Optional Extras (from *Writing a Compiler in Go*)
- [ ] Add support for strings, arrays, hash maps
- [ ] Extend with built-in functions (e.g., `len`, `first`)
- [ ] Add indexing and slicing
- [ ] Explore compiling to bytecode later
