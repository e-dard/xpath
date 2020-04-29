A tiny function for fetching a web-page, applying an xpath expression and then returning the count of nodes as influx line-protocol.

For use with the Telegraf [exec](https://github.com/influxdata/telegraf/tree/master/plugins/inputs/exec) plugin.

Usage:

The function requires three arguments:

  - URL
  - xpath expressino
  - tagset (influx measurment and tag pairs)

```
$ xpath "https://www.example.com" "//button[contains(text(), 'In stock')]" "xpath,retailer=company"
```