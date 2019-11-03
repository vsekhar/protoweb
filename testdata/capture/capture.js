const puppeteer = require('puppeteer');
var argv = require('minimist')(process.argv.slice(2));

(async () => {
    const browser = await puppeteer.launch({
      headless: true,
      executablePath: argv['chromepath']
    });
    const page = await browser.newPage();
    await page.setViewport({ width: 1366, height: 768});
    page.on('response', async (response) => {
      if (response.status() >= 300 && response.status() < 400) {
        // ignore redirects
        return;
      }
      reqheaders = response.request().headers();
      for (header in reqheaders) {
        console.log(header + "=" + reqheaders[header]);
      }
      respheaders = response.headers();
      for (header in respheaders) {
        console.log(header + "=" + respheaders[header]);
      }
    });
    await page.goto('https://nytimes.com', {
      waitUntil: 'networkidle2'
    });
  
    await browser.close();
  })();
