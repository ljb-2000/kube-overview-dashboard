package kubernetes

func GetPodsNumber() int {
    return len(pods.Items)
}

func GetPodsNamesPerNamespaces() int {
    //TODO
    return 0
}

func GetPodsNamesPerNodes() int {
    //TODO
    return 0
}
