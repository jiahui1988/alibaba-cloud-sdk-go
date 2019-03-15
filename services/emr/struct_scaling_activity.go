package emr

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

// ScalingActivity is a nested struct in emr response
type ScalingActivity struct {
	BizId         string `json:"BizId" xml:"BizId"`
	StartTime     int    `json:"StartTime" xml:"StartTime"`
	EndTime       int    `json:"EndTime" xml:"EndTime"`
	InstanceIds   string `json:"InstanceIds" xml:"InstanceIds"`
	TotalCapacity int    `json:"TotalCapacity" xml:"TotalCapacity"`
	Cause         string `json:"Cause" xml:"Cause"`
	Description   string `json:"Description" xml:"Description"`
	Status        string `json:"Status" xml:"Status"`
	Transition    string `json:"Transition" xml:"Transition"`
	ScalingRuleId string `json:"ScalingRuleId" xml:"ScalingRuleId"`
	ExpectNum     int    `json:"ExpectNum" xml:"ExpectNum"`
}