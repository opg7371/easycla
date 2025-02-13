// Copyright The Linux Foundation and each contributor to CommunityBridge.
// SPDX-License-Identifier: MIT

package signatures

import (
	"github.com/communitybridge/easycla/cla-backend-go/gen/v1/models"
	v2Models "github.com/communitybridge/easycla/cla-backend-go/gen/v2/models"
)

// SignatureCompanyID is a simple data model to hold the signature ID and come company details for CCLA's
type SignatureCompanyID struct {
	SignatureID string
	CompanyID   string
	CompanySFID string
	CompanyName string
}

//ApprovalCriteria struct representing approval criteria values
type ApprovalCriteria struct {
	UserEmail      string
	GitHubUsername string
	GitlabUsername string
}

//ApprovalList data model
type ApprovalList struct {
	Criteria                string
	ApprovalList            []string
	Action                  string
	ClaGroupID              string
	ClaGroupName            string
	CompanyID               string
	Version                 string
	EmailApprovals          []string
	DomainApprovals         []string
	GitHubUsernameApprovals []string
	GitHubUsernames         []string
	GitHubOrgApprovals      []string
	GitlabUsernameApprovals []string
	GitlabOrgApprovals      []string
	GitlabUsernames         []string
	GerritICLAECLAs         []string
	ICLAs                   []*models.IclaSignature
	ECLAs                   []*models.Signature
	CLAManager              *models.User
	ManagersInfo            []ClaManagerInfoParams
	CCLASignature           *models.Signature
}

// GerritUserResponse is a data structure to hold the gerrit user query response
type GerritUserResponse struct {
	gerritGroupResponse *v2Models.GerritGroupResponse
	queryType           string
	Error               error
}

// ICLAUserResponse is struct that supports ICLAUsers
type ICLAUserResponse struct {
	ICLASignature *models.IclaSignature
	Error         error
}

const (
	//CCLAICLA representing user removal under CCLA + ICLA
	CCLAICLA = "CCLAICLA"
	//CCLAICLAECLA representing user removal under CCLA + ICLA +ECLA
	CCLAICLAECLA = "CCLAICLAECLA"
	//CCLA representing normal use case of user under CCLA
	CCLA = "ICLA"
	//ICLA representing individual use case
	ICLA = "ICLA"
)
