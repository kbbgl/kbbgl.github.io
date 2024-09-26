# Match Using Iterations

## using the `{x}` will match whatever comes before it `x` amount of times

```bash
echo "12341234444412354444" | grep -P "4{3}" 
# will match 3 times '444'
```

## using the `{x,y}` will match a range from `x` to `y`

```bash
echo "12341234444344412354444" | grep -P "4{1,4}" 
# will match 4, 4444, 444, 4444[
```
