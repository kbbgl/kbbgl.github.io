---
slug: unit-testing-python
title: Unit Testing in Python
authors: [kgal-pan]
tags: [python, ut, unittesting, unittest, test]
---


## Pytest

#### Run Specific Method in Module

```bash
pytest -vv -s SecurityScorecard_test.py::test_get_domain_services
```

#### Run Specific Method in Class (ignore warning)

```bash
pytest --disable-warnings -s -vv $MODULE_NAME_test.py::$CLASS_NAME::$METHOD_NAME
```

#### Ignore function from Pytest

```python
@pytest.mark.skip(reason="no way of currently testing this")
def test_the_unknown():
    ...
```

#### Raising Errors

```python
with pytest.raises(CustomException) as exc_info:
    raises_custom_exception()

exception_raised = exc_info.value
assert SOME_STRING in str(exception_raised)
```

#### Using `pytest_mock`
```python
from pytest_mock import MockerFixture

def test_auth(mocker: MockerFixture):
    pass
```


#### Freezig Time
```python
from datetime import datetime
from freezegun import freeze_time

@freeze_time("1970-01-01 00:00:00")
def test_first_fetch():
    pass
```

## Coverage
#### Running Coverage Report

From within the `/Packs/PACK_NAME/Integrations/INTEGRATION_NAME` folder:

Using `coverage`:

```bash
coverage run -m pytest $INTEGRATION_NAME_test.py

coverage html

open htmlcov/index.html
```

Using `pytest-cov`:

```bash
pip install pytest-cov

pytest --cov-report html --cov=$(pwd) $INTEGRATION_NAME_test.py

open htmlcov/index.html
```

#### Ignore function from Coverage

```python
def main():  # pragma: no cover
    """
        PARSE AND VALIDATE INTEGRATION PARAMS
```


