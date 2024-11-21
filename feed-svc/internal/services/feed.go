package services

import (
	"context"

	"net/http"

	"github.com/asrma7/playpal/feed-svc/internal/db"
	"github.com/asrma7/playpal/feed-svc/internal/models"
	"github.com/asrma7/playpal/feed-svc/pkg/pb"
)

type Server struct {
	DBHandler db.DB
	pb.UnimplementedFeedServiceServer
}

func (s *Server) CreateFeed(ctx context.Context, req *pb.CreateFeedRequest) (*pb.CreateFeedResponse, error) {
	var feed models.Feed

	feed.Title = req.Title
	feed.Author = req.Author
	feed.DateTime = req.Datetime
	feed.Players = req.Players
	feed.Location = req.Location

	if result := s.DBHandler.DB.Create(&feed); result.Error != nil {
		return &pb.CreateFeedResponse{
			Status: http.StatusConflict,
			Error:  result.Error.Error(),
		}, nil
	}
	return &pb.CreateFeedResponse{
		Status: http.StatusCreated,
		Id:     feed.Id,
	}, nil
}

func (s *Server) FindOneFeed(ctx context.Context, req *pb.FindOneRequest) (*pb.FindOneResponse, error) {
	var feed models.Feed

	if result := s.DBHandler.DB.First(&feed, req.Id); result.Error != nil {
		return &pb.FindOneResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}
	return &pb.FindOneResponse{
		Status: http.StatusOK,
		Feed: &pb.FindOneFeed{
			Id:       feed.Id,
			Title:    feed.Title,
			Author:   feed.Author,
			Datetime: feed.DateTime,
			Players:  feed.Players,
			Location: feed.Location,
		},
	}, nil
}

func (s *Server) FindAllFeeds(ctx context.Context, req *pb.FindAllRequest) (*pb.FindAllResponse, error) {
	var feeds []models.Feed

	rows, err := s.DBHandler.DB.Model(&models.Feed{}).Rows()

	if err != nil {
		return &pb.FindAllResponse{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var feed models.Feed
		err := s.DBHandler.DB.ScanRows(rows, &feed)

		if err != nil {
			return nil, err
		}

		feeds = append(feeds, feed)
	}

	var outFeeds []*pb.FindOneFeed
	for _, feed := range feeds {
		var f pb.FindOneFeed
		f.Id = feed.Id
		f.Title = feed.Title
		f.Author = feed.Author
		f.Datetime = feed.DateTime
		f.Players = feed.Players
		f.Location = feed.Location

		outFeeds = append(outFeeds, &f)
	}

	return &pb.FindAllResponse{
		Status: http.StatusOK,
		Feeds:  outFeeds,
	}, nil
}
