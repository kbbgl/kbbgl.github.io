# Environmental Variables

You can check to see if the variable is in the dictionaries returned by `globals()` and `locals()`.

For a local variable:

```python
if locals().get('abc'):
    print(abc)
```

For a global variable:

```bash
if globals().get('abc'):
    print(abc)
```

For an environment variable:

```python
if os.environ.get('abc')=='True':
    #abc is set to True
```
