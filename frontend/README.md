# FRONTEND SERVICE

frontend service of dhammanava search system

## Prerequisites

Make sure you have the following tools installed on your machine:

- Node.js (v16 or higher)
- yarn (if you have npm delete `yarn.lock`)

## Run locally

to run this project locally simply open your terminal and type

1. Install dependencies (depends on your prefer package manager)

```bash
yarn
```

or

```bash
npm install
```

2. Start development server

```bash
yarn run dev
```

or

```bash
npm run dev
```

## Build

This project use docker to build and containerize every service however if you want to build this project you can run

```bash
yarn build
```

or

```bash
npm run build
```

Then you will get a `dist` directory, you can run a server and point the request to this directory to access the server e.g. start nginx server and pass the request to this directory

Alternatively, to quickly view the result of build you can run

```bash
yarn preview
```

or

```bash
npm run preview
```

## Tech stack

**Build Tools** : Vite

**Framework** : React Typescript

**Package manager** : yarn

**UI Framework** : Chakra UI

**Rich Text editor** : Tiptap

## Project structure

###

    .
    ├── deploy/nginx            # nginx service used in deploy process to point to dist directory
    ├── dist                    # Complie files
    ├── public                  # Public directory
    ├── src                     # Source files
        ├──  components         # Components directory
        ├──  constant           # Constant directory
        ├──  functions          # Functions directory
        ├──  hook               # Custom hook directory
        ├──  models             # Data model directory
        ├──  pages              # Pages directory
        ├──  service            # Service directory (call backend services)
        ├──  theme              # Theme directory (styled Chakra UI)
        ├──  App.tsx
        ├──  index.tsx
    ├── package.json            # Package files
    ├── yarn.lock               # Lock file for package
    ├── Dockerfile              # Docker file use to build docker image
    └── README.md
