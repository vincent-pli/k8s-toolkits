package main

import (
	"context"
	"flag"
	"log"

	"github.com/cloudevents/sdk-go/pkg/cloudevents"
	"github.com/cloudevents/sdk-go/pkg/cloudevents/client"
	"github.com/knative/eventing-sources/pkg/kncloudevents"
)

const (
	CloudEventType = "tekton.reply"
)

// Adapter converts incoming GitLab webhook events to CloudEvents
type Adapter struct {
	client client.Client
}

func main() {
	a := new(Adapter)
	var err error
	a.client, err = kncloudevents.NewDefaultClient(*getSink())

	event := cloudevents.Event{
		Context: cloudevents.EventContextV02{
			ID:   "001",
			Type: CloudEventType,
			// Source:     *source,
			// Extensions: extensions,
		}.AsV02(),
		// Data: payload,
	}
	_, err = a.client.Send(context.TODO(), event)

	if err != nil {
		log.Printf("unexpected error sending GitLab event: %s", err)
	}

}

func getSink() *string {
	sink := flag.String("sink", "", "uri to send events to")

	flag.Parse()

	if sink == nil || *sink == "" {
		log.Fatalf("No sink given")
	}

	return sink
}
