module.exports = {
  title: 'Benthos',
  tagline: 'The stream processor for mundane tasks',
  url: 'https://benthos.dev',
  baseUrl: '/',
  favicon: 'img/favicon.ico',
  organizationName: 'Jeffail',
  projectName: 'benthos',
  themeConfig: {
    prism: {
      theme: require('./src/plugins/prism_themes/monokai'),
    },
    navbar: {
      title: 'Benthos',
      logo: {
        alt: 'Benthos Blobfish',
        src: 'img/nav-logo.svg',
      },
      links: [
        {to: 'docs/about', label: 'Docs', position: 'left'},
        {to: 'https://www.jeffail.uk', label: 'Blog', position: 'left'},
        {to: 'https://www.youtube.com/playlist?list=PL9hWaP-BQh2rvNuM29bTLlL0hYk6cqyT5', label: 'Videos', position: 'left'},
        {to: 'https://github.com/Jeffail/benthos/releases/latest', label: 'Download', position: 'right'},
        {
          href: 'https://github.com/Jeffail/benthos',
          label: 'GitHub',
          position: 'right',
        },
      ],
    },
    footer: {
      style: 'dark',
      links: [
        {
          title: 'Docs',
          items: [
            {
              label: 'Getting Started',
              to: 'docs/guides/getting_started',
            },
            {
              label: 'Components',
              to: 'docs/components',
            },
          ],
        },
        {
          title: 'Community',
          items: [
            {
              label: 'Gitter',
              href: 'https://gitter.im/jeffail-benthos/community',
            },
          ],
        },
        {
          title: 'Social',
          items: [
            {
              label: 'Blog',
              to: 'blog',
            },
            {
              label: 'GitHub',
              href: 'https://github.com/Jeffail/benthos',
            },
            {
              label: 'Twitter',
              href: 'https://twitter.com/jeffail',
            },
          ],
        },
      ],
      copyright: `Copyright © ${new Date().getFullYear()} Ashley Jeffs.`,
    },
  },
  presets: [
    [
      '@docusaurus/preset-classic',
      {
        docs: {
          sidebarPath: require.resolve('./sidebars.js'),
          editUrl:
            'https://github.com/Jeffail/benthos/edit/master/website/',
        },
        theme: {
          customCss: require.resolve('./src/css/custom.css'),
        },
      },
    ],
  ],
  scripts: [
    '/js/dark_default.js',
  ],
};
