package main

import (
    "net/http"

    "github.com/lcraver/games-api/handlers"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
    Route{
        "Index",
        "GET",
        "/",
        handlers.Index,
    },
    Route{
        "GameHandler",
        "GET",
        "/games/{title}",
        handlers.GetByName,
    },
    Route{
        "GameHandlerIndex",
        "GET",
        "/games",
        handlers.GetAll,
    },
}
