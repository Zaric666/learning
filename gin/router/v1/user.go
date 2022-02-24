package v1

import (
	"context"
	"github.com/Zaric666/learning/gin/entity"
	userRpc "github.com/Zaric666/learning/grpc/simple_grpc/user"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"time"
)

const (
	address = "localhost:50051"
)

type User struct{}

func (u User) Index(c *gin.Context) {
	user := entity.User{}
	res := entity.Result{}

	if err := c.ShouldBind(&user); err != nil {
		res.SetCode(entity.CodeError)
		res.SetMessage(err.Error())
		c.JSON(http.StatusForbidden, res)
		c.Abort()
		return
	}

	// rpc
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	userClient := userRpc.NewUserClient(conn)

	// 设定请求超时时间 3s
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	// UserIndex 请求
	userIndexReponse, err := userClient.UserIndex(ctx, &userRpc.UserIndexRequest{Page: 1, PageSize: 12})
	if err != nil {
		log.Printf("user index could not greet: %v", err)
	}

	if 0 == userIndexReponse.Err {
		// 包含 UserEntity 的数组列表
		userEntityList := userIndexReponse.Data
		data := make(map[string]interface{})
		for _, row := range userEntityList {
			data[row.Name] = entity.User{Name: row.Name, Age: int(row.Age)}
		}
		res.SetCode(entity.CodeError)
		res.SetData(data)
		c.JSON(http.StatusOK, res)
	} else {
		log.Printf("user index error: %d", userIndexReponse.Err)
	}
}
