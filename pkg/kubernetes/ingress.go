package kubernetes

func GetIngressNumber() int {
    return len(ingress.Items)
}

func GetIngressNamesPerNamespaces() int {
    //TODO
    return 0
}
