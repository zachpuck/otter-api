package main

import "time"

type Otter struct {
    Id          int         `json:"id"`
    Name        string      `json:"name"`
    Cute        bool        `json:"cute"`
    Born        time.Time   `json:"born"`
}

type Otters []Otter