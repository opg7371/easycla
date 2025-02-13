// Copyright The Linux Foundation and each contributor to CommunityBridge.
// SPDX-License-Identifier: MIT

package sign

import (
	"errors"
	"fmt"
	"strings"

	log "github.com/communitybridge/easycla/cla-backend-go/logging"
	"github.com/sirupsen/logrus"

	"github.com/communitybridge/easycla/cla-backend-go/projects_cla_groups"

	"github.com/LF-Engineering/lfx-kit/auth"
	"github.com/communitybridge/easycla/cla-backend-go/gen/v2/models"
	"github.com/communitybridge/easycla/cla-backend-go/gen/v2/restapi/operations"
	"github.com/communitybridge/easycla/cla-backend-go/gen/v2/restapi/operations/sign"
	"github.com/communitybridge/easycla/cla-backend-go/utils"
	"github.com/communitybridge/easycla/cla-backend-go/v2/organization-service/client/organizations"
	"github.com/go-openapi/runtime/middleware"
)

// Configure API call
func Configure(api *operations.EasyclaAPI, service Service) {
	// Retrieve a list of available templates
	api.SignRequestCorporateSignatureHandler = sign.RequestCorporateSignatureHandlerFunc(
		func(params sign.RequestCorporateSignatureParams, user *auth.User) middleware.Responder {
			reqID := utils.GetRequestID(params.XREQUESTID)
			ctx := utils.ContextWithRequestAndUser(params.HTTPRequest.Context(), reqID, user) // nolint
			utils.SetAuthUserProperties(user, params.XUSERNAME, params.XEMAIL)
			f := logrus.Fields{
				"functionName":   "v2.sign.handlers.SignRequestCorporateSignatureHandler",
				utils.XREQUESTID: ctx.Value(utils.XREQUESTID),
				"CompanyID":      params.Input.CompanySfid,
				"ProjectSFID":    params.Input.ProjectSfid,
				"authUserName":   utils.StringValue(params.XUSERNAME),
				"authUserEmail":  utils.StringValue(params.XEMAIL),
			}

			if !utils.IsUserAuthorizedForProjectOrganizationTree(ctx, user, utils.StringValue(params.Input.ProjectSfid), utils.StringValue(params.Input.CompanySfid), utils.DISALLOW_ADMIN_SCOPE) {
				msg := fmt.Sprintf("user %s does not have access to Request Corporate Signature with Project|Organization scope tree of %s | %s - allow admin scope: false",
					user.UserName, utils.StringValue(params.Input.ProjectSfid), utils.StringValue(params.Input.CompanySfid))
				log.WithFields(f).Warn(msg)
				return sign.NewRequestCorporateSignatureForbidden().WithPayload(utils.ErrorResponseForbidden(reqID, msg))
			}

			resp, err := service.RequestCorporateSignature(ctx, utils.StringValue(params.XUSERNAME), params.Authorization, params.Input)
			if err != nil {
				if strings.Contains(err.Error(), "does not exist") {
					return sign.NewRequestCorporateSignatureNotFound().WithPayload(errorResponse(reqID, err))
				}
				if strings.Contains(err.Error(), "internal server error") {
					return sign.NewRequestCorporateSignatureInternalServerError().WithPayload(errorResponse(reqID, err))
				}
				if err == projects_cla_groups.ErrProjectNotAssociatedWithClaGroup {
					return sign.NewRequestCorporateSignatureBadRequest().WithPayload(errorResponse(reqID, err))
				}
				if err == ErrCCLANotEnabled || err == ErrTemplateNotConfigured {
					return sign.NewRequestCorporateSignatureBadRequest().WithPayload(errorResponse(reqID, err))
				}
				if _, ok := err.(*organizations.ListOrgUsrAdminScopesNotFound); ok {
					formatErr := errors.New("user role scopes not found for cla-signatory role ")
					return sign.NewRequestCorporateSignatureNotFound().WithPayload(errorResponse(reqID, formatErr))
				}
				if _, ok := err.(*organizations.CreateOrgUsrRoleScopesConflict); ok {
					formatErr := errors.New("user role scope conflict")
					return sign.NewRequestCorporateSignatureConflict().WithPayload(errorResponse(reqID, formatErr))
				}
				if err == ErrNotInOrg {
					return sign.NewRequestCorporateSignatureConflict().WithPayload(errorResponse(reqID, err))
				}
				return sign.NewRequestCorporateSignatureBadRequest().WithPayload(errorResponse(reqID, err))
			}
			return sign.NewRequestCorporateSignatureOK().WithPayload(resp)
		})
}

type codedResponse interface {
	Code() string
}

func errorResponse(reqID string, err error) *models.ErrorResponse {
	code := ""
	if e, ok := err.(codedResponse); ok {
		code = e.Code()
	}

	e := models.ErrorResponse{
		Code:       code,
		Message:    err.Error(),
		XRequestID: reqID,
	}

	return &e
}
