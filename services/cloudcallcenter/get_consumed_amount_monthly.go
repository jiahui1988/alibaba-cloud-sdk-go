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

// GetConsumedAmountMonthly invokes the cloudcallcenter.GetConsumedAmountMonthly API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getconsumedamountmonthly.html
func (client *Client) GetConsumedAmountMonthly(request *GetConsumedAmountMonthlyRequest) (response *GetConsumedAmountMonthlyResponse, err error) {
	response = CreateGetConsumedAmountMonthlyResponse()
	err = client.DoAction(request, response)
	return
}

// GetConsumedAmountMonthlyWithChan invokes the cloudcallcenter.GetConsumedAmountMonthly API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getconsumedamountmonthly.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetConsumedAmountMonthlyWithChan(request *GetConsumedAmountMonthlyRequest) (<-chan *GetConsumedAmountMonthlyResponse, <-chan error) {
	responseChan := make(chan *GetConsumedAmountMonthlyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetConsumedAmountMonthly(request)
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

// GetConsumedAmountMonthlyWithCallback invokes the cloudcallcenter.GetConsumedAmountMonthly API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getconsumedamountmonthly.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetConsumedAmountMonthlyWithCallback(request *GetConsumedAmountMonthlyRequest, callback func(response *GetConsumedAmountMonthlyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetConsumedAmountMonthlyResponse
		var err error
		defer close(result)
		response, err = client.GetConsumedAmountMonthly(request)
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

// GetConsumedAmountMonthlyRequest is the request struct for api GetConsumedAmountMonthly
type GetConsumedAmountMonthlyRequest struct {
	*requests.RpcRequest
	Phone      string           `position:"Query" name:"phone"`
	PageSize   requests.Integer `position:"Query" name:"pageSize"`
	EndTime    requests.Integer `position:"Query" name:"EndTime"`
	StartTime  requests.Integer `position:"Query" name:"StartTime"`
	PageNumber requests.Integer `position:"Query" name:"pageNumber"`
}

// GetConsumedAmountMonthlyResponse is the response struct for api GetConsumedAmountMonthly
type GetConsumedAmountMonthlyResponse struct {
	*responses.BaseResponse
	RequestId      string       `json:"RequestId" xml:"RequestId"`
	Success        bool         `json:"Success" xml:"Success"`
	Code           string       `json:"Code" xml:"Code"`
	Message        string       `json:"Message" xml:"Message"`
	HttpStatusCode int          `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Consumptions   Consumptions `json:"Consumptions" xml:"Consumptions"`
}

// CreateGetConsumedAmountMonthlyRequest creates a request to invoke GetConsumedAmountMonthly API
func CreateGetConsumedAmountMonthlyRequest() (request *GetConsumedAmountMonthlyRequest) {
	request = &GetConsumedAmountMonthlyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "GetConsumedAmountMonthly", "", "")
	request.Method = requests.POST
	return
}

// CreateGetConsumedAmountMonthlyResponse creates a response to parse from GetConsumedAmountMonthly response
func CreateGetConsumedAmountMonthlyResponse() (response *GetConsumedAmountMonthlyResponse) {
	response = &GetConsumedAmountMonthlyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}