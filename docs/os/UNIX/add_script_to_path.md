# Add Script to Path

```bash
nano ~/.zshrc
```

Then add script path:

```bash
# ~/.zshrc
export PATH="~/scripts:$PATH"
```

```bash
ls ~/scripts/
# bastion.sh
```

```bash
source ~/.zshrc
```

```bash
bastion.sh
```
