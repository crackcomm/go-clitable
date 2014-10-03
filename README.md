# Table in Console for Golang

Example table:

```
|----------------------------------------------------------------------------------------|
| Name              | Host                 | Type             | _id                      |
|----------------------------------------------------------------------------------------|
| MongoLab          | mongolab.com         | MongoDB Provider | 52518c5d56357d17ec000002 |
|----------------------------------------------------------------------------------------|
| Google App Engine | appengine.google.com | App Engine       | 52518ff356357d17ec000004 |
|----------------------------------------------------------------------------------------|
| Heroku            | heroku.com           | App Engine       | 5251918e56357d17ec000005 |
|----------------------------------------------------------------------------------------|
```

## Horizontal table

Example table:

```
|---------------------------------|
| Name | MongoLab                 |
|---------------------------------|
| Host | mongolab.com             |
|---------------------------------|
| Type | MongoDB Provider         |
|---------------------------------|
| _id  | 52518c5d56357d17ec000002 |
|---------------------------------|
```

```Go
table.PrintHorizontal(map[string]string{
	"Name": "MongoLab",
	"Host": "mongolab.com",
	// ...
})
```

## Markdown table

```
table := New(...)
table.Markdown = true
table.Print()
```


Example table:

```
| Name              | Host                 | Type             | _id                      |
| ----------------- | -------------------- | ---------------- | ------------------------ |
| MongoLab          | mongolab.com         | MongoDB Provider | 52518c5d56357d17ec000002 |
| Google App Engine | appengine.google.com | App Engine       | 52518ff356357d17ec000004 |
| Heroku            | heroku.com           | App Engine       | 5251918e56357d17ec000005 |
```

| Name              | Host                 | Type             | _id                      |
| ----------------- | -------------------- | ---------------- | ------------------------ |
| MongoLab          | mongolab.com         | MongoDB Provider | 52518c5d56357d17ec000002 |
| Google App Engine | appengine.google.com | App Engine       | 52518ff356357d17ec000004 |
| Heroku            | heroku.com           | App Engine       | 5251918e56357d17ec000005 |
