package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("public", func() {
	Description("A mock service to test different service communication models")
	Error("unauthenticated")
	GRPC(func() {
		Response("unauthenticated", CodeUnauthenticated)
	})
	Method("batchGRPC", func() {
		Description("Receives an array of payloads")
		Payload(TestPayloadBatch)
		Result(ArrayOf(ResponsePayload), "an example response")
		GRPC(func() {
			Response(CodeOK)
		})
	})
	Method("streamedBatchGRPC", func() {
		Description("Receives an array of payloads")
		StreamingPayload(TestPayload)
		StreamingResult(ResponsePayload, "an example response")
		GRPC(func() {
			Response(CodeOK)
		})
	})
})

var TestPayload = Type("test_payload", func() {
	Description("an example payload")
	Field(1, "first_field", String, "")
	Field(2, "second_field", String, "")
	Field(3, "third_field", String, "")
	Field(4, "organization_id", UInt32, "")
	Required("first_field", "second_field", "third_field", "organization_id")
})

var ResponsePayload = Type("response_payload", func() {
	Description("an example payload")
	Field(1, "first_field", String, "")
	Field(2, "fourth_field", String, "")
	Required("first_field", "fourth_field")
})

var TestPayloadBatch = Type("test_payload_batch", func() {
	Description("an example payload")
	Field(1, "records", ArrayOf(TestPayload), "")
})
