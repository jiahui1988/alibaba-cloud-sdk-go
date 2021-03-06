package polardbx

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

// CheckHealth invokes the polardbx.CheckHealth API synchronously
func (client *Client) CheckHealth(request *CheckHealthRequest) (response *CheckHealthResponse, err error) {
	response = CreateCheckHealthResponse()
	err = client.DoAction(request, response)
	return
}

// CheckHealthWithChan invokes the polardbx.CheckHealth API asynchronously
func (client *Client) CheckHealthWithChan(request *CheckHealthRequest) (<-chan *CheckHealthResponse, <-chan error) {
	responseChan := make(chan *CheckHealthResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CheckHealth(request)
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

// CheckHealthWithCallback invokes the polardbx.CheckHealth API asynchronously
func (client *Client) CheckHealthWithCallback(request *CheckHealthRequest, callback func(response *CheckHealthResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CheckHealthResponse
		var err error
		defer close(result)
		response, err = client.CheckHealth(request)
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

// CheckHealthRequest is the request struct for api CheckHealth
type CheckHealthRequest struct {
	*requests.RpcRequest
}

// CheckHealthResponse is the response struct for api CheckHealth
type CheckHealthResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Status    string `json:"Status" xml:"Status"`
}

// CreateCheckHealthRequest creates a request to invoke CheckHealth API
func CreateCheckHealthRequest() (request *CheckHealthRequest) {
	request = &CheckHealthRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardbx", "2020-02-02", "CheckHealth", "polardbx", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCheckHealthResponse creates a response to parse from CheckHealth response
func CreateCheckHealthResponse() (response *CheckHealthResponse) {
	response = &CheckHealthResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
