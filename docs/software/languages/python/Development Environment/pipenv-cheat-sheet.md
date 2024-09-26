# `pipenv` Cheat Sheet

## Activate

```bash
pipenv shell
```

## Install from requirements.txt

```bash
pipenv install -r ./requirements.txt
```

## Get Python Interpreter (for VSCode)

```bash
find $(pipenv --venv) -name "python"
```

## Ignore pipfile

```bash
pipenv install --ignore-pipfile
```

## Set lockfile - before deployment

```bash
pipenv lock
```

## Removing Old Environment

When we get the following error:

```bash
❯ pipenv shell
Usage: pipenv shell [OPTIONS] [SHELL_ARGS]...

ERROR:: --system is intended to be used for pre-existing Pipfile installation, not installation of specific packages. Aborting.
```

We can remove the virtual environment to fix:
```bash
❯ rm -r $(pipenv --venv)

❯ pipenv shell
```

## Run with pipenv

```bash
pipenv run *
```
