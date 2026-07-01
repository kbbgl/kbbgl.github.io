// remark config used by the pre-commit link check (.husky/pre-commit).
//
// This is a Docusaurus site, so many links are site-absolute slug URLs
// (e.g. `/docs/managing-aks-azure`, `/blog/some-post`) that resolve through
// Docusaurus routing rather than the filesystem. `remark-validate-links`
// only understands real file paths, so it reports those valid links as
// missing files/headings.
//
// `skipPathPatterns` is matched against the resolved "absolute path + hash".
// Docusaurus routes resolve to an extensionless final segment
// (`.../azure-identity`), whereas real links keep their extension
// (`.../certificate.md`). Skipping extensionless targets silences the false
// positives while still validating genuine relative file links.
export default {
  plugins: [
    "remark-mdx",
    ["remark-validate-links", { skipPathPatterns: ["/[^/.]+(#.*)?$"] }],
  ],
};
