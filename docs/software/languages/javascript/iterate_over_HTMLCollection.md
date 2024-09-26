# Iterate over HTMLCollection

```javascript

var table = document.getElementByTagName("table")

var rows = table[0].rows

for (row of rows){
    console.log(row.children[1].innerText + "," + row.children[2].innerText + "," + row.children[3].innerText)
}
```
