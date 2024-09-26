# Functions

```javascript
.=function() { 
   var x = new Date();
   x.setDate(1);
   x.setMonth(x.getMonth()-1);
   return ['January', 'February', 'March', 'April', 'May', 'June', 'July', 'August', 'September', 'October', 'November', 'December'][x.getMonth()];
}()
```

## With Arguments

For example, `val` here is the result of the `parseJSON` transformation:

```javascript
.=function() { 
   return  val.type + " " + val.riskType;
}()
```
