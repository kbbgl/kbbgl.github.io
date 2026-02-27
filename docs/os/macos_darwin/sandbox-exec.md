---
slug: os-macos-darwin-sandbox-exec
title: "`sandbox-exec` (macOS): command-line sandboxing"
authors: [kbbgl]
tags: [os, macos, darwin, sandbox_exec]
---

# `sandbox-exec` (macOS): command-line sandboxing

`sandbox-exec` is a built-in macOS command-line utility that runs a program inside a sandbox defined by a **sandbox profile** (a small Scheme/LISP-like policy file). The basic idea is to **deny or allow specific operations** (network, file reads/writes, process exec, etc.) so a command can only access what you explicitly permit.

Source: [sandbox-exec: macOS's Little-Known Command-Line Sandboxing Tool | Igor's Techno Club (2025-04-17)](https://igorstechnoclub.com/sandbox-exec/)

## Quick start

Run a command under a profile file:

```bash
sandbox-exec -f profile.sb command_to_run
```

## Sandbox profile basics

Profiles are built from parenthesized rules. Typical structure:

- `(version 1)`
- Default policy: `(deny default)` or `(allow default)`
- Specific allow/deny rules

Rules can target paths using:

- **literal** path: `(literal "/path/to/file")`
- **regex**: `(regex "^/System")`
- **subpath** (directory subtree): `(subpath "/Library")`

## Two common strategies

### Deny-by-default (more secure)

Start from “nothing is allowed”, then open up only what’s required:

```lisp
(version 1)
(deny default)
(allow file-read-data (regex "^/usr/lib"))
(allow process-exec (literal "/usr/bin/python3"))
```

This is ideal for untrusted code, but you’ll likely need iterations to make the program functional.

### Allow-by-default (more convenient)

Start from “everything is allowed”, then block specific risky operations:

```lisp
(version 1)
(allow default)
(deny network*)
(deny file-write* (regex "^/Users"))
```

This is easier to adopt, but it’s also easier to miss a capability you *meant* to deny.

## Practical examples

### Sandboxed terminal session (no network + limited personal file reads)

Create `terminal-sandbox.sb`:

```lisp
(version 1)
(allow default)
(deny network*)
(deny file-read-data (regex "^/Users/[^/]+/(Documents|Pictures|Desktop)"))
```

Then start a sandboxed shell:

```bash
sandbox-exec -f terminal-sandbox.sb zsh
```

### Use pre-built system profiles

macOS ships profiles under `/System/Library/Sandbox/Profiles`. You can use one directly:

```bash
sandbox-exec -f /System/Library/Sandbox/Profiles/weatherd.sb command
```

These can be a good starting point to copy/adapt.

### Import and extend an existing profile

```lisp
(version 1)
(import "/System/Library/Sandbox/Profiles/bsd.sb")
(deny network*)
```

### Inline profile (good for quick “no network”)

```bash
alias sandbox-no-network='sandbox-exec -p "(version 1)(allow default)(deny network*)"'
```

Example:

```bash
sandbox-no-network curl -v https://google.com
```

Note: the source article reports this style didn’t reliably prevent network access for at least one GUI app invocation (Firefox).

## Debugging: “why did my app fail in the sandbox?”

### Console.app

Open Console (Applications → Utilities → Console), then search for “sandbox” and your app name. Look for messages containing “deny” to see what was blocked.

### Live sandbox violation logs

Stream sandbox logs:

```bash
log stream --style compact --predicate 'sender=="Sandbox"'
```

Filter for an app (example from the source uses “python”):

```bash
log stream --style compact --predicate 'sender=="Sandbox" and eventMessage contains "python"'
```

## Useful use cases

- **Run untrusted scripts/tools**: deny network and restrict reads outside a narrow allowlist.
- **Privacy guardrails**: explicitly deny reads of personal directories (Documents/Pictures/Desktop) for a subprocess.
- **Safer automation**: wrap a one-off data transform so it can’t exfiltrate or overwrite files outside a working directory.
- **Developer testing**: quickly test how an app behaves when permissions are constrained (before doing deeper App Sandbox work).

## Limitations / caveats

- Apple discourages relying on `sandbox-exec` long-term in favor of App Sandbox (per the source article’s “deprecation status” note).
- Non-trivial apps can require lots of trial-and-error to permit the minimum needed capabilities.
- Major macOS updates may change behavior.
- No GUI for authoring profiles; you need to iterate using logs.

## Sources

- [sandbox-exec: macOS's Little-Known Command-Line Sandboxing Tool | Igor's Techno Club (2025-04-17)](https://igorstechnoclub.com/sandbox-exec/)
