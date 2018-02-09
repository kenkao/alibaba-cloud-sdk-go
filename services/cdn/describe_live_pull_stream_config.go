package cdn

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

func (client *Client) DescribeLivePullStreamConfig(request *DescribeLivePullStreamConfigRequest) (response *DescribeLivePullStreamConfigResponse, err error) {
	response = CreateDescribeLivePullStreamConfigResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeLivePullStreamConfigWithChan(request *DescribeLivePullStreamConfigRequest) (<-chan *DescribeLivePullStreamConfigResponse, <-chan error) {
	responseChan := make(chan *DescribeLivePullStreamConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLivePullStreamConfig(request)
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

func (client *Client) DescribeLivePullStreamConfigWithCallback(request *DescribeLivePullStreamConfigRequest, callback func(response *DescribeLivePullStreamConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLivePullStreamConfigResponse
		var err error
		defer close(result)
		response, err = client.DescribeLivePullStreamConfig(request)
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

type DescribeLivePullStreamConfigRequest struct {
	*requests.RpcRequest
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	DomainName    string           `position:"Query" name:"DomainName"`
}

type DescribeLivePullStreamConfigResponse struct {
	*responses.BaseResponse
	RequestId         string                                          `json:"RequestId" xml:"RequestId"`
	LiveAppRecordList LiveAppRecordListInDescribeLivePullStreamConfig `json:"LiveAppRecordList" xml:"LiveAppRecordList"`
}

func CreateDescribeLivePullStreamConfigRequest() (request *DescribeLivePullStreamConfigRequest) {
	request = &DescribeLivePullStreamConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLivePullStreamConfig", "", "")
	return
}

func CreateDescribeLivePullStreamConfigResponse() (response *DescribeLivePullStreamConfigResponse) {
	response = &DescribeLivePullStreamConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
