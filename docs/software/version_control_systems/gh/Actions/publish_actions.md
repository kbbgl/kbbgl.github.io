# How to Publish Actions to GitHub Actions Marketplace

We use [this boilerplate](https://github.com/actions/javascript-action) and click on 'Use this template', create the repo and clone it locally.

Then we install dependencies using:

```bash
npm install
```

Then change the code in yml/`index.js`, install all missing dependencies and use `ncc` to build the bundle as we saw in [Creating Custom Actions](creating_custom_actions.md#javascript).

To the [action yml](../actions/greet/action.yml), we also need to add the `branding` section:

```yml
outputs:
  # ...
runs:
  # ...
branding:
  # https://docs.github.com/en/actions/creating-actions/metadata-syntax-for-github-actions#brandingcolor
  color: purple
  # https://feathericons.com/
  icon: 'alert-octagon'
```

Then fill out `README.md` that explains how to use the Action.	 

https://docs.github.com/en/actions/creating-actions/metadata-syntax-for-github-actions