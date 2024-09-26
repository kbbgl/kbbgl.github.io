# Locate DOM-changing Code

It can be difficult to determine which function is responsible for updating a specific HTML DOM element when an event occurs.

To locate a process, right-click any HTML element in the `Elements` panel and select one of the options from the `Break on` sub-menu:

Choose:

- **subtree modifications** to monitor when the element or any child element is changed
- **attribute modifications** to monitor when the element’s attributes, such as class, are changed, or
- **node removal** to monitor when the element is removed from the DOM.

A breakpoint is automatically activated in the `Sources` panel when such an event occurs.
