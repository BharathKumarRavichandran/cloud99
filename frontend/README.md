# Frontend for Cloud99

### Prerequirements
* Install [npm](https://docs.npmjs.com/downloading-and-installing-node-js-and-npm)

### Project Installation

1. Clone the repository - `git clone <remote-url>`
2. Go to the project directory - `cd <cloned-repo>`
3. Install dependencies - `npm install`

### Development Build
* To serve files, use the command: `npm run serve:dev`
* To build files: `npm run build:dev`

### Production Build
1. Build files: `npm run build:prod`
2. Serve files: `node prodserver.cloud99.js`

### Linter Instructions
Linter has been set up to run before each serve.
1. To find lint errors - `npm run lint`
2. To handle simple fixes automatically - `npm run lint:fix`

#### Note
* If you face some version incompatability issue while installing/running, check your `node` and `npm` version and ensure it is compatible with the project.