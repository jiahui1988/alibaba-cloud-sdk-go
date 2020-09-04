package cloudcallcenter

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

// DownloadTask invokes the cloudcallcenter.DownloadTask API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/downloadtask.html
func (client *Client) DownloadTask(request *DownloadTaskRequest) (response *DownloadTaskResponse, err error) {
	response = CreateDownloadTaskResponse()
	err = client.DoAction(request, response)
	return
}

// DownloadTaskWithChan invokes the cloudcallcenter.DownloadTask API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/downloadtask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DownloadTaskWithChan(request *DownloadTaskRequest) (<-chan *DownloadTaskResponse, <-chan error) {
	responseChan := make(chan *DownloadTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DownloadTask(request)
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

// DownloadTaskWithCallback invokes the cloudcallcenter.DownloadTask API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/downloadtask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DownloadTaskWithCallback(request *DownloadTaskRequest, callback func(response *DownloadTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DownloadTaskResponse
		var err error
		defer close(result)
		response, err = client.DownloadTask(request)
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

// DownloadTaskRequest is the request struct for api DownloadTask
type DownloadTaskRequest struct {
	*requests.RpcRequest
	InstanceId string           `position:"Query" name:"InstanceId"`
	Channel    string           `position:"Query" name:"Channel"`
	TaskId     requests.Integer `position:"Query" name:"TaskId"`
}

// DownloadTaskResponse is the response struct for api DownloadTask
type DownloadTaskResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Data           Data   `json:"Data" xml:"Data"`
}

// CreateDownloadTaskRequest creates a request to invoke DownloadTask API
func CreateDownloadTaskRequest() (request *DownloadTaskRequest) {
	request = &DownloadTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "DownloadTask", "", "")
	request.Method = requests.POST
	return
}

// CreateDownloadTaskResponse creates a response to parse from DownloadTask response
func CreateDownloadTaskResponse() (response *DownloadTaskResponse) {
	response = &DownloadTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}