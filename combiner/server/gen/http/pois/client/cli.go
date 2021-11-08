// Code generated by goa v3.5.2, DO NOT EDIT.
//
// pois HTTP client CLI support package
//
// Command:
// $ goa gen github.com/derbexuk/poieventservice/server/design

package client

import (
	"encoding/json"
	"fmt"

	pois "github.com/derbexuk/poieventservice/server/gen/pois"
)

// BuildPostPayload builds the payload for the pois post endpoint from CLI
// flags.
func BuildPostPayload(poisPostBody string, poisPostPath string) (*pois.PostPayload, error) {
	var err error
	var body PostRequestBody
	{
		err = json.Unmarshal([]byte(poisPostBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"pois\": [\n         {\n            \"deactivated\": false,\n            \"description\": \"Perferendis consequatur et.\",\n            \"geojson\": \"Voluptates eum enim aut.\",\n            \"geopath\": {\n               \"geopoints\": [\n                  [\n                     0.6467556,\n                     0.28522083,\n                     0.46987006,\n                     0.22842954\n                  ],\n                  [\n                     0.935357,\n                     0.68563277\n                  ]\n               ],\n               \"id\": \"59b287534defd02dace3671b\",\n               \"title\": \"Dolorum qui.\",\n               \"type\": \"Ea enim sunt animi et.\"\n            },\n            \"id\": \"59b287534defd02dace3671b\",\n            \"location\": \"Placeat maxime sunt perferendis vel sunt.\",\n            \"path\": \"Ut voluptas.\",\n            \"properties\": {\n               \"Dolores voluptas sunt et quia voluptatem.\": \"Dolores sed.\",\n               \"Officiis voluptatem corporis voluptatem accusamus dolores.\": \"Saepe molestias in.\",\n               \"Ut aut dolore qui iste asperiores hic.\": \"Mollitia nihil cum et voluptate et suscipit.\"\n            },\n            \"refs\": [\n               \"Aliquid cupiditate.\",\n               \"Nihil sint voluptate incidunt excepturi totam.\"\n            ],\n            \"title\": \"Vitae qui in dolorum officiis laborum dolor.\"\n         },\n         {\n            \"deactivated\": false,\n            \"description\": \"Perferendis consequatur et.\",\n            \"geojson\": \"Voluptates eum enim aut.\",\n            \"geopath\": {\n               \"geopoints\": [\n                  [\n                     0.6467556,\n                     0.28522083,\n                     0.46987006,\n                     0.22842954\n                  ],\n                  [\n                     0.935357,\n                     0.68563277\n                  ]\n               ],\n               \"id\": \"59b287534defd02dace3671b\",\n               \"title\": \"Dolorum qui.\",\n               \"type\": \"Ea enim sunt animi et.\"\n            },\n            \"id\": \"59b287534defd02dace3671b\",\n            \"location\": \"Placeat maxime sunt perferendis vel sunt.\",\n            \"path\": \"Ut voluptas.\",\n            \"properties\": {\n               \"Dolores voluptas sunt et quia voluptatem.\": \"Dolores sed.\",\n               \"Officiis voluptatem corporis voluptatem accusamus dolores.\": \"Saepe molestias in.\",\n               \"Ut aut dolore qui iste asperiores hic.\": \"Mollitia nihil cum et voluptate et suscipit.\"\n            },\n            \"refs\": [\n               \"Aliquid cupiditate.\",\n               \"Nihil sint voluptate incidunt excepturi totam.\"\n            ],\n            \"title\": \"Vitae qui in dolorum officiis laborum dolor.\"\n         },\n         {\n            \"deactivated\": false,\n            \"description\": \"Perferendis consequatur et.\",\n            \"geojson\": \"Voluptates eum enim aut.\",\n            \"geopath\": {\n               \"geopoints\": [\n                  [\n                     0.6467556,\n                     0.28522083,\n                     0.46987006,\n                     0.22842954\n                  ],\n                  [\n                     0.935357,\n                     0.68563277\n                  ]\n               ],\n               \"id\": \"59b287534defd02dace3671b\",\n               \"title\": \"Dolorum qui.\",\n               \"type\": \"Ea enim sunt animi et.\"\n            },\n            \"id\": \"59b287534defd02dace3671b\",\n            \"location\": \"Placeat maxime sunt perferendis vel sunt.\",\n            \"path\": \"Ut voluptas.\",\n            \"properties\": {\n               \"Dolores voluptas sunt et quia voluptatem.\": \"Dolores sed.\",\n               \"Officiis voluptatem corporis voluptatem accusamus dolores.\": \"Saepe molestias in.\",\n               \"Ut aut dolore qui iste asperiores hic.\": \"Mollitia nihil cum et voluptate et suscipit.\"\n            },\n            \"refs\": [\n               \"Aliquid cupiditate.\",\n               \"Nihil sint voluptate incidunt excepturi totam.\"\n            ],\n            \"title\": \"Vitae qui in dolorum officiis laborum dolor.\"\n         },\n         {\n            \"deactivated\": false,\n            \"description\": \"Perferendis consequatur et.\",\n            \"geojson\": \"Voluptates eum enim aut.\",\n            \"geopath\": {\n               \"geopoints\": [\n                  [\n                     0.6467556,\n                     0.28522083,\n                     0.46987006,\n                     0.22842954\n                  ],\n                  [\n                     0.935357,\n                     0.68563277\n                  ]\n               ],\n               \"id\": \"59b287534defd02dace3671b\",\n               \"title\": \"Dolorum qui.\",\n               \"type\": \"Ea enim sunt animi et.\"\n            },\n            \"id\": \"59b287534defd02dace3671b\",\n            \"location\": \"Placeat maxime sunt perferendis vel sunt.\",\n            \"path\": \"Ut voluptas.\",\n            \"properties\": {\n               \"Dolores voluptas sunt et quia voluptatem.\": \"Dolores sed.\",\n               \"Officiis voluptatem corporis voluptatem accusamus dolores.\": \"Saepe molestias in.\",\n               \"Ut aut dolore qui iste asperiores hic.\": \"Mollitia nihil cum et voluptate et suscipit.\"\n            },\n            \"refs\": [\n               \"Aliquid cupiditate.\",\n               \"Nihil sint voluptate incidunt excepturi totam.\"\n            ],\n            \"title\": \"Vitae qui in dolorum officiis laborum dolor.\"\n         }\n      ]\n   }'")
		}
	}
	var path string
	{
		path = poisPostPath
	}
	v := &pois.PostPayload{}
	if body.Pois != nil {
		v.Pois = make([]*pois.PoiPayload, len(body.Pois))
		for i, val := range body.Pois {
			v.Pois[i] = marshalPoiPayloadRequestBodyToPoisPoiPayload(val)
		}
	}
	v.Path = path

	return v, nil
}

// BuildShowPayload builds the payload for the pois show endpoint from CLI
// flags.
func BuildShowPayload(poisShowPath string) (*pois.ShowPayload, error) {
	var path string
	{
		path = poisShowPath
	}
	v := &pois.ShowPayload{}
	v.Path = path

	return v, nil
}

// BuildListByPathPayload builds the payload for the pois ListByPath endpoint
// from CLI flags.
func BuildListByPathPayload(poisListByPathPath string) (*pois.ListByPathPayload, error) {
	var path string
	{
		path = poisListByPathPath
	}
	v := &pois.ListByPathPayload{}
	v.Path = path

	return v, nil
}

// BuildListByReferencePayload builds the payload for the pois ListByReference
// endpoint from CLI flags.
func BuildListByReferencePayload(poisListByReferencePath string) (*pois.ListByReferencePayload, error) {
	var path string
	{
		path = poisListByReferencePath
	}
	v := &pois.ListByReferencePayload{}
	v.Path = path

	return v, nil
}

// BuildUpdatePayload builds the payload for the pois update endpoint from CLI
// flags.
func BuildUpdatePayload(poisUpdateBody string, poisUpdatePath string) (*pois.UpdatePayload, error) {
	var err error
	var body UpdateRequestBody
	{
		err = json.Unmarshal([]byte(poisUpdateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"poi\": {\n         \"deactivated\": false,\n         \"description\": \"Perferendis consequatur et.\",\n         \"geojson\": \"Voluptates eum enim aut.\",\n         \"geopath\": {\n            \"geopoints\": [\n               [\n                  0.6467556,\n                  0.28522083,\n                  0.46987006,\n                  0.22842954\n               ],\n               [\n                  0.935357,\n                  0.68563277\n               ]\n            ],\n            \"id\": \"59b287534defd02dace3671b\",\n            \"title\": \"Dolorum qui.\",\n            \"type\": \"Ea enim sunt animi et.\"\n         },\n         \"id\": \"59b287534defd02dace3671b\",\n         \"location\": \"Placeat maxime sunt perferendis vel sunt.\",\n         \"path\": \"Ut voluptas.\",\n         \"properties\": {\n            \"Dolores voluptas sunt et quia voluptatem.\": \"Dolores sed.\",\n            \"Officiis voluptatem corporis voluptatem accusamus dolores.\": \"Saepe molestias in.\",\n            \"Ut aut dolore qui iste asperiores hic.\": \"Mollitia nihil cum et voluptate et suscipit.\"\n         },\n         \"refs\": [\n            \"Aliquid cupiditate.\",\n            \"Nihil sint voluptate incidunt excepturi totam.\"\n         ],\n         \"title\": \"Vitae qui in dolorum officiis laborum dolor.\"\n      }\n   }'")
		}
	}
	var path string
	{
		path = poisUpdatePath
	}
	v := &pois.UpdatePayload{}
	if body.Poi != nil {
		v.Poi = marshalPoiPayloadRequestBodyToPoisPoiPayload(body.Poi)
	}
	v.Path = path

	return v, nil
}

// BuildDeactivatePayload builds the payload for the pois deactivate endpoint
// from CLI flags.
func BuildDeactivatePayload(poisDeactivatePath string) (*pois.DeactivatePayload, error) {
	var path string
	{
		path = poisDeactivatePath
	}
	v := &pois.DeactivatePayload{}
	v.Path = path

	return v, nil
}

// BuildDeletePayload builds the payload for the pois delete endpoint from CLI
// flags.
func BuildDeletePayload(poisDeletePath string) (*pois.DeletePayload, error) {
	var path string
	{
		path = poisDeletePath
	}
	v := &pois.DeletePayload{}
	v.Path = path

	return v, nil
}