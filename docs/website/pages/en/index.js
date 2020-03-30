/**
 * Copyright (c) 2017-present, Facebook, Inc.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

const React = require('react');

const CompLibrary = require('../../core/CompLibrary.js');

const MarkdownBlock = CompLibrary.MarkdownBlock; /* Used to read markdown */
const Container = CompLibrary.Container;
const GridBlock = CompLibrary.GridBlock;

const fs = require('fs');

class HomeSplash extends React.Component {
  render() {
    const {siteConfig, language = ''} = this.props;
    const {baseUrl, docsUrl} = siteConfig;
    const docsPart = `${docsUrl ? `${docsUrl}/` : ''}`;
    const langPart = `${language ? `${language}/` : ''}`;
    const docUrl = doc => `${baseUrl}${docsPart}${langPart}${doc}`;

    const SplashContainer = props => (
      <div className="homeContainer">
        <div className="homeSplashFade">
          <div className="wrapper homeWrapper">{props.children}</div>
        </div>
      </div>
    );

    const Logo = props => (
      <div className="projectLogo">
        <img src={props.img_src} alt="Project Logo" />
      </div>
    );

    const ProjectTitle = props => (
      <h2 className="projectTitle">
        {props.title}
        <small>{props.tagline}</small>
      </h2>
    );

    const PromoSection = props => (
      <div className="section promoSection">
        <div className="promoRow">
          <div className="pluginRowBlock">{props.children}</div>
        </div>
      </div>
    );

    const Button = props => (
      <div className="pluginWrapper buttonWrapper">
        <a className="button" href={props.href} target={props.target}>
          {props.children}
        </a>
      </div>
    );

    return (
      <SplashContainer>
        <Logo img_src={`${baseUrl}img/undraw_monitor.svg`} />
        <div className="inner">
          <ProjectTitle tagline={siteConfig.tagline} title={siteConfig.title} />
          <PromoSection>
            <Button href="#why">Why this project?</Button>
            <Button href="#tools">Which tools?</Button>
            <Button href={docUrl('tutorial/run-locally.html')}>Getting Started</Button>
          </PromoSection>
        </div>
      </SplashContainer>
    );
  }
}

class Index extends React.Component {
  render() {
    const {config: siteConfig, language = ''} = this.props;
    const {baseUrl} = siteConfig;

    const Block = props => (
      <Container
        padding={['bottom', 'top']}
        id={props.id}
        background={props.background}>
        <GridBlock
          align="center"
          contents={props.children}
          layout={props.layout}
        />
      </Container>
    );

    const WhyCallout = () => (
      <Block id="why" background="dark">
        {[
          {
            content:
              '<div style="text-align:left;">We know it\'s sometimes overwhelming to learn all about your tools before you can work with them. '+
              'We hope you can use this repository as reference material on your development journey. '+
              '<br><br>'+
              'Please feel free to use this project as a template for your web application. It\'s a nice '+
              'way to kick-start your work so you don\'t need to spend time on your framework - the piece we really enjoy. ' +
              '<br><br>'+
              'We chose a sample notepad application because it\'s a concept that is easy to understand and uses many '+
              'of the pieces you will use in any web application you build. This project uses Mithril on the front-end (UI) '+
              ' and Go on the back-end (API). This project is designed to show good development and CI/CD practices as '+
              'well as integrations between '+
              'modern development tools.'+
              '</div>',
            image: `https://user-images.githubusercontent.com/2394539/76177148-ac753e00-6189-11ea-963b-bff38b29e8ed.gif`,
            imageAlign: 'right',
            title: 'Why this project?',
          },
        ]}
      </Block>
    );

    const ToolsCallout = () => {
      const toolsLeft = fs.readFileSync('../docs/homepage/tools-left.md').toString();
      const ToolsRight = fs.readFileSync('../docs/homepage/tools-right.md').toString();

      return (
      <Container id="tools" padding={['bottom', 'top']}>
      <div>
      <h3>Which tools does this project use?</h3>
      </div>
      <div className="gridBlock">
      <div
        className="blockElement imageAlignSide imageAlignLeft twoByGridBlock">
        <MarkdownBlock>
        {toolsLeft}
        </MarkdownBlock>
      </div>
      <div
        className="blockElement imageAlignSide imageAlignLeft twoByGridBlock">
        <MarkdownBlock>
        {ToolsRight}
        </MarkdownBlock>
      </div>
      </div>
    </Container>
      );
    };

    return (
      <div>
        <HomeSplash siteConfig={siteConfig} language={language} />
        <div className="mainContainer">
          <WhyCallout />
          <ToolsCallout />
        </div>
      </div>
    );
  }
}

module.exports = Index;
