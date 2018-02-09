package ddospro

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

func (client *Client) DescribeDomainBlackWhiteList(request *DescribeDomainBlackWhiteListRequest) (response *DescribeDomainBlackWhiteListResponse, err error) {
	response = CreateDescribeDomainBlackWhiteListResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeDomainBlackWhiteListWithChan(request *DescribeDomainBlackWhiteListRequest) (<-chan *DescribeDomainBlackWhiteListResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainBlackWhiteListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainBlackWhiteList(request)
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

func (client *Client) DescribeDomainBlackWhiteListWithCallback(request *DescribeDomainBlackWhiteListRequest, callback func(response *DescribeDomainBlackWhiteListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainBlackWhiteListResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainBlackWhiteList(request)
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

type DescribeDomainBlackWhiteListRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Domain          string           `position:"Query" name:"Domain"`
}

type DescribeDomainBlackWhiteListResponse struct {
	*responses.BaseResponse
	RequestId string                             `json:"RequestId" xml:"RequestId"`
	Data      DataInDescribeDomainBlackWhiteList `json:"Data" xml:"Data"`
}

func CreateDescribeDomainBlackWhiteListRequest() (request *DescribeDomainBlackWhiteListRequest) {
	request = &DescribeDomainBlackWhiteListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DDoSPro", "2017-07-25", "DescribeDomainBlackWhiteList", "", "")
	return
}

func CreateDescribeDomainBlackWhiteListResponse() (response *DescribeDomainBlackWhiteListResponse) {
	response = &DescribeDomainBlackWhiteListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
