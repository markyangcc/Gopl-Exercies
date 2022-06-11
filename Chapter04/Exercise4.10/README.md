## reference
https://docs.github.com/en/rest/search#search-issues-and-pull-requests

https://docs.github.com/en/rest/search

## Other way to solve this
In fact, github had offered the api since:someday

But, handle the time conversion when call the api is a bit complicated to write.
So, I choose to get all the issuse, and filter by myself.

## call github api samples using curl
```bash
curl \
  -H "Accept: application/vnd.github.v3+json" \
  "https://api.github.com/search/issues?q=repo:golang/go+updated:>=2022-06-01"
```

IF you using this way(github api), make sure you call like this,
`go run main.go repo:golang/go is:open updated:\>=2022-06-01`

`\(backslash)` is needed. In Linux, I think greater-than is resolved as a redirect symbol, it's not the way we need.