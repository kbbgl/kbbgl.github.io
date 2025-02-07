import { themes as prismThemes } from "prism-react-renderer";
import type { Config } from "@docusaurus/types";
import type * as Preset from "@docusaurus/preset-classic";

const config: Config = {
  title: "Kobbi's Knowledgebase",
  tagline: "",
  favicon: "img/favicon.svg",
  url: "https://kbbgl.github.io",
  baseUrl: "/",
  organizationName: "kbbgl",
  projectName: "kbbgl.github.io",
  onBrokenLinks: "warn",
  onBrokenMarkdownLinks: "warn",
  i18n: {
    defaultLocale: "en",
    locales: ["en"],
  },

  presets: [
    [
      "classic",
      {
        docs: {
          sidebarPath: "./sidebars.ts",
        },
        blog: {
          showReadingTime: true,
          feedOptions: {
            type: ["rss", "atom"],
            xslt: true,
          },
          onInlineTags: "warn",
          onInlineAuthors: "warn",
          onUntruncatedBlogPosts: "warn",
        },
        sitemap: {
          lastmod: 'datetime',
          changefreq: 'daily',
          priority: 0.5,
          ignorePatterns: ['/tags/**'],
          filename: 'sitemap.xml'
        },
        theme: {
          customCss: "./src/css/custom.css",
        },
      } satisfies Preset.Options,
    ],
  ],

  themeConfig: {
    image: "img/docusaurus-social-card.jpg",
    navbar: {
      title: "Kobbi's Knowledgebase",
      logo: {
        alt: "My Site Logo",
        src: "img/logo.svg",
      },
      items: [
        {
          type: "docSidebar",
          sidebarId: "docsSidebar",
          position: "left",
          label: "Docs",
        },
        {
          to: "/blog",
          sidebarId: "blogSidebar",
          position: "left",
          label: "Blog",
        },
        {
          to: "/hack_blog",
          sidebarId: "hackBlogSidebar",
          position: "left",
          label: "Hack Blog",
        },
      ],
    },
    footer: {
      style: "dark",
      links: [
        {
          title: "Knowledgebase",
          items: [
            {
              label: "Knowledgebase",
              to: "/docs",
            },
          ],
        },
        {
          title: "More",
          items: [
            {
              label: "Blog",
              to: "/blog",
            },
            {
              label: "Hack Blog",
              to: "/hack_blog",
            },
            {
              label: "GitHub",
              href: "https://github.com/kbbgl",
            },
          ],
        },
      ],
      copyright: `Copyright © ${new Date().getFullYear()} Kobbi's Knowledgebase, Built with Docusaurus.`,
    },
    prism: {
      theme: prismThemes.github,
      darkTheme: prismThemes.dracula,
      additionalLanguages: [
        "armasm",
        "bash",
        "batch",
        "c",
        "docker",
        "erlang",
        "go",
        "markup",
        "ini",
        "javascript",
        "json",
        "nginx",
        "powershell",
        "python",
        "yaml",
        "excel-formula",
        "css",
        "protobuf",
      ],
    },
    // https://docusaurus.io/docs/search#connecting-algolia
    algolia: {
      appId: process.env.ALGOLIA_APP_ID,
      apiKey: process.env.ALGOLIA_API_KEY,
      indexName: process.env.ALGOLIA_INDEX_NAME
    }
  } satisfies Preset.ThemeConfig,
};

export default config;
