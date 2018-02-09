package mts

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

func (client *Client) QueryMediaWorkflowExecutionList(request *QueryMediaWorkflowExecutionListRequest) (response *QueryMediaWorkflowExecutionListResponse, err error) {
	response = CreateQueryMediaWorkflowExecutionListResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) QueryMediaWorkflowExecutionListWithChan(request *QueryMediaWorkflowExecutionListRequest) (<-chan *QueryMediaWorkflowExecutionListResponse, <-chan error) {
	responseChan := make(chan *QueryMediaWorkflowExecutionListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryMediaWorkflowExecutionList(request)
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

func (client *Client) QueryMediaWorkflowExecutionListWithCallback(request *QueryMediaWorkflowExecutionListRequest, callback func(response *QueryMediaWorkflowExecutionListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryMediaWorkflowExecutionListResponse
		var err error
		defer close(result)
		response, err = client.QueryMediaWorkflowExecutionList(request)
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

type QueryMediaWorkflowExecutionListRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	RunIds               string           `position:"Query" name:"RunIds"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
}

type QueryMediaWorkflowExecutionListResponse struct {
	*responses.BaseResponse
	RequestId                  string                                                      `json:"RequestId" xml:"RequestId"`
	NonExistRunIds             NonExistRunIds                                              `json:"NonExistRunIds" xml:"NonExistRunIds"`
	MediaWorkflowExecutionList MediaWorkflowExecutionListInQueryMediaWorkflowExecutionList `json:"MediaWorkflowExecutionList" xml:"MediaWorkflowExecutionList"`
}

func CreateQueryMediaWorkflowExecutionListRequest() (request *QueryMediaWorkflowExecutionListRequest) {
	request = &QueryMediaWorkflowExecutionListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "QueryMediaWorkflowExecutionList", "", "")
	return
}

func CreateQueryMediaWorkflowExecutionListResponse() (response *QueryMediaWorkflowExecutionListResponse) {
	response = &QueryMediaWorkflowExecutionListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
