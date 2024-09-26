# How to Create Custom GitHub Actions

We can create our own actions that utilize JavaScript or Docker containers.

## JavaScript

```bash
mkdir .github/actions/hello

touch .github/actions/hello/action.yml
```

An example of a Javascript custom action can be found in [here](./greet/action.yml).

Some very useful actions for JavaScript can be found in https://github.com/actions/toolkit. For example, in order to get inputs and set outputs, we would need to install the `core` and `github` packages:

```bash
npm install @actions/core
npm install @actions/github
```

And then use them in [`index.js`](./greet/index.js)

### Dependencies

Since the job will run in a VM, NPM packages will not be available. We can push the `node_modules` folder but a better solution is to use [`ncc`](https://github.com/vercel/ncc) and compile all dependencies into one bundle.

```bash
npm install -D @zeit/ncc

npx ncc build .github/actions/greet/index.js -o .github/actions/greet/dist
```

And change the action to use the bundle:

```yml
runs:
  using: 'node12'
  main: 'dist/index.js' # reference the bundle
```

## Docker

To create an action that will run in a Docker container, we need to create:

- An [action YAML](./greet-docker/action.yml) that will hold the inputs, outputs and Docker run specification:

 ```yml
 runs:
 using: 'docker'
 # We can use a docker image and tag from DockerHub or a Dockerfile
 # image: '$IMAGE_NAME:$TAG_NAME'
 image: 'Dockerfile'
 ```

- A [`Dockerfile`](./greet-docker/entrypoint.sh)
- An [`entrypoint.sh`](./greet-docker/entrypoint.sh)
