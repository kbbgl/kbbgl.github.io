# Run `diff` with specific lines to ignore

Make sure to install [`delta`](https://dandavison.github.io/delta/installation.html)

```bash
diff -u <(sed -E '/id|"x"|"y"/d' some.yml) <(sed -E '/id|"x"|"y"/d' some_DEV.yml) | delta --line-numbers
```

To jump between diff sections, use n/N.
