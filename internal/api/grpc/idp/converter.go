package idp

import (
	obj_grpc "github.com/caos/zitadel/internal/api/grpc/object"
	"github.com/caos/zitadel/internal/domain"
	iam_model "github.com/caos/zitadel/internal/iam/model"
	user_model "github.com/caos/zitadel/internal/user/model"
	idp_pb "github.com/caos/zitadel/pkg/grpc/idp"
)

func IDPViewsToPb(idps []*iam_model.IDPConfigView) []*idp_pb.IDP {
	resp := make([]*idp_pb.IDP, len(idps))
	for i, idp := range idps {
		resp[i] = ModelIDPViewToPb(idp)
	}
	return resp
}

func ModelIDPViewToPb(idp *iam_model.IDPConfigView) *idp_pb.IDP {
	return &idp_pb.IDP{
		Id:          idp.IDPConfigID,
		State:       ModelIDPStateToPb(idp.State),
		Name:        idp.Name,
		StylingType: ModelIDPStylingTypeToPb(idp.StylingType),
		Owner:       ModelIDPProviderTypeToPb(idp.IDPProviderType),
		Config:      ModelIDPViewToConfigPb(idp),
		Details: obj_grpc.ToViewDetailsPb(
			idp.Sequence,
			idp.CreationDate,
			idp.ChangeDate,
			"", //TODO: backend
		),
	}
}

func IDPViewToPb(idp *domain.IDPConfigView) *idp_pb.IDP {
	mapped := &idp_pb.IDP{
		Id:          idp.AggregateID,
		State:       IDPStateToPb(idp.State),
		Name:        idp.Name,
		StylingType: IDPStylingTypeToPb(idp.StylingType),
		Config:      IDPViewToConfigPb(idp),
		Details:     obj_grpc.ToViewDetailsPb(idp.Sequence, idp.CreationDate, idp.ChangeDate, ""), //TODO: resource owner in view
	}
	return mapped
}

func ExternalIDPViewsToLoginPolicyLinkPb(links []*iam_model.IDPProviderView) []*idp_pb.IDPLoginPolicyLink {
	l := make([]*idp_pb.IDPLoginPolicyLink, len(links))
	for i, link := range links {
		l[i] = ExternalIDPViewToLoginPolicyLinkPb(link)
	}
	return l
}

func ExternalIDPViewToLoginPolicyLinkPb(link *iam_model.IDPProviderView) *idp_pb.IDPLoginPolicyLink {
	return &idp_pb.IDPLoginPolicyLink{
		IdpId:   link.IDPConfigID,
		IdpName: link.Name,
		IdpType: idp_pb.IDPType_IDP_TYPE_OIDC,
	}
}

func IDPsToUserLinkPb(res []*user_model.ExternalIDPView) []*idp_pb.IDPUserLink {
	links := make([]*idp_pb.IDPUserLink, len(res))
	for i, link := range res {
		links[i] = ExternalIDPViewToUserLinkPb(link)
	}
	return links
}

func ExternalIDPViewToUserLinkPb(link *user_model.ExternalIDPView) *idp_pb.IDPUserLink {
	return &idp_pb.IDPUserLink{
		UserId:           link.UserID,
		IdpId:            link.IDPConfigID,
		IdpName:          link.IDPName,
		ProvidedUserId:   link.ExternalUserID,
		ProvidedUserName: link.UserDisplayName,
		//TODO: as soon as saml is implemented we need to switch here
		IdpType: idp_pb.IDPType_IDP_TYPE_OIDC,
	}
}

func IDPStateToPb(state domain.IDPConfigState) idp_pb.IDPState {
	switch state {
	case domain.IDPConfigStateActive:
		return idp_pb.IDPState_IDP_STATE_ACTIVE
	case domain.IDPConfigStateInactive:
		return idp_pb.IDPState_IDP_STATE_INACTIVE
	default:
		return idp_pb.IDPState_IDP_STATE_UNSPECIFIED
	}
}

func ModelIDPStateToPb(state iam_model.IDPConfigState) idp_pb.IDPState {
	switch state {
	case iam_model.IDPConfigStateActive:
		return idp_pb.IDPState_IDP_STATE_ACTIVE
	case iam_model.IDPConfigStateInactive:
		return idp_pb.IDPState_IDP_STATE_INACTIVE
	default:
		return idp_pb.IDPState_IDP_STATE_UNSPECIFIED
	}
}

func IDPStylingTypeToDomain(stylingType idp_pb.IDPStylingType) domain.IDPConfigStylingType {
	switch stylingType {
	case idp_pb.IDPStylingType_STYLING_TYPE_GOOGLE:
		return domain.IDPConfigStylingTypeGoogle
	default:
		return domain.IDPConfigStylingTypeUnspecified
	}
}

func ModelIDPStylingTypeToPb(stylingType iam_model.IDPStylingType) idp_pb.IDPStylingType {
	switch stylingType {
	case iam_model.IDPStylingTypeGoogle:
		return idp_pb.IDPStylingType_STYLING_TYPE_GOOGLE
	default:
		return idp_pb.IDPStylingType_STYLING_TYPE_UNSPECIFIED
	}
}

func IDPStylingTypeToPb(stylingType domain.IDPConfigStylingType) idp_pb.IDPStylingType {
	switch stylingType {
	case domain.IDPConfigStylingTypeGoogle:
		return idp_pb.IDPStylingType_STYLING_TYPE_GOOGLE
	default:
		return idp_pb.IDPStylingType_STYLING_TYPE_UNSPECIFIED
	}
}

func ModelIDPViewToConfigPb(config *iam_model.IDPConfigView) *idp_pb.IDP_OidcConfig {
	return &idp_pb.IDP_OidcConfig{
		OidcConfig: &idp_pb.OIDCConfig{
			ClientId:           config.OIDCClientID,
			Issuer:             config.OIDCIssuer,
			Scopes:             config.OIDCScopes,
			DisplayNameMapping: ModelMappingFieldToPb(config.OIDCIDPDisplayNameMapping),
			UsernameMapping:    ModelMappingFieldToPb(config.OIDCUsernameMapping),
		},
	}
}

func IDPViewToConfigPb(config *domain.IDPConfigView) *idp_pb.IDP_OidcConfig {
	return &idp_pb.IDP_OidcConfig{
		OidcConfig: &idp_pb.OIDCConfig{
			ClientId:           config.OIDCClientID,
			Issuer:             config.OIDCIssuer,
			Scopes:             config.OIDCScopes,
			DisplayNameMapping: MappingFieldToPb(config.OIDCIDPDisplayNameMapping),
			UsernameMapping:    MappingFieldToPb(config.OIDCUsernameMapping),
		},
	}
}

func OIDCConfigToPb(config *domain.OIDCIDPConfig) *idp_pb.IDP_OidcConfig {
	return &idp_pb.IDP_OidcConfig{
		OidcConfig: &idp_pb.OIDCConfig{
			ClientId:           config.ClientID,
			Issuer:             config.Issuer,
			Scopes:             config.Scopes,
			DisplayNameMapping: MappingFieldToPb(config.IDPDisplayNameMapping),
			UsernameMapping:    MappingFieldToPb(config.UsernameMapping),
		},
	}
}

func FieldNameToModel(fieldName idp_pb.IDPFieldName) iam_model.IDPConfigSearchKey {
	switch fieldName {
	// case admin.IdpSearchKey_IDPSEARCHKEY_IDP_CONFIG_ID: //TODO: not implemented in proto
	// 	return iam_model.IDPConfigSearchKeyIdpConfigID
	case idp_pb.IDPFieldName_IDP_FIELD_NAME_NAME:
		return iam_model.IDPConfigSearchKeyName
	default:
		return iam_model.IDPConfigSearchKeyUnspecified
	}
}

func ModelMappingFieldToPb(mappingField iam_model.OIDCMappingField) idp_pb.OIDCMappingField {
	switch mappingField {
	case iam_model.OIDCMappingFieldEmail:
		return idp_pb.OIDCMappingField_OIDC_MAPPING_FIELD_EMAIL
	case iam_model.OIDCMappingFieldPreferredLoginName:
		return idp_pb.OIDCMappingField_OIDC_MAPPING_FIELD_PREFERRED_USERNAME
	default:
		return idp_pb.OIDCMappingField_OIDC_MAPPING_FIELD_UNSPECIFIED
	}
}

func MappingFieldToPb(mappingField domain.OIDCMappingField) idp_pb.OIDCMappingField {
	switch mappingField {
	case domain.OIDCMappingFieldEmail:
		return idp_pb.OIDCMappingField_OIDC_MAPPING_FIELD_EMAIL
	case domain.OIDCMappingFieldPreferredLoginName:
		return idp_pb.OIDCMappingField_OIDC_MAPPING_FIELD_PREFERRED_USERNAME
	default:
		return idp_pb.OIDCMappingField_OIDC_MAPPING_FIELD_UNSPECIFIED
	}
}

func MappingFieldToDomain(mappingField idp_pb.OIDCMappingField) domain.OIDCMappingField {
	switch mappingField {
	case idp_pb.OIDCMappingField_OIDC_MAPPING_FIELD_EMAIL:
		return domain.OIDCMappingFieldEmail
	case idp_pb.OIDCMappingField_OIDC_MAPPING_FIELD_PREFERRED_USERNAME:
		return domain.OIDCMappingFieldPreferredLoginName
	default:
		return domain.OIDCMappingFieldUnspecified
	}
}

func ModelIDPProviderTypeToPb(typ iam_model.IDPProviderType) idp_pb.IDPOwnerType {
	switch typ {
	case iam_model.IDPProviderTypeOrg:
		return idp_pb.IDPOwnerType_IDP_OWNER_TYPE_ORG
	case iam_model.IDPProviderTypeSystem:
		return idp_pb.IDPOwnerType_IDP_OWNER_TYPE_SYSTEM
	default:
		return idp_pb.IDPOwnerType_IDP_OWNER_TYPE_UNSPECIFIED
	}
}

func IDPProviderTypeFromPb(typ idp_pb.IDPOwnerType) domain.IdentityProviderType {
	switch typ {
	case idp_pb.IDPOwnerType_IDP_OWNER_TYPE_ORG:
		return domain.IdentityProviderTypeOrg
	case idp_pb.IDPOwnerType_IDP_OWNER_TYPE_SYSTEM:
		return domain.IdentityProviderTypeSystem
	default:
		return domain.IdentityProviderTypeOrg
	}
}

func IDPProviderTypeModelFromPb(typ idp_pb.IDPOwnerType) iam_model.IDPProviderType {
	switch typ {
	case idp_pb.IDPOwnerType_IDP_OWNER_TYPE_ORG:
		return iam_model.IDPProviderTypeOrg
	case idp_pb.IDPOwnerType_IDP_OWNER_TYPE_SYSTEM:
		return iam_model.IDPProviderTypeSystem
	default:
		return iam_model.IDPProviderTypeOrg
	}
}

func IDPIDQueryToModel(query *idp_pb.IDPIDQuery) *iam_model.IDPConfigSearchQuery {
	return &iam_model.IDPConfigSearchQuery{
		Key:    iam_model.IDPConfigSearchKeyIdpConfigID,
		Method: domain.SearchMethodEquals,
		Value:  query.Id,
	}
}

func IDPNameQueryToModel(query *idp_pb.IDPNameQuery) *iam_model.IDPConfigSearchQuery {
	return &iam_model.IDPConfigSearchQuery{
		Key:    iam_model.IDPConfigSearchKeyName,
		Method: obj_grpc.TextMethodToModel(query.Method),
		Value:  query.Name,
	}
}

func IDPOwnerTypeQueryToModel(query *idp_pb.IDPOwnerTypeQuery) *iam_model.IDPConfigSearchQuery {
	return &iam_model.IDPConfigSearchQuery{
		Key:    iam_model.IDPConfigSearchKeyIdpProviderType,
		Method: domain.SearchMethodEquals,
		Value:  IDPProviderTypeModelFromPb(query.OwnerType),
	}
}
