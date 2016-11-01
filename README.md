# urlsplit

A helper to create FreeMarker template for a path.

Usage:

```
$urlsplit https://nd.akana.dev:9982/api10029live/people/AC1

/api10029live/people/AC1
0
1 api10029live
2 people
3 AC1
<#assign
  url = message.getProperty("http.request.line")?split(" ")
  path = url[1]?split("/")
  id = path[3]
>
```


