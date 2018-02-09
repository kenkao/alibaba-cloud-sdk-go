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

type GetDomainDetailModel struct {
	GmtCreated              string                           `json:"GmtCreated" xml:"GmtCreated"`
	GmtModified             string                           `json:"GmtModified" xml:"GmtModified"`
	SourceType              string                           `json:"SourceType" xml:"SourceType"`
	DomainStatus            string                           `json:"DomainStatus" xml:"DomainStatus"`
	SourcePort              int                              `json:"SourcePort" xml:"SourcePort"`
	CdnType                 string                           `json:"CdnType" xml:"CdnType"`
	Cname                   string                           `json:"Cname" xml:"Cname"`
	HttpsCname              string                           `json:"HttpsCname" xml:"HttpsCname"`
	DomainName              string                           `json:"DomainName" xml:"DomainName"`
	Description             string                           `json:"Description" xml:"Description"`
	ServerCertificateStatus string                           `json:"ServerCertificateStatus" xml:"ServerCertificateStatus"`
	ServerCertificate       string                           `json:"ServerCertificate" xml:"ServerCertificate"`
	Region                  string                           `json:"Region" xml:"Region"`
	Scope                   string                           `json:"Scope" xml:"Scope"`
	CertificateName         string                           `json:"CertificateName" xml:"CertificateName"`
	ResourceGroupId         string                           `json:"ResourceGroupId" xml:"ResourceGroupId"`
	Sources                 SourcesInDescribeCdnDomainDetail `json:"Sources" xml:"Sources"`
	SourceModels            SourceModels                     `json:"SourceModels" xml:"SourceModels"`
}
