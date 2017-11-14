package apiv1

import (
	// _ "cloud.google.com/go"
	"context"
	"fmt"
	"log"

	"github.com/konfortes/firego_poc/models"

	"github.com/fatih/structs"
	"google.golang.org/api/iterator"

	"cloud.google.com/go/firestore"
)

// SuppliersController ...
type SuppliersController struct {
	BaseAPIController
	FirestoreClient  *firestore.Client
	FirestoreContext context.Context
}

type hash map[string]interface{}

// Prepare ...
func (c *SuppliersController) Prepare() {
	c.initFirebaseClient()
}

// Finish ...
func (c *SuppliersController) Finish() {
	// TODO: is there a way to defer from Prepare() ?
	c.FirestoreClient.Close()
}

// GetSupplier ...
// @router /suppliers/:supplierID [GET]
func (c *SuppliersController) GetSupplier(supplierID string) {
	suppliers := c.FirestoreClient.Collection("suppliers").Where("supplier_id", "==", supplierID)
	docs := suppliers.Documents(c.FirestoreContext)
	doc, _ := docs.Next()

	s := success{Data: hash{"supplier": doc.Data()}}
	c.Data["json"] = &s
	c.ServeJSON()
}

// GetSuppliers ...
// @router /suppliers [GET]
func (c *SuppliersController) GetSuppliers() {

	var suppliers = hash{}

	iter := c.FirestoreClient.Collection("suppliers").Documents(c.FirestoreContext)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		suppliers[doc.Ref.ID] = doc.Data()
	}

	s := success{Data: hash{"suppliers": suppliers}}
	c.Data["json"] = &s
	c.ServeJSON()
}

// Create ...
// @router /suppliers [POST]
func (c *SuppliersController) Create() {

	ref, _, err := c.FirestoreClient.Collection("suppliers").Add(c.FirestoreContext, hash{
		"supplier_id": c.getStringParam("supplier_id"),
		"name":        c.getStringParam("name"),
		"actions":     []hash{{"type": "pickup"}, {"type": "drop_off"}},
	})
	if err != nil {
		panic("create failed")
	}

	s := success{Data: hash{"ID": ref.ID}}
	c.Data["json"] = &s
	c.ServeJSON()
}

// Update ...
// @router /suppliers/:supplierID [PUT]
func (c *SuppliersController) Update(supplierID string) {
	supplier := c.parseSupplier()

	suppliers := c.FirestoreClient.Collection("suppliers").Where("supplier_id", "==", supplierID)
	doc, _ := suppliers.Documents(c.FirestoreContext).Next()

	// .Set overrides empty values
	// _, err := doc.Ref.Set(ctx, supplier)
	_, err := doc.Ref.UpdateMap(c.FirestoreContext, structs.Map(supplier))

	if err != nil {
		fmt.Printf("error: %v\n", err)
		panic("update failed")
	}

	s := success{Data: hash{"success": true}}
	c.Data["json"] = &s
	c.ServeJSON()
}

func (c *SuppliersController) parseSupplier() models.Supplier {
	supplier := models.Supplier{}

	supplier.Name = c.getStringParam("name")

	return supplier
}

func (c *SuppliersController) initFirebaseClient() {
	// export GOOGLE_APPLICATION_CREDENTIALS="/Users/konfortes/workspace/golang/src/github.com/konfortes/firego_poc/keyfile.json"
	ctx := context.Background()
	projectID := "actions-f7aae"
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	c.FirestoreClient = client
	c.FirestoreContext = ctx
}
