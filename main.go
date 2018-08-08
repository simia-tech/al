package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", index)
	r.Mount("/js", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	r.Mount("/wasm", http.StripPrefix("/wasm/", http.FileServer(http.Dir("build"))))

	s := http.Server{
		Addr:    "localhost:4000",
		Handler: r,
	}

	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, `<DOCTYPE html>
<html>
  <head>
    <title>Tests</title>
    <script src="/js/wasm_exec.js"></script>
    <script type="text/javascript">
      function fetchAndInstantiate(url, importObject) {
        return fetch(url).then(response =>
          response.arrayBuffer()
        ).then(bytes =>
          WebAssembly.instantiate(bytes, importObject)
        ).then(results =>
          results.instance
        );
      }
      var go = new Go();
      var mod = fetchAndInstantiate("/wasm/test.wasm", go.importObject);
      window.onload = function() {
        mod.then(function(instance) {
          go.run(instance);
        });
      };
    </script>
  </head>
  <body>
    <h1>Tests</h1>
  </body>
</html>
`)
}
