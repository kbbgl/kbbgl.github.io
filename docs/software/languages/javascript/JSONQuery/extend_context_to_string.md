# Extend Context to String

When running the following command, the output `VendorName` will be output as a list/array:

```javascript
!SearchIncidentsV2 query="type:TP and partnervendorid:=100000" extend-context=`VendorName=.=val.Contents.data[0].name` ignore-outputs=true
```

However, if we want it to be a string, we can choose the first element (`VendorName=.[0]=val.Contents.data[0].name`) of the list/array in the middle of the context extension argument:

```javascript
!SearchIncidentsV2 query="type:TP and partnervendorid:=100000" extend-context=`VendorName=.[0]=val.Contents.data[0].name` ignore-outputs=true
```
