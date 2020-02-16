package main


import (
    "flag"
    "strings"
    "fmt"
    "io/ioutil"
    "gopkg.in/yaml.v2"
    "network"
)


type NetworkConfigDefaults struct {
    Defaults struct {
        Domain string `yaml:"domain"`
        Dns []string `yaml:"dns"`
    } `yaml:"defaults"`
}

type NetworkConfig struct {
    Hosts []struct {
        Name string `yaml:"name"`
        Vips []struct {
            Address string `yaml:"address"`
            Network string `yaml:"network"`
        } `yaml:"vips"`
        Interfaces []struct {
            Name string `yaml:"name"`
            Interface string `yaml:"interface"`
            Address string `yaml:"address"`
            Gateway string `yaml:"gateway"`
        } `yaml:"interfaces"`
    } `yaml:"hosts"`
}




func main() {

    hostname := flag.String("hostname", "dns-1", "The hostname to generate the config for")
    environment := flag.String("env", "prod", "The environment you'r host running in")

    flag.Parse()

    nodeParts := strings.Split(*hostname, "-")
    nodeNumberA, nodeRoleA := nodeParts[len(nodeParts)-1], nodeParts[:len(nodeParts)-1]

    nodeRole := strings.Join(nodeRoleA, "-")
    nodeNumber := nodeNumberA



    networkFile := "./config/env/" + *environment + "/network.yaml"
    usersFile := "./config/env/" + *environment + "/users.yaml"
    roleFile := "./config/roles/" + nodeRole + ".yaml"



    networkFileContent, err := ioutil.ReadFile(networkFile)
    if err != nil {
        fmt.Printf("Error reading YAML file: %s\n", err)
        return
    }
    //fmt.Println("networkFileContent:", string(networkFileContent))


    var networkConfigDefaults NetworkConfigDefaults
    err = yaml.Unmarshal(networkFileContent, &networkConfigDefaults)
    if err != nil {
        fmt.Printf("Error parsing YAML file: %s\n", err)
    }

    fmt.Printf("networkConfigDefaults: %v\n", networkConfigDefaults)



    var networkConfig NetworkConfig
    err = yaml.Unmarshal(networkFileContent, &networkConfig)
    if err != nil {
        fmt.Printf("Error parsing YAML file: %s\n", err)
    }

    fmt.Printf("networkConfig: %v\n", networkConfig)


    fmt.Println("")
    fmt.Println("")
    fmt.Println("")
    fmt.Println("")
    fmt.Println("networkFile:", networkFile)
    fmt.Println("usersFile:", usersFile)
    fmt.Println("roleFile:", roleFile)
    fmt.Println("nodeNumber:", nodeNumber)
    fmt.Println("nodeRole:", nodeRole)
    fmt.Println("hostname:", *hostname)
    fmt.Println("environment:", *environment)
    fmt.Println("tail:", flag.Args())


}