curl --location --request POST 'http://localhost:8080/' \
--header 'Content-Type: application/cloudevents+json' \
--data-raw '{
	"specversion": "1.0",
	"type": "dev.knative.samples.helloworld",
	"source": "dev.knative.samples/helloworldsource",
	"id": "536808d3-88be-4077-9d7a-a3f162705f79",
	"data": {
		"msg": "Hello Knative2!"
	}
}'