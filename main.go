package main

import (
	"fmt"
	"log"
	"os"

	proto "github.com/golang/protobuf/proto"

	"github.com/confus1on/go-protobuff/pb"
)

func main() {
	fmt.Println("Belajar Protocol Buffers : ")

	// set data `tasg` protobuff
	tags := []*pb.Article_Tags{
		&pb.Article_Tags{
			TagId: 1,
			Name:  "Golang",
		},
		&pb.Article_Tags{
			TagId: 2,
			Name:  "PHP",
		},
		&pb.Article_Tags{
			TagId: 3,
			Name:  "Javascript",
		},
	}

	socmedStats := make(map[string]*pb.Article_SocialMediaStatisticField)
	socmedStats["facebook"] = &pb.Article_SocialMediaStatisticField{
		Like:     10,
		Share:    20,
		Comments: 5,
	}
	socmedStats["twitter"] = &pb.Article_SocialMediaStatisticField{
		Like:     20,
		Share:    30,
		Comments: 40,
	}

	// set data `oneof` protobuff
	process := &pb.Article_Update{
		Update: true,
	}

	article := &pb.Article{
		Id:                   1,
		Title:                "Belajar Membuat Protobuff",
		Content:              "Lorem Ipsum Dolor Amet",
		Status:               pb.Article_DRAFT,
		Tags:                 tags,
		SocialMediaStatistic: socmedStats,
		ProcessOneof:         process,
	}

	data, err := proto.Marshal(article)
	if err != nil {
		log.Fatalf("Marshalling Error : %v", err.Error())
		os.Exit(1)
	}

	fmt.Printf("%+v \n", string(data))
	fmt.Println(len(data))
	fmt.Println("*=================== Unmarshalling =====================*")

	unmarshArticle := &pb.Article{}
	err = proto.Unmarshal(data, unmarshArticle)
	if err != nil {
		log.Fatalf("Error Unmarshalling : %+v ", err.Error())
	}

	fmt.Println(unmarshArticle)

}
