// GENERATED CODE - DO NOT EDIT
// This file provides a way of creating URL's based on all the actions
// found in all the controllers.
package routes

import "github.com/revel/revel"


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).URL
}

func (_ tStatic) ServeDir(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeDir", args).URL
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).URL
}

func (_ tStatic) ServeModuleDir(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModuleDir", args).URL
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).URL
}

func (_ tTestRunner) Suite(
		suite string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	return revel.MainRouter.Reverse("TestRunner.Suite", args).URL
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).URL
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).URL
}


type tAdmin struct {}
var Admin tAdmin


func (_ tAdmin) Administration(
		begin_date_input interface{},
		end_date_input interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "begin_date_input", begin_date_input)
	revel.Unbind(args, "end_date_input", end_date_input)
	return revel.MainRouter.Reverse("Admin.Administration", args).URL
}


type tApp struct {}
var App tApp


func (_ tApp) Login(
		message string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "message", message)
	return revel.MainRouter.Reverse("App.Login", args).URL
}

func (_ tApp) Admin(
		uid string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "uid", uid)
	return revel.MainRouter.Reverse("App.Admin", args).URL
}

func (_ tApp) User(
		uid string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "uid", uid)
	return revel.MainRouter.Reverse("App.User", args).URL
}

func (_ tApp) Auth(
		email string,
		password string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "email", email)
	revel.Unbind(args, "password", password)
	return revel.MainRouter.Reverse("App.Auth", args).URL
}

func (_ tApp) Inscription(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Inscription", args).URL
}

func (_ tApp) SignIn(
		nom string,
		prenom string,
		email string,
		password string,
		phone string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "nom", nom)
	revel.Unbind(args, "prenom", prenom)
	revel.Unbind(args, "email", email)
	revel.Unbind(args, "password", password)
	revel.Unbind(args, "phone", phone)
	return revel.MainRouter.Reverse("App.SignIn", args).URL
}


type tClient struct {}
var Client tClient


func (_ tClient) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Client.Index", args).URL
}

func (_ tClient) Facture(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Client.Facture", args).URL
}

func (_ tClient) ProcessDemande(
		address string,
		date string,
		hour string,
		motif string,
		phone string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "address", address)
	revel.Unbind(args, "date", date)
	revel.Unbind(args, "hour", hour)
	revel.Unbind(args, "motif", motif)
	revel.Unbind(args, "phone", phone)
	return revel.MainRouter.Reverse("Client.ProcessDemande", args).URL
}

func (_ tClient) DeleteDemande(
		id int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("Client.DeleteDemande", args).URL
}

func (_ tClient) Demande(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Client.Demande", args).URL
}


