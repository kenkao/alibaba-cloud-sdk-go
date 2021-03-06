package outboundbot

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

// ModifyJobGroup invokes the outboundbot.ModifyJobGroup API synchronously
// api document: https://help.aliyun.com/api/outboundbot/modifyjobgroup.html
func (client *Client) ModifyJobGroup(request *ModifyJobGroupRequest) (response *ModifyJobGroupResponse, err error) {
	response = CreateModifyJobGroupResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyJobGroupWithChan invokes the outboundbot.ModifyJobGroup API asynchronously
// api document: https://help.aliyun.com/api/outboundbot/modifyjobgroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyJobGroupWithChan(request *ModifyJobGroupRequest) (<-chan *ModifyJobGroupResponse, <-chan error) {
	responseChan := make(chan *ModifyJobGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyJobGroup(request)
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

// ModifyJobGroupWithCallback invokes the outboundbot.ModifyJobGroup API asynchronously
// api document: https://help.aliyun.com/api/outboundbot/modifyjobgroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyJobGroupWithCallback(request *ModifyJobGroupRequest, callback func(response *ModifyJobGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyJobGroupResponse
		var err error
		defer close(result)
		response, err = client.ModifyJobGroup(request)
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

// ModifyJobGroupRequest is the request struct for api ModifyJobGroup
type ModifyJobGroupRequest struct {
	*requests.RpcRequest
	Description   string    `position:"Query" name:"Description"`
	CallingNumber *[]string `position:"Query" name:"CallingNumber"  type:"Repeated"`
	ScriptId      string    `position:"Query" name:"ScriptId"`
	InstanceId    string    `position:"Query" name:"InstanceId"`
	StrategyJson  string    `position:"Query" name:"StrategyJson"`
	JobGroupId    string    `position:"Query" name:"JobGroupId"`
	Name          string    `position:"Query" name:"Name"`
	ScenarioId    string    `position:"Query" name:"ScenarioId"`
}

// ModifyJobGroupResponse is the response struct for api ModifyJobGroup
type ModifyJobGroupResponse struct {
	*responses.BaseResponse
	Code           string   `json:"Code" xml:"Code"`
	HttpStatusCode int      `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string   `json:"Message" xml:"Message"`
	RequestId      string   `json:"RequestId" xml:"RequestId"`
	Success        bool     `json:"Success" xml:"Success"`
	JobGroup       JobGroup `json:"JobGroup" xml:"JobGroup"`
}

// CreateModifyJobGroupRequest creates a request to invoke ModifyJobGroup API
func CreateModifyJobGroupRequest() (request *ModifyJobGroupRequest) {
	request = &ModifyJobGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OutboundBot", "2019-12-26", "ModifyJobGroup", "outboundbot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyJobGroupResponse creates a response to parse from ModifyJobGroup response
func CreateModifyJobGroupResponse() (response *ModifyJobGroupResponse) {
	response = &ModifyJobGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
