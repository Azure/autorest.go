package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"tests/acceptancetests/utils"
	"tests/generated/azurereport"
	"tests/generated/report"
)

const testServerPath = "../../../node_modules/@microsoft.azure/autorest.testserver"

func main() {
	err := startServer()
	if err != nil {
		panic(fmt.Sprintf("Error starting server: %v\n", err))
	}
	allPass := true
	runTests(&allPass)
	getReport()
	getAzureReport()
	err = stopServer()
	if err != nil {
		fmt.Printf("Error stopping server: %v\n", err)
	}
	if !allPass {
		fmt.Println("Not all tests passed")
		os.Exit(1)
	}
}

func startServer() error {
	fmt.Println("Go Tests.......")
	install := exec.Command("npm", "install")
	install.Dir = testServerPath
	server := exec.Command("npm", "start")
	server.Dir = testServerPath

	if err := install.Run(); err != nil {
		return err
	}
	return server.Start()
}

func stopServer() error {
	server := exec.Command("npm", "stop")
	server.Dir = testServerPath
	return server.Run()
}

func runTests(allPass *bool) {
	fmt.Println("Run tests")
	testSuites := []string{
		"arraygroup",
		"booleangroup",
		"bytegroup",
		//"complexgroup",
		//"dategroup",
		//"datetimegroup",
		//"datetimerfc1123group",
		//"dictionarygroup",
		//"durationgroup",
		//"headergroup",
		"httpInfrastructuregroup",
		"integergroup",
		//"modelflatteninggroup",
		//"numbergroup",
		//"requiredoptionalgroup",
		"stringgroup",
		//"urlgroup",
		//"validationgroup",
		//"custombaseurlgroup",
		//"filegroup",
		//"formdatagroup",
		//"paginggroup",
		//"morecustombaseurigroup",
	}

	for _, suite := range testSuites {
		fmt.Printf("Run test (go test ./acceptancetests/%vtest -v) ...\n", suite)
		tests := exec.Command("go", "test", fmt.Sprintf("./acceptancetests/%vtest", suite), "-v")
		var stdout, stderr bytes.Buffer
		tests.Stdout, tests.Stderr = &stdout, &stderr
		err := tests.Run()
		fmt.Println(stdout.String())
		fmt.Println(stderr.String())
		if err != nil {
			fmt.Printf("Error! %v\n", err)
			*allPass = false
		}
		if len(stderr.String()) >= 2 && stderr.String()[:2] != "OK" {
			*allPass = false
		}
	}
}

func getReport() {
	var reportClient = report.NewManagementClient(utils.NewPipeline())
	res, err := reportClient.GetReport(context.Background(), nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
	printReport(res.Value, "")
}

func getAzureReport() {
	var reportClient = azurereport.NewManagementClient(utils.NewPipeline())
	res, err := reportClient.GetReport(context.Background(), nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
	printReport(res.Value, "Azure")
}

func printReport(res map[string]int32, report string) {
	count := 0
	for key, val := range res {
		if val <= 0 {
			fmt.Println(key, val)
			count++
		}
	}
	total := len(res)
	fmt.Printf("\nReport:	Passed(%v)  Not Run(%v)\n", total-count, count)
	fmt.Printf("Go %s Done.......\n\n", report)
}
