package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models/schema"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("xuiagwc0s3jkqd4")
		if err != nil {
			return err
		}

		// add
		new_phone := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "wkbd337c",
			"name": "phone",
			"type": "text",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), new_phone); err != nil {
			return err
		}
		collection.Schema.AddField(new_phone)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("xuiagwc0s3jkqd4")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("wkbd337c")

		return dao.SaveCollection(collection)
	})
}
