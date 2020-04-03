/**
 * Copyright (c) 2017-present, Facebook, Inc.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

// See https://docusaurus.io/docs/site-config for all the possible
// site configuration options.

// List of projects/orgs using your project for the users page.
const users = [
  {
    caption: 'User1',
    // You will need to prepend the image path with your baseUrl
    // if it is not '/', like: '/test-site/img/image.jpg'.
    // image: '/gomithrilapp/img/undraw_open_source.svg',
    // infoLink: 'https://www.facebook.com',
    // pinned: true,
  },
];

const siteConfig = {
  title: 'gomithrilapp', // Title for your website.
  tagline: 'Sample Notepad Application in Go and Mithril',
  url: 'https://josephspurrier.github.io', // Your website URL
  baseUrl: '/gomithrilapp/', // Base URL for your project */
  // For github.io type URLs, you would set the url and baseUrl like:
  //   url: 'https://facebook.github.io',
  //   baseUrl: '/test-site/',

  // Google Analytics tracking code
  gaTrackingId: 'UA-162306385-1',

  // Used for publishing and more
  projectName: 'gomithrilapp',
  organizationName: 'josephspurrier',
  // For top-level user or org sites, the organization is still the same.
  // e.g., for the https://JoelMarcey.github.io site, it would be set like...
  //   organizationName: 'JoelMarcey'

  // For no header links in the top nav bar -> headerLinks: [],
  headerLinks: [
    {doc: 'tutorial/run-locally', label: 'Docs'},
    {blog: true, label: 'Blog'},
    {href: 'https://petstore.swagger.io/?url=https://raw.githubusercontent.com/josephspurrier/gomithrilapp/master/src/app/ui/static/swagger.json', label: 'API'},
    {page: 'help', label: 'Help'},
    {href: "http://github.com/josephspurrier/gomithrilapp", label: 'GitHub'},
  ],

  // If you have users set above, you add it here:
  users,

  /* path to images for header/footer */
  //headerIcon: 'img/favicon.ico',
  //footerIcon: 'img/favicon.ico',
  //favicon: 'img/favicon.ico',

  /* Colors for website */
  colors: {
    primaryColor: '#3A445D',
    secondaryColor: '#B8336A',
  },

  blogSidebarTitle:{
    default: 'Blog Posts',
    all: 'Blog Posts'
  },
  blogSidebarCount: 'ALL',

  /* Custom fonts for website */
  /*
  fonts: {
    myFont: [
      "Times New Roman",
      "Serif"
    ],
    myOtherFont: [
      "-apple-system",
      "system-ui"
    ]
  },
  */

  // This copyright info is used in /core/Footer.js and blog RSS/Atom feeds.
  copyright: `Copyright © ${new Date().getFullYear()} Joseph Spurrier`,

  highlight: {
    // Highlight.js theme to use for syntax highlighting in code blocks.
    theme: 'hybrid',
  },

  // Add custom scripts here that would be placed in <script> tags.
  scripts: ['https://buttons.github.io/buttons.js'],

  // On page navigation for the current documentation page.
  onPageNav: 'separate',
  // No .html extensions for paths.
  cleanUrl: true,

  // Open Graph and Twitter card images.
  ogImage: 'img/undraw_online.svg',
  twitterImage: 'img/undraw_tweetstorm.svg',

  // For sites with a sizable amount of content, set collapsible to true.
  // Expand/collapse the links and subcategories under categories.
  // docsSideNavCollapsible: true,

  // Show documentation's last contributor's name.
  // enableUpdateBy: true,

  // Show documentation's last update time.
  // enableUpdateTime: true,

  // You may provide arbitrary config keys to be used as needed by your
  // template. For example, if you need your repo's URL...
  repoUrl: 'https://github.com/josephspurrier/gomithrilapp',
};

module.exports = siteConfig;
