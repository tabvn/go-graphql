package dev

import "net/http"

var Content = []byte(`
<!DOCTYPE html>
<head>
  <title>Go GraphQL</title>
  <style>body {height: 100vh; margin: 0; width: 100%; overflow: hidden;}</style>
  <link rel="stylesheet" href="//cdn.jsdelivr.net/sweetalert/1.1.3/sweetalert.css" />
  <link rel="stylesheet" href="//cdn.jsdelivr.net/graphiql/0.6.3/graphiql.css" />
  <script src="//cdn.jsdelivr.net/sweetalert/1.1.3/sweetalert.min.js"></script>
  <script src="//cdn.jsdelivr.net/fetch/0.9.0/fetch.min.js"></script>
  <script src="//cdn.jsdelivr.net/react/0.14.7/react.min.js"></script>
  <script src="//cdn.jsdelivr.net/react/0.14.7/react-dom.min.js"></script>
  <script src="//cdn.jsdelivr.net/graphiql/0.6.3/graphiql.min.js"></script>
  <script>
    (function () {
      
      document.addEventListener('DOMContentLoaded', function () {
        
          var endpoint = window.location.origin + '/api';
          function fetcher(params) {
            var options = {
              method: 'post',
              headers: {'Accept': 'application/json', 'Content-Type': 'application/json'},
              body: JSON.stringify(params),
              credentials: 'include',
            };
            return fetch(endpoint, options)
              .then(function (res) { return res.json() });
          }
          var body = React.createElement(GraphiQL, {fetcher: fetcher, query: '', variables: ''});
          ReactDOM.render(body, document.body);
      });
    }());
  </script>
</head>
<body>
</body>
`)

// ServeGraphiQL is a handler function for HTTP servers
func ServeGraphiQL(res http.ResponseWriter, req *http.Request) {
	res.Write(Content)
}
