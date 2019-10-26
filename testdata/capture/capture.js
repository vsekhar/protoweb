const puppeteer = require('puppeteer-core');
const commander = require('minimist');
var argv = require('minimist')(process.argv.slice(2));

(async () => {
    const browser = await puppeteer.launch({
      headless: true,
      executablePath: argv['chromepath']
    });
    const page = await browser.newPage();
    await page.setViewport({ width: 1366, height: 768});
    await page.goto('https://google.com');
    await page.screenshot({path: 'google.png'});
  
    await browser.close();
  })();
