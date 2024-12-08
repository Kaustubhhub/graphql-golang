package database

import (
	"context"
	"time"

	"github.com/kaustubhhub/graphql-project/graph/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	Client *mongo.Client
}

var connectionString string = "mongodb://localhost:27017"

func Connect() *DB {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil
	}

	return &DB{Client: client}
}

func (db *DB) GetJob(id string) *model.JobListing {
	var jobListing model.JobListing
	return &jobListing
}

func (db *DB) GetJobs() []*model.JobListing {
	var jobListings []*model.JobListing
	return jobListings
}

func (db *DB) CreateJobListing(createInfo model.CreateJobListingInput) *model.JobListing {
	var returnJobListing model.JobListing
	return &returnJobListing
}

func (db *DB) UpdateJobListing(jobId string, jobInfo model.UpdateJobListingInput) *model.JobListing {
	var jobListing model.JobListing
	return &jobListing
}

func (db *DB) DeleteJobListing(jobId string) *model.DeleteJobResponse {
	return &model.DeleteJobResponse{DeleteJobID: jobId}
}
