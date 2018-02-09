package ecs

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

type Image struct {
	Progress             string               `json:"Progress" xml:"Progress"`
	ImageId              string               `json:"ImageId" xml:"ImageId"`
	ImageName            string               `json:"ImageName" xml:"ImageName"`
	ImageVersion         string               `json:"ImageVersion" xml:"ImageVersion"`
	Description          string               `json:"Description" xml:"Description"`
	Size                 int                  `json:"Size" xml:"Size"`
	ImageOwnerAlias      string               `json:"ImageOwnerAlias" xml:"ImageOwnerAlias"`
	IsSupportIoOptimized bool                 `json:"IsSupportIoOptimized" xml:"IsSupportIoOptimized"`
	IsSupportCloudinit   bool                 `json:"IsSupportCloudinit" xml:"IsSupportCloudinit"`
	OSName               string               `json:"OSName" xml:"OSName"`
	Architecture         string               `json:"Architecture" xml:"Architecture"`
	Status               string               `json:"Status" xml:"Status"`
	ProductCode          string               `json:"ProductCode" xml:"ProductCode"`
	IsSubscribed         bool                 `json:"IsSubscribed" xml:"IsSubscribed"`
	CreationTime         string               `json:"CreationTime" xml:"CreationTime"`
	IsSelfShared         string               `json:"IsSelfShared" xml:"IsSelfShared"`
	OSType               string               `json:"OSType" xml:"OSType"`
	Platform             string               `json:"Platform" xml:"Platform"`
	Usage                string               `json:"Usage" xml:"Usage"`
	IsCopied             bool                 `json:"IsCopied" xml:"IsCopied"`
	DiskDeviceMappings   DiskDeviceMappings   `json:"DiskDeviceMappings" xml:"DiskDeviceMappings"`
	Tags                 TagsInDescribeImages `json:"Tags" xml:"Tags"`
}
