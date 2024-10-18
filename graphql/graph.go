package main

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/SaurabPoudel/go-grpc-graphql-micro/account"
	"github.com/SaurabPoudel/go-grpc-graphql-micro/catalog"
	"github.com/SaurabPoudel/go-grpc-graphql-micro/order"
)

// import "github.com/99designs/gqlgen/graphql"

type Server struct {
	accountClient *account.Client
	catalogClient *catalog.Client
	orderClient   *order.Client
}

func NewGraphQLServer(accountUrl, catalogUrl, orderUrl string) (*Server, error) {
	//
	// 	accountClient, err := account.NewClient(accountUrl)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	//
	// 	catalogClient, err := catalog.NewClient(catalogUrl)
	// 	if err != nil {
	// 		accountClient.Close()
	// 		return nil, err
	// 	}
	//
	// 	orderClient, err := order.NewClient(orderUrl)
	// 	if err != nil {
	// 		accountClient.Close()
	// 		catalogClient.Close()
	// 		return nil, err
	// 	}
	//
	return &Server{
		// 		accountClient,
		// 		catalogClient,
		// 		orderClient,
	}, nil
}

//	func (s *Server) Mutation() MutationResolver {
//		return &mutationResolver{
//			server: s,
//		}
//	}
//
// func (s *Server) Mutation() QueryResolver {
//
//		return &queryResolver{
//			server: s,
//		}
//	}
//
// func (s *Server) Account() AccountResolver {
//
//		return &accountResolver{
//			server: s,
//		}
//	}
func (s *Server) ToExecutableSchema() graphql.ExecutableSchema {
	return NewExecutableSchema(Config{
		Resolvers: s,
	})
}
