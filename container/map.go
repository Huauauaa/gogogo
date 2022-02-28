package container

import "fmt"

func mapDemo() {
    scene := make(map[string]int)
    scene["route"] = 66
    scene["brazil"] = 4
    scene["china"] = 960

    for k, v := range scene {
        fmt.Println(k, v)
    }

    delete(scene, "brazil")

    for k := range scene {
        fmt.Println(k)
    }
}
