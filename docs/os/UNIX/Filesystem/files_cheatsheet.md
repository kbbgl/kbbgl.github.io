# Reading Files

## View File from Bottom

```bash
tac filename.txt
```

## Concatenate Multiple Files

```bash
cat file1.txt file2.txt
```

## Number lines in `cat`

```bash
cat -n list.txt
```

## View File with `less` from end

```bash
less +G list.txt
```

## View File from end and tail

```bash
less +F app.log
```

## Start `less` from first instance of string 'fail'

```bash
less +/fail app.log
```

## Find a file

```bash
locate $FILENAME
```

## `less` with Line Numbers

```bash
less -N $FILE
```

## Find Largest Files in Directory

```bash
sudo du -h --max-depth=1 /var/log/program
```

## Recursive Read First and Last Lines from File

```bash
find . -name "*log*" | while read -r file; do echo $file; head -n 1 $file; tail -n 1 $file; done;
```

## Creating Files

```bash
touch {config,main}.py
```
