package kubernetes

import (
  "fmt"
  "strings"
  "sort"
)

func GetNodesNumber() int {
    return len(nodes.Items)
}

func GetNodesAvailableNumber() int {
    available_nodes := 0
    for _, node := range nodes.Items {
        for _, c := range node.Status.Conditions {
          fmt.Printf("   %s %s\n", c.Type, c.Status)
          if ((c.Type == "Ready") && (c.Status == "True")) {
              available_nodes += 1
          }
        }
    }
    return available_nodes
}


func GetNodesName() []string {
    var nn []string //nodesname
    for _, node := range nodes.Items {
        nn = append(nn, node.Name)
    }
    return nn
}

func GetNodesLabels() map[string]map[string]string {
    nl := make(map[string]map[string]string) // Nodes labels

    for _, node := range nodes.Items {
        nl[node.Name] = make (map[string]string)
        var labels []string
        for k, v := range node.Labels {
            labels = append(labels, k+"="+v)
        }
        sort.Strings(labels)
        nl[node.Name]["labels"] = strings.Join(labels, " ")
    }
    return nl
}
