# JSON lint

```
    go get github.com/moxar/jsonlint
    echo '{"foo": "bar", "lorem": 1, "ipsum": false, "dolores": null, "ikamor": {"foo": ["bar", "foo", "bar"]}}' | jsonlint
    echo '{"foo": "bar", "lorem": 1, "ipsum": false, "dolores": null, "ikamor": {"foo": ["bar", "foo", "bar"]}}' | jsonlint | pygmentize -l json
```
