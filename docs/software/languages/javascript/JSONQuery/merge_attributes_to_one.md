# Merge Attributes into Object

```javascript
!Set key=attributes.manager value=m

!Set key=attributes.displayName value=u

!Print value=${attributes} extend-context=`UserAndManager=.={"user": val.displayName, "manager": val.manager}`
```
