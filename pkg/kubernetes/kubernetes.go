package kubernetes

import (
  client "k8s.io/kubernetes/pkg/client/unversioned"
  "os"
  "k8s.io/kubernetes/pkg/api"
  "k8s.io/kubernetes/pkg/apis/extensions"
  "k8s.io/kubernetes/pkg/client/restclient"
  "log"
)

var c *client.Client
var nodes *api.NodeList
var svc *api.ServiceList
var pods *api.PodList
var ingress *extensions.IngressList

func Connect(token string, ) {
  log.SetOutput(os.Stdout)

  config := restclient.Config{
      Host:        "https://kubetest.com:8443",
      BearerToken: string(token),
      Insecure: true,
  }

  var err error
  c, err = client.New(&config)
  checkErr(err, "Can't connect to Kubernetes API:")

  nodes, err = c.Nodes().List(api.ListOptions{})
  checkErr(err, "Can't get nodes:")

  svc, err = c.Services(api.NamespaceAll).List(api.ListOptions{})
  checkErr(err, "Can't get service:")

  pods, err = c.Pods(api.NamespaceAll).List(api.ListOptions{})
  checkErr(err, "Can't get pods:")

  ingress, err = c.Ingress(api.NamespaceAll).List(api.ListOptions{})
  checkErr(err, "Can't get ingress:")
}

func checkErr(err error, m string){
  if err != nil {
      log.Fatalln(m, err)
  }
}
