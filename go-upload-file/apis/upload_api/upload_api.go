package upload_api

import (
	"fmt"
	"net/http"
)

func UploadFile(response http.ResponseWriter, request *http.Request) {

	// limit file 10 MB
	request.ParseMultipartForm(10 * 1024 * 1024)

	file, handler, err := request.FormFile("myFile")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	// get key from form body
	testing := request.Form.Get("testing")
	userName := request.Form.Get("userName")

	fmt.Println("file info")
	fmt.Println("file name: ", handler.Filename)
	fmt.Println("file size: ", handler.Size)
	fmt.Println("file type: ", handler.Header)

	fmt.Println(testing)
	fmt.Println(userName)

}
