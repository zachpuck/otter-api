package main

import "net/http"

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
        Index,
    },
    Route{
        "OtterIndex",
        "GET",
        "/otters",
        OtterIndex,
    },
    Route{
        "OtterShow",
        "GET",
        "/otters/{otterId}",
        OtterShow,
    },
    Route{
        "OtterCreate",
        "POST",
        "/otters",
        OtterCreate,
    },
}