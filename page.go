package main

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type PageData struct {
	Name string `yaml:"name"`
	Port int    `yaml:"port"`
}

type Page struct {
	Title    string     `yaml:"title"`
	PageData []PageData `yaml:"page_data"`
}

func pageContent() (*Page, error) {
	var page Page
	p, err := filepath.Abs(contentFile)
	if err != nil {
		return &page, err
	}

	content, err := os.ReadFile(p)
	if err != nil {
		return &page, err
	}

	err = yaml.Unmarshal(content, &page)
	if err != nil {
		return &page, err
	}

	return &page, nil
}

var pageTmpl = `<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="UTF-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
  <meta http-equiv="X-UA-Compatible" content="ie=edge" />
  <link rel="stylesheet" href="static/style.css" />
  <title>{{.Title}}</title>
</head>

<body>
  <h1>
    {{.Title}}
  </h1>
  <ul>
    {{range .PageData}}
    <li>
      <a target="_blank" href=":{{.Port}}">{{.Name}}</a>
    </li>
    {{end}}
  </ul>
</body>

</html>


<script>
  document.addEventListener(
    "click",
    function (event) {
      var target = event.target;
      if (target.tagName.toLowerCase() === "a") {
        var port = target.getAttribute("href").match(/^:(\d+)(.*)/);
        if (port) {
          target.href = port[2];
          target.port = port[1];
        }
      }
    },
    false
  );
</script>
`
