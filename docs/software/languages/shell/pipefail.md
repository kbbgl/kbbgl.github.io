# Handling Failures in Pipes

[Source](https://www.howtogeek.com/782514/how-to-use-set-and-pipefail-in-bash-scripts-on-linux/)

## Using `set -e`

When there is no pipe but just a sequence of commands, we can use `set -e`:

```bash
#!/bin/bash

# This will ensure that the script exists
# if any command fails and return code is non-zero
set -e

echo This will happen
ls some_file
echo This will not happen
```

## Failures in Pipes

The return code in a sequence of piped commands is the last one. So if one of the piped commands fails but the last command succeeds, we will erroneously get back a `0` return code and dependent scripts will continue execution.

For example, if we pipe a `false` (which returns a non-zero code) into `true` (which returns a zero code), we will get a `0` return code:

```bash
> false | true

> echo $?
0
```

Bash has an array variable called `PIPESTATUS` which captures all of the return codes from each program in the pipe chain:

```bash
> false | true | false | true

> echo "${PIPESTATUS[0]} ${PIPESTATUS[1]} ${PIPESTATUS[2]} ${PIPESTATUS[3]}"
1 0 1 0
```

`PIPESTATUS` only holds the return codes until the next program runs, and trying to determine which return code goes with which program can get very messy very quickly.

This is where `set -o` (options) and `pipefail` come in.

```bash
set -eo pipefail

echo This will happen first
cat script-99.sh | wc -l
echo This will not be printed
```

This will return a `1` and will not execute the second `echo`.

## Catching Uninitialized Variables

Uninitialized variables can be difficult to spot in a real-world script. If we try to `echo` the value of an uninitialized variable, `echo` simply prints a blank line. It doesn’t raise an error message. The rest of the script will continue to execute.

We can trap this type of error using the `set -u` (unset) option.

```bash
#!/bin/bash 

set -eou pipefail

echo "$notset" 
echo "Another echo command"

# notset: unbound variable
# echo $?
# 1
```

If we want to initialize a value:

```bash
#!/bin/bash 

set -euo pipefail

if [ -z "${New_Var:-}" ]; then 
 echo "New_Var has no value assigned to it." 
fi
```

## Print Command and Parameters using `set -x`
