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

// TaskProgressInfo is a nested struct in polardbx response
type TaskProgressInfo struct {
	DBName           string `json:"DBName" xml:"DBName"`
	BeginTime        string `json:"BeginTime" xml:"BeginTime"`
	ProgressInfo     string `json:"ProgressInfo" xml:"ProgressInfo"`
	FinishTime       string `json:"FinishTime" xml:"FinishTime"`
	TaskAction       string `json:"TaskAction" xml:"TaskAction"`
	TaskId           string `json:"TaskId" xml:"TaskId"`
	Progress         string `json:"Progress" xml:"Progress"`
	Status           string `json:"Status" xml:"Status"`
	TaskErrorCode    string `json:"TaskErrorCode" xml:"TaskErrorCode"`
	TaskErrorMessage string `json:"TaskErrorMessage" xml:"TaskErrorMessage"`
	ScaleOutToken    string `json:"ScaleOutToken" xml:"ScaleOutToken"`
}
