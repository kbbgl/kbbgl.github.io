# `pipenv` Current Interpreter

```bash
#!/bin/bash

set -e

export PIPENV_VERBOSITY=-1

VENV=$(pipenv --venv)

echo "$VENV/bin/python"
```
