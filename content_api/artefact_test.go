package content_api_test

import (
	. "github.com/alphagov/metadata-api/content_api"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Artefact", func() {
	Describe("ParseArtefact", func() {
		It("returns an error when it can't parse the string", func() {
			artefact, err := ParseArtefactResponse([]byte(""))

			Expect(err).ToNot(BeNil())
			Expect(artefact).To(BeNil())
		})

		It("parses a Content API response into an Artefact struct", func() {
			exampleArtefactResponse := `{
  "id": "https://www.gov.uk/api/driving-licence-fees.json",
  "web_url": "https://www.gov.uk/driving-licence-fees",
  "title": "Driving licence fees",
  "format": "answer",
  "updated_at": "2014-06-27T14:21:48+01:00",
  "tags": [],
  "related": [],
  "details": {
    "need_ids": ["100567"],
    "business_proposition": true,
    "description": "this is a test artefact",
    "language": "en",
    "need_extended_font": false,
    "alternative_title": "",
    "body": "foo",
    "parts": [
      {
		"web_url": "https://www.gov.uk/housing-benefit/overview",
		"slug": "overview",
		"order": 1,
		"title": "Overview",
		"body": "<p>You could get Housing Benefit to help you pay your rent.</p>"
      },
      {
		"web_url": "https://www.gov.uk/housing-benefit/what-youll-get",
		"slug": "what-youll-get",
		"order": 2,
		"title": "What you'll get",
		"body": "<p>You may get help with all or part of your rent.</p> "
      }
    ]
  },
  "_response_info": {
    "status": "ok"
  }
}`

			artefact, err := ParseArtefactResponse([]byte(exampleArtefactResponse))

			Expect(err).To(BeNil())
			Expect(artefact).To(Equal(&Artefact{
				ID:     "https://www.gov.uk/api/driving-licence-fees.json",
				WebURL: "https://www.gov.uk/driving-licence-fees",
				Title:  "Driving licence fees",
				Format: "answer",
				Details: Detail{
					NeedIDs:             []string{"100567"},
					BusinessProposition: true,
					Description:         "this is a test artefact",
	                Parts: []Part{
	                    Part{
	                        WebURL: "https://www.gov.uk/housing-benefit/overview",
	                        Title: "Overview",
	                    },
	                    Part{
	                        WebURL: "https://www.gov.uk/housing-benefit/what-youll-get",
	                        Title: "What you'll get",
	                    },
	                },
				},
			}))
		})
	})
})
