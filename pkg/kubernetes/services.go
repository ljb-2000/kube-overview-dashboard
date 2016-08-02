package kubernetes

func GetServicesNumber() int {
    return len(svc.Items)
}


func GetServicesNamesPerNamespaces() int {
    // Services
    // svc by namespaces
    sn := make(map[string]int)
    for _, s := range svc.Items {
        sn[s.Namespace]+=1
        //fmt.Printf("%s\n", s.Name)
    }
    return len(svc.Items)
}
