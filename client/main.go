package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"

	pb "golang.cafe/protobuf/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed server's certificate
	pemServerCA, err := ioutil.ReadFile("../ca/ca-cert.pem")
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}

	// Create the credentials and return it
	config := &tls.Config{
		RootCAs: certPool,
	}

	return credentials.NewTLS(config), nil
}

func main() {
	var conn *grpc.ClientConn

	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}

	// grpc.Dial(*serverAddress, grpc.WithTransportCredentials(tlsCredentials))
	conn, err = grpc.Dial("0.0.0.0:9000", grpc.WithTransportCredentials(tlsCredentials))

	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer conn.Close()

	c := pb.NewRouteGuideClient(conn)

	point := pb.Point{Latitude: 100, Longitude: 200}

	response, err := c.GetFeature(context.Background(), &point)

	if err != nil {
		log.Fatalf("Error when calling GetFeature: %v", err)
	}

	log.Printf("Response from Server: %v", response)
}
