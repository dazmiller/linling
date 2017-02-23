package controllers

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	this.TplName = "release/index.html"
	//this.Layout = "public/layout.html"

}

func (this *HomeController) Prod() {
	this.LayoutSections = make(map[string]string)
	this.TplName = "release/product_view.html"
	this.Layout = "release/layout.html"
	this.LayoutSections["header"] = "release/header.html"
	this.LayoutSections["Content"] = "release/product_view.html"
	this.LayoutSections["Footer"] = "release/footer.html"
}

func (this *HomeController) Prod2() {
	this.LayoutSections = make(map[string]string)
	this.TplName = "release/product_view2.html"
	this.Layout = "release/layout.html"
	this.LayoutSections["header"] = "release/header.html"
	this.LayoutSections["Content"] = "release/product_view2.html"
	this.LayoutSections["Footer"] = "release/footer.html"
}

func (this *HomeController) Prod3() {
	this.LayoutSections = make(map[string]string)
	this.TplName = "release/product_view3.html"
	this.Layout = "release/layout.html"
	this.LayoutSections["header"] = "release/header.html"
	this.LayoutSections["Content"] = "release/product_view3.html"
	this.LayoutSections["Footer"] = "release/footer.html"
}

func (this *HomeController) Cart() {
	this.LayoutSections = make(map[string]string)
	this.TplName = "release/addcart.html"
	this.Layout = "release/layout.html"
	this.LayoutSections["header"] = "release/header.html"
	this.LayoutSections["Feature"] = "release/feature.html"
	this.LayoutSections["Content"] = "release/addcart.html"
	this.LayoutSections["Footer"] = "release/footer.html"
}
