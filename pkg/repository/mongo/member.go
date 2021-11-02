package mongo

import (
	"context"
	"golang-interview-project-masaru-ohashi/cmd/common"
	"golang-interview-project-masaru-ohashi/pkg/team"
	"reflect"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type mongoRepository struct {
	client   *mongo.Client
	database string
	timeout  time.Duration
}

func newMongoClient(mongoURL string, mongoTimeout int) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(mongoTimeout)*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	return client, nil
}

func NewMongoRepository(mongoURL, mongoDB string, mongoTimeout int) (team.MemberRepository, error) {
	repo := &mongoRepository{
		timeout:  time.Duration(mongoTimeout) * time.Second,
		database: mongoDB,
	}
	client, err := newMongoClient(mongoURL, mongoTimeout)
	if err != nil {
		return nil, errors.Wrap(err, "repository.NewMongoRepo")
	}
	repo.client = client
	return repo, nil
}

func (r *mongoRepository) DbGetAll() (members []interface{}, err error) {
	collection := r.client.Database(r.database).Collection("members")
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		return members, err
	}
	for cursor.Next(context.Background()) {
		resultMap := map[string]interface{}{}
		err := cursor.Decode(&resultMap)
		if err != nil {
			return members, err
		}
		agreement := resultMap["agreement"].(string)
		if agreement == common.CONTRACTOR {
			var tags []string
			tgs := resultMap["tags"]
			s := reflect.ValueOf(tgs)
			for i := 0; i < s.Len(); i++ {
				tags = append(tags, s.Index(i).Elem().String())
			}
			member := team.Contractor{
				Colaborator: team.Colaborator{
					ID:        resultMap["_id"].(primitive.ObjectID),
					Name:      resultMap["name"].(string),
					Agreement: resultMap["agreement"].(string),
					CreatedAt: resultMap["created_at"].(int64),
					Tags:      tags,
				},
				Duration: int(resultMap["duration"].(int32)),
			}
			members = append(members, member)
		} else if agreement == string(common.EMPLOYEE) {
			var tags []string
			tgs := resultMap["tags"]
			s := reflect.ValueOf(tgs)
			for i := 0; i < s.Len(); i++ {
				tags = append(tags, s.Index(i).Elem().String())
			}
			member := team.Employee{
				Colaborator: team.Colaborator{
					Name:      resultMap["name"].(string),
					Agreement: resultMap["agreement"].(string),
					CreatedAt: resultMap["created_at"].(int64),
					Tags:      tags,
				},
				Role: resultMap["role"].(string),
			}
			members = append(members, member)
		}
	}
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.Wrap(team.ErrMemberNotFound, "repository.Member.GetAll")
		}
		return nil, errors.Wrap(err, "repository.Member.GetAll")
	}
	return members, nil
}

func (r *mongoRepository) DbGet(name string) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection("members")
	filter := bson.M{"name": name}
	resultMap := map[string]interface{}{}
	err := collection.FindOne(ctx, filter).Decode(&resultMap)
	agreement := resultMap["agreement"].(string)
	var member team.Member
	if agreement == common.CONTRACTOR {
		var tags []string
		tgs := resultMap["tags"]
		s := reflect.ValueOf(tgs)
		for i := 0; i < s.Len(); i++ {
			tags = append(tags, s.Index(i).Elem().String())
		}
		member = &team.Contractor{
			Colaborator: team.Colaborator{
				ID:        resultMap["_id"].(primitive.ObjectID),
				Name:      resultMap["name"].(string),
				Agreement: resultMap["agreement"].(string),
				CreatedAt: resultMap["created_at"].(int64),
				Tags:      tags,
			},
			Duration: int(resultMap["duration"].(int32)),
		}
	} else if agreement == string(common.EMPLOYEE) {
		var tags []string
		tgs := resultMap["tags"]
		s := reflect.ValueOf(tgs)
		for i := 0; i < s.Len(); i++ {
			tags = append(tags, s.Index(i).Elem().String())
		}
		member = &team.Employee{
			Colaborator: team.Colaborator{
				ID:        resultMap["_id"].(primitive.ObjectID),
				Name:      resultMap["name"].(string),
				Agreement: resultMap["agreement"].(string),
				CreatedAt: resultMap["created_at"].(int64),
				Tags:      tags,
			},
			Role: resultMap["role"].(string),
		}
	}
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.Wrap(team.ErrMemberNotFound, "repository.Member.Get")
		}
		return nil, errors.Wrap(err, "repository.Member.Get")
	}
	return member, nil
}

func (r *mongoRepository) DbCreate(member interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	collection := r.client.Database(r.database).Collection("members")
	collection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{{Key: "name", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	)
	switch v := member.(type) {
	case *team.Contractor:
		v = member.(*team.Contractor)
		_, err := collection.InsertOne(
			ctx,
			bson.M{
				"name":       v.GetName(),
				"duration":   v.Duration,
				"agreement":  v.GetAgreement(),
				"tags":       v.GetTags(),
				"created_at": v.GetCreatedAt(),
			},
		)
		if err != nil {
			return errors.Wrap(err, "repository.Contractor.Create")
		}
	case *team.Employee:
		v = member.(*team.Employee)
		_, err := collection.InsertOne(
			ctx,
			bson.M{
				"name":       v.GetName(),
				"agreement":  v.GetAgreement(),
				"role":       v.Role,
				"tags":       v.GetTags(),
				"created_at": v.GetCreatedAt(),
			},
		)
		if err != nil {
			return errors.Wrap(err, "repository.Employee.Create")
		}
	}
	return nil
}

func (r *mongoRepository) DbUpdate(member interface{}) error {
	switch member.(type) {
	case *team.Contractor:
	case *team.Employee:
		println(member.(*team.Employee).Colaborator.ID.Hex())
		collection := r.client.Database(r.database).Collection("members")
		filter := bson.M{"_id": bson.M{"$eq": "6181232b0f190acdebe3a92e"}}
		update := bson.M{"$set": bson.M{"role": member.(*team.Employee).Role}}
		collection.UpdateOne(
			context.Background(),
			filter,
			update,
		)
	}
	return nil
}

func (r *mongoRepository) DbDelete(member interface{}) error {
	return nil
}
