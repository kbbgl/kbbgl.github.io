# GitHub Actions

GitHub Actions are defined by workflows. Workflows are specified in:

```bash
.github/workflows/mWorkflow.yml
```

Workflow have jobs. Jobs that have steps. Steps can use custom or community actions.

See [this](./workflows/community_actions.yml) to see community actions or [this](./greet/action.yml) to see a custom action.

## Secrets

Secrets can be set in the repository by going to GitHub > Settings > Secrets.

For example, to add debug output, we can set:

```bash
ACTIONS_STEP_DEBUG=true
```
