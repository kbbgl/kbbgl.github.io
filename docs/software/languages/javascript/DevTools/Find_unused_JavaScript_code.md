# Find Unused JavaScript Code

Chrome’s Coverage panel allows you to quickly locate JavaScript (and CSS) code that has — and has not — been used. To start, open Coverage from the More tools sub-menu in the DevTools menu. Reload the page and the panel will show the percentage of unused code with a bar chart:

Click any JavaScript file and unused code is highlighted with a red bar in the line number gutter.

If you are testing a Single Page App, you can navigate around and use features to update the unused code indicators. However, be wary that coverage information is reset when you reload or navigate to a new page. This is a great tool to understand how much of an external depednencies are you using in your code, if you’re importing a 100kb library and only using 1% of it, you’ll clearly see it here too.
