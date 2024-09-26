# Conditional Breakpoints

Clicking a line number of a file opened in the `Sources` panel adds a breakpoint. It halts a script at that point during execution so you can step through code to inspect variables, the call stack, etc.

Breakpoints are not always practical, e.g. if an error occurs during the last iteration of a loop which runs 1,000 times. A conditional breakpoint however, can be set so it will only trigger when a certain condition is met, e.g. i > 999. To set one, right-click a line number, choose `Add conditional breakpoint`…, and enter a condition expression.
