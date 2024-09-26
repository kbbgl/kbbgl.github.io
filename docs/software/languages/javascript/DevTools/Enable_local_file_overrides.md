# Enable Local File Overrides

Chrome allows any HTTP request to use a local file on your device rather than fetch it over the network. This could allow you to:

- edit scripts or styles on a live site without requiring build tools
- develop a site offline which normally requests essential files from a third-party, or domain
- temporarily replace an unnecessary script such as analytics.

Create a directory on your local PC where overrides files will be stored, e.g. localfiles, then open Chrome’s DevTools Sources panel. Open the Overrides tab in the left-hand pane, click + Select folder for overrides, and choose the directory you created. You’ll be prompted to Allow files to be saved locally and the directory will appear:

Now open the Page tab and locate any source file. There are two ways to add it as a local override:

1. right-click the file and choose Save for overrides, or
2. open the file, make an edit, and save it with Ctrl | Cmd + S.

It will also be present in the Overrides tab and the `localfiles` directory. The file can be edited within Chrome or using any code editor — the updated version is used whenever the page is loaded again.
