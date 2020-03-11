package alimt

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

// TranslateGeneral invokes the alimt.TranslateGeneral API synchronously
// api document: https://help.aliyun.com/api/alimt/translategeneral.html
func (client *Client) TranslateGeneral(request *TranslateGeneralRequest) (response *TranslateGeneralResponse, err error) {
	response = CreateTranslateGeneralResponse()
	err = client.DoAction(request, response)
	return
}

// TranslateGeneralWithChan invokes the alimt.TranslateGeneral API asynchronously
// api document: https://help.aliyun.com/api/alimt/translategeneral.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) TranslateGeneralWithChan(request *TranslateGeneralRequest) (<-chan *TranslateGeneralResponse, <-chan error) {
	responseChan := make(chan *TranslateGeneralResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.TranslateGeneral(request)
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

// TranslateGeneralWithCallback invokes the alimt.TranslateGeneral API asynchronously
// api document: https://help.aliyun.com/api/alimt/translategeneral.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) TranslateGeneralWithCallback(request *TranslateGeneralRequest, callback func(response *TranslateGeneralResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *TranslateGeneralResponse
		var err error
		defer close(result)
		response, err = client.TranslateGeneral(request)
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

// TranslateGeneralRequest is the request struct for api TranslateGeneral
type TranslateGeneralRequest struct {
	*requests.RpcRequest
	SourceLanguage string `position:"Body" name:"SourceLanguage"`
	SourceText     string `position:"Body" name:"SourceText"`
	FormatType     string `position:"Body" name:"FormatType"`
	Scene          string `position:"Body" name:"Scene"`
	TargetLanguage string `position:"Body" name:"TargetLanguage"`
}

// TranslateGeneralResponse is the response struct for api TranslateGeneral
type TranslateGeneralResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateTranslateGeneralRequest creates a request to invoke TranslateGeneral API
func CreateTranslateGeneralRequest() (request *TranslateGeneralRequest) {
	request = &TranslateGeneralRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("alimt", "2018-10-12", "TranslateGeneral", "alimtct", "openAPI")
	return
}

// CreateTranslateGeneralResponse creates a response to parse from TranslateGeneral response
func CreateTranslateGeneralResponse() (response *TranslateGeneralResponse) {
	response = &TranslateGeneralResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
