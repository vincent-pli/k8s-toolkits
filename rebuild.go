package main

import (
	"context"
	"flag"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"

	cloudevents "github.com/cloudevents/sdk-go"
	"github.com/tidwall/gjson"
)

var (
	eventType string
)

type Details struct {
	repoName     string
	repoURL      string
	revision     string
	filesChanged []string
	files        string
}

func init() {
	flag.StringVar(&eventType, "type", "dev.knative.source.gitlab.Push Hook", "Watches for this CloudEvent Type.")
}

func Run(cmd string, args ...string) (string, error) {
	out, err := exec.Command(cmd, args...).CombinedOutput()
	if err != nil {
		log.Print("Failed running: %s %s\n%s\n",
			cmd, strings.Join(args, " "), err)
	}
	return string(out), err
}

func receive(event cloudevents.Event, resp *cloudevents.EventResponse) {
	log.Printf("Received CloudEvent,\n%s", event)
	if event.Type() == eventType {
		details := parse(event)
		replace(details)
		triggerPipeline()
	}
}

func parse(event cloudevents.Event) Details {
	details := Details{}
	repoName := gjson.GetBytes(event.Data.([]byte), "repository.name").Raw
	repoURL := gjson.GetBytes(event.Data.([]byte), "repository.url").Raw
	revision := gjson.GetBytes(event.Data.([]byte), "after").Raw
	files := gjson.GetBytes(event.Data.([]byte), "commits.0.modified").Raw
	// revision := event.Data.after
	// files := event.Data.commits[0].modified
	details.repoName = repoName
	details.repoURL = repoURL
	details.revision = revision
	details.files = files
	log.Printf("Parsed CloudEvent,\n%s", details)

	return details
}

func replace(details Details) {

}

func triggerPipeline() {
	token, err := ioutil.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/token")
	if err != nil {
		log.Print("Can't load API token:%s\n", err)
		return
	}

	args := []string{"kubectl", "--token=" + string(token),
		"apply", "-f", "/gitresource.yaml"}
	if out, err := Run(args[0], args[1:]...); err != nil {
		log.Print("Error create git resource: %s\n%s\n", out, err)
		return
	}

	args = []string{"kubectl", "--token=" + string(token),
		"apply", "-f", "/pipelinerun.yaml"}
	if out, err := Run(args[0], args[1:]...); err != nil {
		log.Print("Error create pipelinerun: %s\n%s\n", out, err)
		return
	}
}

func main() {
	ce, err := cloudevents.NewDefaultClient()
	if err != nil {
		log.Fatalf("failed to create CloudEvent client, %s", err)
	}

	log.Fatal(ce.StartReceiver(context.Background(), receive))
}
