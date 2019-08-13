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

// DescribeInstanceList invokes the ddosbgp.DescribeInstanceList API synchronously
// api document: https://help.aliyun.com/api/ddosbgp/describeinstancelist.html
func (client *Client) DescribeInstanceList(request *DescribeInstanceListRequest) (response *DescribeInstanceListResponse, err error) {
	response = CreateDescribeInstanceListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeInstanceListWithChan invokes the ddosbgp.DescribeInstanceList API asynchronously
// api document: https://help.aliyun.com/api/ddosbgp/describeinstancelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeInstanceListWithChan(request *DescribeInstanceListRequest) (<-chan *DescribeInstanceListResponse, <-chan error) {
	responseChan := make(chan *DescribeInstanceListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeInstanceList(request)
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

// DescribeInstanceListWithCallback invokes the ddosbgp.DescribeInstanceList API asynchronously
// api document: https://help.aliyun.com/api/ddosbgp/describeinstancelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeInstanceListWithCallback(request *DescribeInstanceListRequest, callback func(response *DescribeInstanceListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeInstanceListResponse
		var err error
		defer close(result)
		response, err = client.DescribeInstanceList(request)
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

// DescribeInstanceListRequest is the request struct for api DescribeInstanceList
type DescribeInstanceListRequest struct {
	*requests.RpcRequest
	Ip              string                     `position:"Query" name:"Ip"`
	Orderby         string                     `position:"Query" name:"Orderby"`
	Remark          string                     `position:"Query" name:"Remark"`
	ResourceGroupId string                     `position:"Query" name:"ResourceGroupId"`
	SourceIp        string                     `position:"Query" name:"SourceIp"`
	InstanceIdList  string                     `position:"Query" name:"InstanceIdList"`
	PageNo          requests.Integer           `position:"Query" name:"PageNo"`
	Orderdire       string                     `position:"Query" name:"Orderdire"`
	PageSize        requests.Integer           `position:"Query" name:"PageSize"`
	DdosRegionId    string                     `position:"Query" name:"DdosRegionId"`
	InstanceType    string                     `position:"Query" name:"InstanceType"`
	IpVersion       string                     `position:"Query" name:"IpVersion"`
	Tag             *[]DescribeInstanceListTag `position:"Query" name:"Tag"  type:"Repeated"`
}

// DescribeInstanceListTag is a repeated param struct in DescribeInstanceListRequest
type DescribeInstanceListTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// DescribeInstanceListResponse is the response struct for api DescribeInstanceList
type DescribeInstanceListResponse struct {
	*responses.BaseResponse
	RequestId    string     `json:"RequestId" xml:"RequestId"`
	Total        int64      `json:"Total" xml:"Total"`
	InstanceList []Instance `json:"InstanceList" xml:"InstanceList"`
}

// CreateDescribeInstanceListRequest creates a request to invoke DescribeInstanceList API
func CreateDescribeInstanceListRequest() (request *DescribeInstanceListRequest) {
	request = &DescribeInstanceListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddosbgp", "2018-07-20", "DescribeInstanceList", "ddosbgp", "openAPI")
	return
}

// CreateDescribeInstanceListResponse creates a response to parse from DescribeInstanceList response
func CreateDescribeInstanceListResponse() (response *DescribeInstanceListResponse) {
	response = &DescribeInstanceListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
