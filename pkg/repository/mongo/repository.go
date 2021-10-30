package mongo

import (
	"context"
	"golang-interview-project-masaru-ohashi/pkg/member"
	m "golang-interview-project-masaru-ohashi/pkg/member"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
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

func NewMongoRepository(mongoURL, mongoDB string, mongoTimeout int) (member.MemberRepository, error) {
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

func (r *mongoRepository) Find(name string) (*member.Member, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	member := &member.Member{}
	collection := r.client.Database(r.database).Collection("members")
	filter := bson.M{"name": name}
	err := collection.FindOne(ctx, filter).Decode(&member)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.Wrap(m.ErrMemberNotFound, "repository.Member.Find")
		}
		return nil, errors.Wrap(err, "repository.Redirect.Find")
	}
	return member, nil
}

func (r *mongoRepository) Store(member *member.Member) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection("members")
	_, err := collection.InsertOne(
		ctx,
		bson.M{
			"code":       member.Code,
			"name":       member.Name,
			"tags":       member.Tags,
			"created_at": member.CreatedAt,
		},
	)
	if err != nil {
		return errors.Wrap(err, "repository.Member.Store")
	}
	return nil
}
