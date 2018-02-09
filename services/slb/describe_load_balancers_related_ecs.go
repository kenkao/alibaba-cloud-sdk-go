package slb

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

func (client *Client) DescribeLoadBalancersRelatedEcs(request *DescribeLoadBalancersRelatedEcsRequest) (response *DescribeLoadBalancersRelatedEcsResponse, err error) {
	response = CreateDescribeLoadBalancersRelatedEcsResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeLoadBalancersRelatedEcsWithChan(request *DescribeLoadBalancersRelatedEcsRequest) (<-chan *DescribeLoadBalancersRelatedEcsResponse, <-chan error) {
	responseChan := make(chan *DescribeLoadBalancersRelatedEcsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLoadBalancersRelatedEcs(request)
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

func (client *Client) DescribeLoadBalancersRelatedEcsWithCallback(request *DescribeLoadBalancersRelatedEcsRequest, callback func(response *DescribeLoadBalancersRelatedEcsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLoadBalancersRelatedEcsResponse
		var err error
		defer close(result)
		response, err = client.DescribeLoadBalancersRelatedEcs(request)
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

type DescribeLoadBalancersRelatedEcsRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	AccessKeyId          string           `position:"Query" name:"access_key_id"`
	Tags                 string           `position:"Query" name:"Tags"`
	LoadBalancerId       string           `position:"Query" name:"LoadBalancerId"`
}

type DescribeLoadBalancersRelatedEcsResponse struct {
	*responses.BaseResponse
	Message       string                                         `json:"Message" xml:"Message"`
	Success       bool                                           `json:"Success" xml:"Success"`
	RequestId     string                                         `json:"RequestId" xml:"RequestId"`
	LoadBalancers LoadBalancersInDescribeLoadBalancersRelatedEcs `json:"LoadBalancers" xml:"LoadBalancers"`
}

func CreateDescribeLoadBalancersRelatedEcsRequest() (request *DescribeLoadBalancersRelatedEcsRequest) {
	request = &DescribeLoadBalancersRelatedEcsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "DescribeLoadBalancersRelatedEcs", "", "")
	return
}

func CreateDescribeLoadBalancersRelatedEcsResponse() (response *DescribeLoadBalancersRelatedEcsResponse) {
	response = &DescribeLoadBalancersRelatedEcsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
