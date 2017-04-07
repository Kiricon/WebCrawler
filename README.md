# WebCrawler
A Go implementation of a web crawler.

This web crawler recursively concurrently crawls web pages and stores the links that it finds in a local text file.
Writing to a local file will not work with larger domains due to high levels of concurrecy and OS limits on number of read and write operations to a single file.

Terminal show stats regarding the number of pages found and the number ofoncurrent web crawlers running.
