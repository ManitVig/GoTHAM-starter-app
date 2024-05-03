# GoTHAM Quickstart Guide

This guide walks you through setting up a new Go project using the create-gotham-app tool. This tool helps you quickly scaffold a GoTHAM project boilerplate with essential features and facilitates easy import path management.

## Prerequisites

- Go: Download and install the Go programming language from the official website (https://go.dev/doc/install). Ensure you have the appropriate version for your system.
- Templ: The default GoTHAM setup uses the Templ templating engine. Read the official guide for how to setup Templ https://templ.guide/quick-start/installation
- Air: create-gotham-app also uses Air for hot-reload. Read https://github.com/cosmtrek/air?tab=readme-ov-file#installation to install air cli on your system
- JavaScript Runtime and Package Manager (Choose One):
  - Node.js and npm (for npx): Download and install Node.js (https://nodejs.org/en) which includes npm by default.
  - Bun (Default used by create-gotham-app) (for bunx): Download and install Bun from (https://bun.sh/). Bunx comes bundled with Bun
  - Yarn: Install Yarn, a popular alternative package manager for JavaScript, from their website (https://yarnpkg.com/).
  - pnpm: Install pnpm, another package manager known for its efficient dependency management, following their installation instructions (https://pnpm.io/).

## Using `create-gotham-app`

1. Open your terminal or command prompt.
2. Navigate to where you want to create the project
3. run the create-gotham-app script using:
   - `npx create-gotham-app <project-path> <go-project-identifier>` or
   - `bunx create-gotham-app <project-path> <go-project-identifier>`
5. open the project directory `cd <project-path>`
6. Install go packages `templ generate && go mod tidy`
7. Install javascript packages:
   - If using node/npm or yarn or pnpm
     1. run `npm install` or `yarn install` or `pnpm install`
     2. modify the run command in file `.air.toml`. At line 24 change `pre_cmd = ["templ generate", "bun run build"]` to the appropriate command:
      - `pre_cmd = ["templ generate", "npm run build"]` or
      - `pre_cmd = ["templ generate", "yarn run build"]` or
      - `pre_cmd = ["templ generate", "pnpm run build"]`
   - If using bun
     1. run `bun install`
8. Run the project using command `air`

## Additional Notes
- The create-gotham-app script takes care of cloning the GoTHAM starter app template repository, initializing the Go module (go.mod), and modifying import paths within the codebase to reflect your chosen project identifier.
- Remember to replace <project-path> with the actual path where you want to create the project and <go-project-identifier> with your desired Go project identifier that follows Go package naming conventions (lowercase alphanumeric characters with underscores and hyphens).
   
