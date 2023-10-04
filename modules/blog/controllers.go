package blog

import (
	"context"
	"time"

	"github.com/Jeffail/gabs"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"main.go/config"
)

func create(c *fiber.Ctx) error {
	apiBody, err := gabs.ParseJSON([]byte(c.Body()))
	if err != nil {
		return err
	}
	blog := new(Model)
	if apiBody.Path("title").Data() != nil {
		blog.Title = apiBody.Path("title").Data().(string)
	}
	if apiBody.Path("description.short").Data() != nil {
		blog.Description.Short = apiBody.Path("description.short").Data().(string)
	}
	if apiBody.Path("description.long").Data() != nil {
		blog.Description.Long = apiBody.Path("description.long").Data().(string)
	}
	if apiBody.Path("category").Data() != nil {
		blog.Category = apiBody.Path("category").Data().(string)
	}
	if apiBody.Path("created.id").Data() != nil {
		blog.Created.Time = time.Now()
		id := apiBody.Path("created.id").Data().(string)
		blog.Created.ID, err = primitive.ObjectIDFromHex(id)
		if err != nil {
			return err
		}
	}
	if apiBody.Path("monetized.ok").Data() != nil {
		blog.Monetized.Ok = apiBody.Path("monetized.ok").Data().(bool)
	}
	ID, err := config.MI.DB.Collection("blogs").InsertOne(context.Background(), blog)
	if err != nil {
		return c.Status(200).JSON(&fiber.Map{
			"success": false,
			"message": "Error in creating a blog",
			"error":   err,
		})
	}
	blog.ID = ID.InsertedID.(primitive.ObjectID)
	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "Blog created successfully.",
	})
}

func delete(c *fiber.Ctx) error {
	apiBody, err := gabs.ParseJSON([]byte(c.Body()))
	if err != nil {
		return err
	}
	var blogID primitive.ObjectID
	if apiBody.Path("blog_id").Data() == nil {
		return c.Status(200).JSON(&fiber.Map{
			"success": false,
			"message": "Please enter the blog id to delete.",
		})
	} else {
		id := apiBody.Path("blog_id").Data().(string)
		blogID, err = primitive.ObjectIDFromHex(id)
		if err != nil {
			return c.Status(200).JSON(&fiber.Map{
				"success": false,
				"message": "Invalid blog id.",
			})
		}
	}
	filter := bson.M{"_id": blogID}
	_, err = config.MI.DB.Collection("blogs").DeleteOne(context.TODO(), filter)
	if err != nil {
		return c.Status(200).JSON(&fiber.Map{
			"success": false,
			"message": "Error in deleting a blog.",
			"error":   err,
		})
	}
	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "Blog deleted successfully.",
	})
}

func fetchOne(c *fiber.Ctx) error {
	apiBody, err := gabs.ParseJSON([]byte(c.Body()))
	if err != nil {
		return err
	}
	var blogID primitive.ObjectID
	if apiBody.Path("blog_id").Data() == nil {
		return c.Status(200).JSON(&fiber.Map{
			"success": false,
			"message": "Please enter the blog id to delete.",
		})
	} else {
		id := apiBody.Path("blog_id").Data().(string)
		blogID, err = primitive.ObjectIDFromHex(id)
		if err != nil {
			return c.Status(200).JSON(&fiber.Map{
				"success": false,
				"message": "Invalid blog id.",
			})
		}
	}
	blog := Model{}
	filter := bson.M{"_id": blogID}
	err = config.MI.DB.Collection("blogs").FindOne(context.TODO(), filter).Decode(&blog)
	if err != nil {
		return c.Status(200).JSON(&fiber.Map{
			"success": false,
			"message": "Error in finding a blog.",
			"error":   err,
		})
	}
	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "Blog fetched successfully.",
		"blog":    blog,
	})
}

func fetchAll(c *fiber.Ctx) error {
	blogs := []Model{}
	filter := bson.M{}
	opts := options.Find().SetSort(bson.M{"_id": 1})
	cursor, err := config.MI.DB.Collection("blogs").Find(context.TODO(), filter, opts)
	if err != nil {
		return c.Status(200).JSON(&fiber.Map{
			"success": false,
			"message": "Error in finding blogs.",
			"error":   err,
		})
	}
	err = cursor.All(context.Background(), &blogs)
	if err != nil {
		return c.Status(200).JSON(&fiber.Map{
			"success": false,
			"message": "Error in decoding blogs.",
			"error":   err,
		})
	}
	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "Blogs fetched successfully.",
		"blogs":   blogs,
		"count":   len(blogs),
	})
}

func update(c *fiber.Ctx) error {
	apiBody, err := gabs.ParseJSON([]byte(c.Body()))
	if err != nil {
		return err
	}
	set := bson.M{}
	var blogID primitive.ObjectID
	if apiBody.Path("blog_id").Data() == nil {
		return c.Status(200).JSON(&fiber.Map{
			"success": false,
			"message": "Please enter the blog id to update.",
		})
	} else {
		id := apiBody.Path("blog_id").Data().(string)
		blogID, err = primitive.ObjectIDFromHex(id)
		if err != nil {
			return c.Status(200).JSON(&fiber.Map{
				"success": false,
				"message": "Invalid blog id.",
			})
		}
	}
	if apiBody.Path("title").Data() != nil {
		set["title"] = apiBody.Path("title").Data().(string)
	}
	if apiBody.Path("description.short").Data() != nil {
		set["description.short"] = apiBody.Path("description.short").Data().(string)
	}
	if apiBody.Path("description.long").Data() != nil {
		set["description.short"] = apiBody.Path("description.long").Data().(string)
	}
	if apiBody.Path("category").Data() != nil {
		set["category"] = apiBody.Path("category").Data().(string)
	}
	if apiBody.Path("monetized.ok").Data() != nil {
		set["monetized"] = apiBody.Path("monetized.ok").Data().(bool)
	}
	set["updated_at"] = time.Now()
	update := bson.M{"$set": set}
	after := options.After
	upsert := true
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	filter := bson.M{"_id": blogID}
	var updatedData map[string]interface{}
	err = config.MI.DB.Collection("blogs").FindOneAndUpdate(context.TODO(), filter, update, &opt).Decode(&updatedData)
	if err != nil {
		return c.Status(200).JSON(&fiber.Map{
			"success": false,
			"message": "Error in creating a blog.",
			"error":   err,
		})
	}
	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "Blog updated successfully.",
		"data":    updatedData,
	})
}
