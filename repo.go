package main

import "fmt"

var currentId int

var otters Otters

// Give us some seed data
func init() {
    RepoCreateOtter(Otter{Name: "Sea Otter", Cute: true})
    RepoCreateOtter(Otter{Name: "River Otter"})
}

func RepoFindOtter(id int) Otter {
    for _, o := range otters {
        if o.Id == id {
            return o
        }
    }
    // return empty Otter if not found
    return Otter{}
}

func RepoCreateOtter(o Otter) Otter {
    currentId += 1
    o.Id = currentId
    otters = append(otters, o)
    return o
}

func RepoDestroyOtter(id int) error {
    for i, o := range otters {
        if o.Id == id {
            otters = append(otters[:i], otters[i+1:]...)
            return nil
        }
    }
    return fmt.Errorf("Could not find Otter with id of %d to delete", id)
}