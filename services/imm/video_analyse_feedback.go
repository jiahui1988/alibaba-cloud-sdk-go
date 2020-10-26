package imm

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// VideoAnalyseFeedback invokes the imm.VideoAnalyseFeedback API synchronously
func (client *Client) VideoAnalyseFeedback(request *VideoAnalyseFeedbackRequest) (response *VideoAnalyseFeedbackResponse, err error) {
	response = CreateVideoAnalyseFeedbackResponse()
	err = client.DoAction(request, response)
	return
}

// VideoAnalyseFeedbackWithChan invokes the imm.VideoAnalyseFeedback API asynchronously
func (client *Client) VideoAnalyseFeedbackWithChan(request *VideoAnalyseFeedbackRequest) (<-chan *VideoAnalyseFeedbackResponse, <-chan error) {
	responseChan := make(chan *VideoAnalyseFeedbackResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.VideoAnalyseFeedback(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// VideoAnalyseFeedbackWithCallback invokes the imm.VideoAnalyseFeedback API asynchronously
func (client *Client) VideoAnalyseFeedbackWithCallback(request *VideoAnalyseFeedbackRequest, callback func(response *VideoAnalyseFeedbackResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *VideoAnalyseFeedbackResponse
		var err error
		defer close(result)
		response, err = client.VideoAnalyseFeedback(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// VideoAnalyseFeedbackRequest is the request struct for api VideoAnalyseFeedback
type VideoAnalyseFeedbackRequest struct {
	*requests.RpcRequest
	Note       string `position:"Query" name:"Note"`
	Project    string `position:"Query" name:"Project"`
	TaskId     string `position:"Query" name:"TaskId"`
	Frames     string `position:"Query" name:"Frames"`
	Suggestion string `position:"Query" name:"Suggestion"`
	Uri        string `position:"Query" name:"Uri"`
	Scenes     string `position:"Query" name:"Scenes"`
}

// VideoAnalyseFeedbackResponse is the response struct for api VideoAnalyseFeedback
type VideoAnalyseFeedbackResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateVideoAnalyseFeedbackRequest creates a request to invoke VideoAnalyseFeedback API
func CreateVideoAnalyseFeedbackRequest() (request *VideoAnalyseFeedbackRequest) {
	request = &VideoAnalyseFeedbackRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "VideoAnalyseFeedback", "imm", "openAPI")
	request.Method = requests.POST
	return
}

// CreateVideoAnalyseFeedbackResponse creates a response to parse from VideoAnalyseFeedback response
func CreateVideoAnalyseFeedbackResponse() (response *VideoAnalyseFeedbackResponse) {
	response = &VideoAnalyseFeedbackResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}