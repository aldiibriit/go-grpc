package main

import (
	"context"
	"database/sql"
	"fmt"
	"go-grpc-2/database"
	pb "go-grpc-2/proto/echo"
	vulcan "go-grpc-2/proto/vulcan"
	"log"
	"net"
	"net/http"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedEchoServiceServer
	vulcan.UnimplementedMobileAppsServer
}

type MobileApp struct {
	// result     map[string]string `json:"result"`
	APP_ID    string `json:"APP_ID"`
	Title     string `json:"title"`
	TableName string `json:"table_name"`
	Result    string `json:"result"`
}

func (s *server) UnaryEcho(ctx context.Context, in *pb.EchoRequest) (*pb.EchoResponse, error) {
	fmt.Println("Unary Echo")
	fmt.Println("Incoming Request")
	fmt.Printf("Message : %s -> send echo", in.Message)
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.DataLoss, "error unary echo")
	}

	if v, ok := md["timestamps"]; ok {
		fmt.Println("metadata timestamps")
		for i, e := range v {
			fmt.Printf("%d,%s", i, e)
		}
	}

	return &pb.EchoResponse{Message: in.Message}, nil
}

func (s *server) SelectMobileAppByAppID(ctx context.Context, in *vulcan.RequestSelectMobileAppByAppID) (*vulcan.ResponseSelectMobileAppByAppID, error) {
	dbtype, dbname := database.Connect("peduliATM")
	db, err := sql.Open(dbtype, dbname)

	if err != nil || db == nil {
		log.Fatal(http.StatusInternalServerError, "Could not open DB")
	}

	defer db.Close()

	mobileApp := MobileApp{}
	query := fmt.Sprintf(`SELECT APP_ID,title,table_name FROM mobile_apps WHERE APP_ID = "%s"`, in.APP_ID)

	rows, err := db.Query(query)

	if err != nil {
		log.Fatal(http.StatusInternalServerError, err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		// var SIK_ECH_TAHAP sql.NullString
		err := rows.Scan(&mobileApp.APP_ID, &mobileApp.Title, &mobileApp.TableName)
		// tid.SIK_ECH_TAHAP = SIK_ECH_TAHAP.String

		if err != nil {
			log.Fatal(http.StatusInternalServerError, err.Error())
		}
	}

	mapMobileApp := map[string]string{
		"app_id":     mobileApp.APP_ID,
		"table_name": mobileApp.TableName,
		"title":      mobileApp.Title,
	}

	fmt.Println(mapMobileApp)
	test := &vulcan.ResponseSelectMobileAppByAppID{
		Result: mapMobileApp,
	}
	fmt.Println(test)

	return &vulcan.ResponseSelectMobileAppByAppID{
		APP_ID:    mobileApp.APP_ID,
		TableName: mobileApp.TableName,
		Title:     mobileApp.Title,
		Result:    mapMobileApp,
		Status:    "200",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":7177")
	if err != nil {
		fmt.Printf("error : %v", err.Error())
		os.Exit(1)
	}

	s := grpc.NewServer()
	pb.RegisterEchoServiceServer(s, &server{})
	// pb.RegisterEchoServiceServer(s, &server{})
	vulcan.RegisterMobileAppsServer(s, &server{})
	s.Serve(lis)
}
