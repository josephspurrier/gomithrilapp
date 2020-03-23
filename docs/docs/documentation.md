---
id: documentation
title: Documentation
---

## Docusaurus

The documentation for the project is generated through [Docusaurus](https://docusaurus.io/) and then published through [GitHub Pages](https://pages.github.com/). All the documentation is located in the [/docs](https://github.com/josephspurrier/govueapp/tree/master/docs) folder.

The documentation is split into three different types of content:

- [blog](https://josephspurrier.github.io/govueapp/blog/) - these are blog articles that are generated from markdown files [here](https://github.com/josephspurrier/govueapp/tree/master/docs/website/blog).
- [docs](https://josephspurrier.github.io/govueapp/docs/tutorial/run-locally) - these are guides and tutorials that are generated from markdown files [here](https://github.com/josephspurrier/govueapp/tree/master/docs/docs).
- [pages](https://josephspurrier.github.io/govueapp/help) - these are single pages like the homepage and help page that generated from [React](https://docusaurus.io/docs/en/api-pages) components [here](https://github.com/josephspurrier/govueapp/tree/master/docs/website/pages/en).

The [siteConfig.js](https://github.com/josephspurrier/govueapp/blob/master/docs/website/siteConfig.js) file contains all of the site configurations like the baseURL, site title, Google Analytics code, etc. You can also customize the top navigation bar via the `headerLinks` field.

The [sidebars.js](https://github.com/josephspurrier/govueapp/blob/master/docs/website/sidebars.json) file contains the [Sidebar](https://docusaurus.io/docs/en/navigation) configurations. It will create the left sidebar with relative pages when a page loads.

The [Footer.js](https://github.com/josephspurrier/govueapp/blob/master/docs/website/core/Footer.js) file contains the React component for the footer that is visible on every page.

The [static](https://github.com/josephspurrier/govueapp/tree/master/docs/website/static) folder contains the static assets like CSS and images used throughout the documentation.

### Install the npm Dependencies

Install the npm dependencies for building the documentation from Docusaurus.

```bash
# Makefile
make doc-dep

# Manual
cd ${GOPATH}/docs/website
npm install
```

### Run the Docusaurus Live Server

Run the server that displays the documentation locally. Changes are made without having to reload, but you must kill and re-run the commands when you make changes to the **sidebars.js** file.

```bash
# Makefile
make doc-dev

# Manual
cd ${GOPATH}/docs/website
npm start
```

### Publish the Documentation to GitHub Pages

Push the changes from the master branch to GitHub pages for the specified user.

```bash
# Makefile
make doc-publish

# Manual
cd ${GOPATH}/docs/website
GIT_USER=josephspurrier CURRENT_BRANCH=master USE_SSH=true npm run publish-gh-pages
```