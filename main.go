package main

import (
    tm "github.com/buger/goterm"
    "time"
    "fmt"
    "github.com/fsouza/go-dockerclient"
)

func prepare() {
    tm.Clear()
    tm.MoveCursor(0,0)
}

func monitor(containers []docker.APIContainers) {
	results := tm.NewTable(0, 5, 5, ' ', 0)
	fmt.Fprintf(results, "CONTAINER ID\tIMAGE\tCOMMAND\tCREATED\tSTATUS\tSTATE\tNAMES\n")
	for _, container := range containers {
    	fmt.Fprintf(results, "%s\t%s\t%s\t%s\t%s\t%s\t%s\t\n", container.ID[:12], container.Image, container.Command, time.Since(time.Unix(container.Created, 0)), container.Status, container.State, container.Names[0][1:])
	}

	tm.Println(results)
}

func flush() {
    tm.Flush()
    time.Sleep(time.Second)
}

func main() {
    
    endpoint := "unix:///var/run/docker.sock"
	client, err := docker.NewClient(endpoint)
	if err != nil {
		panic(err)
	}

	for {
        prepare()
	    containers, err := client.ListContainers(docker.ListContainersOptions{All: false})
        if err != nil {
            panic(err)
        }
        monitor(containers)
        flush()
    }

}