const HCCrawler = require('headless-chrome-crawler');

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

  await crawler.queue({
    url: 'https://nytimes.com',
    maxDepth: 2,
  });
  // TODO: load sites.txt and enqueue
  await crawler.onIdle();
  await crawler.close();
})();
