package ddosbgp

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

// DeleteIp invokes the ddosbgp.DeleteIp API synchronously
// api document: https://help.aliyun.com/api/ddosbgp/deleteip.html
func (client *Client) DeleteIp(request *DeleteIpRequest) (response *DeleteIpResponse, err error) {
	response = CreateDeleteIpResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteIpWithChan invokes the ddosbgp.DeleteIp API asynchronously
// api document: https://help.aliyun.com/api/ddosbgp/deleteip.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteIpWithChan(request *DeleteIpRequest) (<-chan *DeleteIpResponse, <-chan error) {
	responseChan := make(chan *DeleteIpResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteIp(request)
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

// DeleteIpWithCallback invokes the ddosbgp.DeleteIp API asynchronously
// api document: https://help.aliyun.com/api/ddosbgp/deleteip.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteIpWithCallback(request *DeleteIpRequest, callback func(response *DeleteIpResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteIpResponse
		var err error
		defer close(result)
		response, err = client.DeleteIp(request)
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

// DeleteIpRequest is the request struct for api DeleteIp
type DeleteIpRequest struct {
	*requests.RpcRequest
	ResourceGroupId  string `position:"Query" name:"ResourceGroupId"`
	InstanceId       string `position:"Query" name:"InstanceId"`
	SourceIp         string `position:"Query" name:"SourceIp"`
	IpList           string `position:"Query" name:"IpList"`
	ResourceRegionId string `position:"Query" name:"ResourceRegionId"`
}

// DeleteIpResponse is the response struct for api DeleteIp
type DeleteIpResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteIpRequest creates a request to invoke DeleteIp API
func CreateDeleteIpRequest() (request *DeleteIpRequest) {
	request = &DeleteIpRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddosbgp", "2018-07-20", "DeleteIp", "ddosbgp", "openAPI")
	return
}

// CreateDeleteIpResponse creates a response to parse from DeleteIp response
func CreateDeleteIpResponse() (response *DeleteIpResponse) {
	response = &DeleteIpResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
