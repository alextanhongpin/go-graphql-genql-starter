package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/alextanhongpin/core/test/testutil"
	"github.com/alextanhongpin/graphql-test/graph"
	"github.com/alextanhongpin/graphql-test/graph/generated"
)

func TestServer(t *testing.T) {
	mux := http.NewServeMux()
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	mux.Handle("/query", testutil.DumpHTTPHandler(t, testutil.IgnoreHeaders("Host"))(srv))
	ts := httptest.NewServer(mux)
	defer ts.Close()

	body := strings.NewReader(`{
  "query": "\n\nquery FindTodo{\n  todos {\n    id \n    text\n  }\n}\n",
  "variables": null,
  "operationName": "FindTodo"
}`)

	res, err := ts.Client().Post(ts.URL+"/query", "application/json", body)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(b))
}
