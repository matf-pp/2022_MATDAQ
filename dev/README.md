# Dev tools / environment

## Branching strategy

- We will be using the [GitHub flow](https://docs.github.com/en/get-started/quickstart/github-flow)
- The main idea is to have 2 types of branches
  - main branch
  - feature/fix branches

### Main branch

- This branch hosts only the code that is fully runnable
- Whenever you want to share a new feature or fix with your colleagues, it has to be merged into main by creating a PR from your feature/fix branch

### Feature/fix branches

- These are the branches where you will develop your features and fixes
- You base them on main and create commits on them
  - Note that you are free to create as many local branches (that are based on your feature branch) as you want, just remember that you will be creating the PR from the feature branch
- Only when you're **fully finished** with the feature you create a PR
  - One exception to the rule is when you're stuck on a problem you don't know how to solve 
    - Instead of merging into `main` (because we merge into `main` **only complete** features/fixes) for you colleagues to see your work, you can create a PR which will enable easier commenting and asynchronous communication

## Pull Requests (PRs)

- You can create PRs from the web interface or using the [GitHub CLI tool](https://cli.github.com/) (I recommend the tool of course)
  - Before using the `gh` tool you need to set it up with your credentials
- TODO: checkout what's happening with write permissions to the repo

## Workflow example

- Typical workflow example when you are creating a new feature/fix:
  ```bash
$ git checkout main
$ git pull  
$ git checkout -b branch_name
# now you do your development (git add, git commit...) 
# only when you're sure you've finished the feature/fix you create a PR
$ gh pr create
TODO: add the rest when you fix writing permissions with Ivan
# your code should be reviewed by a colleague before getting merged into main
# hint: you can assign reviewers
  ```

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

