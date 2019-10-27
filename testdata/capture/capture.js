const puppeteer = require('puppeteer-core');
const path = require('path');
const fse = require('fs-extra');
var argv = require('minimist')(process.argv.slice(2));

(async () => {
    const browser = await puppeteer.launch({
      headless: true,
      executablePath: argv['chromepath']
    });
    const page = await browser.newPage();
    await page.setViewport({ width: 1366, height: 768});
    page.on('requestfinished', async (request) => {
      const url = new URL(response.url());

    });
    page.on('response', async (response) => {
      // redirects
      if (response.status() >= 300 && response.status() < 400) {
        return;
      }
      const url = new URL(response.url());
      let filePath = path.resolve(`./output${url.pathname}`);
      if (path.extname(url.pathname).trim() === '') {
        filePath = `${filePath}/index.html`;
      }
      await fse.outputFile(filePath, await response.buffer());
    });
    await page.goto('https://google.com', {
      waitUntil: 'networkidle2'
    });
  
    await browser.close();
  })();
