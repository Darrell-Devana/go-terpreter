# go-terpreter

A tree-walking JavaScript interpreter written in Go, using only the standard library.

This is a personal learning project: the goal is to deepen Go skills beyond CRUD-style backend work, and to demystify how interpreters work under the hood.

## Status

Pre-v1, in active development.

## Scope (v1)

- Lexer, parser, AST, tree-walking evaluator
- Variable declarations (`let`, `const`)
- Integer arithmetic: `+`, `-`, `*`, `/`
- String concatenation
- Built-in `print(...)` function
- Control flow: `if` / `else`, `while` loops
- Comparison operators: `<`, `>`, `===`, `!==`
- Source comments (`//` and `/* */`)
- CLI with `--version` and `--help`
- Both REPL and file-execution modes

## Non-goals (v1)

Arrays, objects, `console.log`, pretty errors with line/column numbers, JS type coercion, `==` loose equality, hoisting / `var`, automatic semicolon insertion, prototypes / `this` / classes, async / promises.

## Project layout

```
cmd/go-terpreter/   binary entrypoint
token/              token types and constants
lexer/              source string -> token stream
ast/                AST node types and interfaces
parser/             token stream -> AST
object/             runtime value types
evaluator/          AST -> evaluated result
repl/               REPL loop
```

## Build and run

```
go build ./...
go run ./cmd/go-terpreter            # REPL
go run ./cmd/go-terpreter script.js  # file mode
```

## Testing

```
go test ./...
```

## Approach

- Go standard library only — no third-party packages.
- Reference: Thorsten Ball's *Writing An Interpreter In Go*, used as a skim reference rather than followed line-by-line.
- Tree-walking interpreter (not a bytecode VM).
- Idiomatic Go: interfaces, error returns, table-driven tests, struct embedding.
