# urlsplit

A helper to create FreeMarker template for a path.

Usage:

```
$ urlsplit https://nd.akana.dev:9982/v1/asoiaf/realms/houses/targaryen

2016/11/01 14:51:22 /v1/asoiaf/realms/houses/targaryen
2016/11/01 14:51:22 0
2016/11/01 14:51:22 1 v1
2016/11/01 14:51:22 2 asoiaf
2016/11/01 14:51:22 3 realms
2016/11/01 14:51:22 4 houses
2016/11/01 14:51:22 5 targaryen
<#assign
  url = message.getProperty("http.request.line")?split(" ")
  path = url[1]?split("/")
  id = path[5] <#-- targaryen -->
>
```


