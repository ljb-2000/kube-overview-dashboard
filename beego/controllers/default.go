package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
	kubernetes "github.com/kube-overview-dashboard/pkg/kubernetes"
)

type MainController struct {
	beego.Controller
}

type NodesController struct {
	beego.Controller
}

type ErrorController struct {
    beego.Controller
}


var kubeToken string = "djxCNG1qdrl47vzrE8tZw4nrb2uJzael"

func (c *ErrorController) Error404() {
	c.TplName = "error.tpl"
	c.Data["error"] = "Page not found"
	c.Data["code"] = "404"
}

func (c *NodesController) Get() {
	c.Data["title"] = "Nodes Overview"

	kubernetes.Connect(kubeToken)

  // Nodes overview template
	tnodes := make(map[string]map[string]string)
	tnodes = kubernetes.GetNodesLabels()

	c.Data["nodes"] = tnodes
	c.Layout = "layout.html"
	c.TplName = "nodes-overview.tpl"
}

func (c *MainController) Get() {
	c.Data["title"] = "Kubernetes Overview Dashboard"

	kubernetes.Connect(kubeToken)

  // Nodes template
	tnodes := make(map[string][]string)
	tnodes["number"] = []string{strconv.Itoa(kubernetes.GetNodesNumber())}
	tnodes["numberAvailable"] = []string{strconv.Itoa(kubernetes.GetNodesAvailableNumber())}
	tnodes["names"] = kubernetes.GetNodesName()

  tservices := make(map[string][]string)
  tservices["number"] = []string{strconv.Itoa(kubernetes.GetServicesNumber())}

  tpods := make(map[string][]string)
	tpods["number"] = []string{strconv.Itoa(kubernetes.GetPodsNumber())}

	tingress := make(map[string][]string)
	tingress["number"] = []string{strconv.Itoa(kubernetes.GetPodsNumber())}

	c.Data["nodes"] = tnodes
	c.Data["services"] = tservices
	c.Data["pods"] = tpods
	c.Data["ingress"] = tingress

	c.Layout = "layout.html"
	c.TplName = "home.tpl"
}
