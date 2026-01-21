---
slug: add-timestamp-history
title: Adding Timestamps to Shell History
authors: [kgal-akl]
tags: [tools, bash, shell, zsh, history]
---

## Bash

```bash
vim ~/.bashrc
```

Add the following line to the end of the file to set the timestamp format. This example uses an ISO-like format (Year-Month-Day Hour:Minute:Second):

```bash
export HISTTIMEFORMAT="%F %T "
```

The space after %T ensures a clean separation between the timestamp and the command itself. You can customize the format using `strftime` options.

```bash
source ~/.bashrc
```

## ZSH

Zsh handles history timestamps slightly differently. You don't need `HISTTIMEFORMAT` for saving, but you do need an option to enable the feature, and a flag to display the output. 

```bash
vim ~/.zshrc
```

Add the following line to enable the saving of timestamps with history entries:

```bash
setopt EXTENDED_HISTORY
```

This option saves the history records in the `~/.zsh_history` file with timestamps in Unix epoch format (e.g., #1506591948).

Add the following line to enable appending history immediately and sharing it across sessions (optional, but recommended):

```bash
setopt INC_APPEND_HISTORY

# or
setopt sharehistory
```

```bash
source ~/.zshrc
```

View your history with timestamps using the built-in `fc` command (which `history` is an alias for in ZSH) and the -i flag for ISO format output:

```bash
history -i
```

You can also use other flags like `-f` (US format), `-E` (European format), or `-t` for a custom `strftime` format. 