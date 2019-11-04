const HCCrawler = require('headless-chrome-crawler');

const sitesFilename = __dirname + "/sites.txt";
const maxDepth = 2;

(async () => {
  const crawler = await HCCrawler.launch({
    customCrawl: (async (page, crawl) => {
      page.on('requestfinished', request => {
        for (h in request.headers) {
          console.log(h + "=" + request.headers[h]);
        }
      });
      return crawl();
    }),
    onSuccess: (result => {
      for (h in result.response.headers) {
        console.log(h + "=" + result.response.headers[h]);
      }
    }),
    onError: (error => {
      console.log(error);
    }),
  });

  var lineReader = require('readline').createInterface({
    input: require('fs').createReadStream(sitesFilename)
  });

  lineReader.on('line', function (line) {
    crawler.queue({url: line, maxDepth: maxDepth})
  });

  await crawler.onIdle();
  await crawler.close();
})();
