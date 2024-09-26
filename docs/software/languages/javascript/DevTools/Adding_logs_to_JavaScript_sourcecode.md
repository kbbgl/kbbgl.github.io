# Adding Logs to Source Code

It’s tempting to liberally insert `console.log()` debugging statements throughout your files but logpoints provide a way to obtain the same information without writing any code.

To add a log point, open a script in the `Sources` panel, right-click any line number, and choose `Add log point`….

Enter an expression such as one you would use in a console command, e.g.

```plaintext
"The value of x is", x
```

The message will appear in the DevTools `console` whenever that line is executed. Log points will usually persist between page refreshes.
