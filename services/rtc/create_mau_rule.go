package rtc

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

// CreateMAURule invokes the rtc.CreateMAURule API synchronously
// api document: https://help.aliyun.com/api/rtc/createmaurule.html
func (client *Client) CreateMAURule(request *CreateMAURuleRequest) (response *CreateMAURuleResponse, err error) {
	response = CreateCreateMAURuleResponse()
	err = client.DoAction(request, response)
	return
}

// CreateMAURuleWithChan invokes the rtc.CreateMAURule API asynchronously
// api document: https://help.aliyun.com/api/rtc/createmaurule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateMAURuleWithChan(request *CreateMAURuleRequest) (<-chan *CreateMAURuleResponse, <-chan error) {
	responseChan := make(chan *CreateMAURuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateMAURule(request)
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

// CreateMAURuleWithCallback invokes the rtc.CreateMAURule API asynchronously
// api document: https://help.aliyun.com/api/rtc/createmaurule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateMAURuleWithCallback(request *CreateMAURuleRequest, callback func(response *CreateMAURuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateMAURuleResponse
		var err error
		defer close(result)
		response, err = client.CreateMAURule(request)
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

// CreateMAURuleRequest is the request struct for api CreateMAURule
type CreateMAURuleRequest struct {
	*requests.RpcRequest
	UseridPrefix  string           `position:"Query" name:"UseridPrefix"`
	ChannelPrefix string           `position:"Query" name:"ChannelPrefix"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	MauTemplateId requests.Integer `position:"Query" name:"MauTemplateId"`
	AppId         string           `position:"Query" name:"AppId"`
	CallBack      string           `position:"Query" name:"CallBack"`
}

// CreateMAURuleResponse is the response struct for api CreateMAURule
type CreateMAURuleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	RuleId    int64  `json:"RuleId" xml:"RuleId"`
}

// CreateCreateMAURuleRequest creates a request to invoke CreateMAURule API
func CreateCreateMAURuleRequest() (request *CreateMAURuleRequest) {
	request = &CreateMAURuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("rtc", "2018-01-11", "CreateMAURule", "", "")
	return
}

// CreateCreateMAURuleResponse creates a response to parse from CreateMAURule response
func CreateCreateMAURuleResponse() (response *CreateMAURuleResponse) {
	response = &CreateMAURuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
