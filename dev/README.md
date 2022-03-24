# Dev tools / environment

## Git hooks
### Automatic commit message formatter and sanitizer

- Implemented as a `commit-msg` git hook which calls a `python` script that implements the following logic for determining whether the commit message is valid:
    - A commit message header consists of 
        ```
            <key>: <header_message>
        ```    
        - `<key>` can be any of the case-insensitive words predetermined by the following table:
            | Key     | 😀 |
            |---------|----|
            | NEW     | 📦 |
            | IMPROVE | 👌 |
            | FIX     | 🐛 |
            | DOC     | 📖 |
            | RELEASE | 🚀 |
            | MERGE   | 🧲 |
        - `<header_message>` will automatically be capitalized and will remove any extra space or the trailing dot.
    - NOTE: currently only the header is automatically formatted, but it can easily be extended to the body as well

#### Example

- Input
    ``` 
        iMpRovE:     refactor the code.

        blah blah...
    ```
- Output
    ``` 
        IMPROVE 👌: Refactor the code

        blah blah...
    ```

