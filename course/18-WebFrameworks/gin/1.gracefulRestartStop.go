package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*

example from https://gin-gonic.com/docs/examples/graceful-restart-or-stop/

*/
func main(){
router := gin.Default()
router.GET("/", func(c  *gin.Context){
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
})
// gin.ListenAndServe(":8080", router)
//router.Run(":8081")
//http.ListenAndServe(":8080", router)

server := &http.Server{
	Addr: ":8080",
	Handler: router,
}

go func(){
	if err := server.ListenAndServe(); err != nil{
		//log.Fatalf("error:%s",err)
		log.Panic(err) // this logs and calls panic, which will looke for defer statements with recover if any otherwise exits
		//panic(err)
	}
}()



}