---
slug: cp-pip-requirements-from-container
title: Copy requirements.txt from Container
authors: [kbbgl]
tags: [docker, pip, python, dependencies, cp, pipenv, containers]
---

Python code can have imports that do not exist in the environment (`pipenv`) but do exist in a Docker container. We can copy the `requirements.txt` file from within the Docker container and then use `pip` to install them.

```bash
docker run -it --rm $USER/$IMAGE_NAME:1.0.0.XXXXX sh

docker cp $CONTAINER_ID:/requirements.txt ./

pipenv install -r requirement.txt

pipenv clean
```

If using Visual Studio Code, reload window to refresh dependencies in Python interpreter.

Full script to pull requirements from image:

```bash
#!/bin/bash

image=$(grep "dockerimage" *.yml | cut -d":" -f2,3 | tr -d '[[:space:]]')

docker run -it --rm $image cat requirements.txt > requirements.txt

pipenv clean

pipenv install
```
