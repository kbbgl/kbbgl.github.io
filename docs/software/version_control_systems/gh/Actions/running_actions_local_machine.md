# Running GitHub Actions on Local Machine

Install [`act`](https://github.com/nektos/act).

To run a particluar workflow on `push` event:

```bash
act push --workflows .github/workflows/flow_control.yml
```

Reuse the container, add environmental variables and set a secret:

```bash
act push \
--workflows .github/workflows/actions.yml \
--reuse \
-s SECRET=ej \
--env-file .env \
```

Sample workflows can be found [here](../Actions)
