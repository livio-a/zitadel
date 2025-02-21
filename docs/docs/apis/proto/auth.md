---
title: zitadel/auth.proto
---
> This document reflects the state from API 1.0 (available from 20.04.2021)


## AuthService {#zitadelauthv1authservice}


### Healthz

> **rpc** Healthz([HealthzRequest](#healthzrequest))
[HealthzResponse](#healthzresponse)






### GetMyUser

> **rpc** GetMyUser([GetMyUserRequest](#getmyuserrequest))
[GetMyUserResponse](#getmyuserresponse)

Returns my full blown user




### ListMyUserChanges

> **rpc** ListMyUserChanges([ListMyUserChangesRequest](#listmyuserchangesrequest))
[ListMyUserChangesResponse](#listmyuserchangesresponse)

Returns the history of the authorized user (each event)




### ListMyUserSessions

> **rpc** ListMyUserSessions([ListMyUserSessionsRequest](#listmyusersessionsrequest))
[ListMyUserSessionsResponse](#listmyusersessionsresponse)

Returns the user sessions of the authorized user of the current useragent




### ListMyRefreshTokens

> **rpc** ListMyRefreshTokens([ListMyRefreshTokensRequest](#listmyrefreshtokensrequest))
[ListMyRefreshTokensResponse](#listmyrefreshtokensresponse)

Returns the refresh tokens of the authorized user




### RevokeMyRefreshToken

> **rpc** RevokeMyRefreshToken([RevokeMyRefreshTokenRequest](#revokemyrefreshtokenrequest))
[RevokeMyRefreshTokenResponse](#revokemyrefreshtokenresponse)

Revokes a single refresh token of the authorized user by its (token) id




### RevokeAllMyRefreshTokens

> **rpc** RevokeAllMyRefreshTokens([RevokeAllMyRefreshTokensRequest](#revokeallmyrefreshtokensrequest))
[RevokeAllMyRefreshTokensResponse](#revokeallmyrefreshtokensresponse)

Revokes all refresh tokens of the authorized user




### UpdateMyUserName

> **rpc** UpdateMyUserName([UpdateMyUserNameRequest](#updatemyusernamerequest))
[UpdateMyUserNameResponse](#updatemyusernameresponse)

Change the user name of the authorize user




### GetMyPasswordComplexityPolicy

> **rpc** GetMyPasswordComplexityPolicy([GetMyPasswordComplexityPolicyRequest](#getmypasswordcomplexitypolicyrequest))
[GetMyPasswordComplexityPolicyResponse](#getmypasswordcomplexitypolicyresponse)

Returns the password complexity policy of my organisation
This policy defines how the password should look




### UpdateMyPassword

> **rpc** UpdateMyPassword([UpdateMyPasswordRequest](#updatemypasswordrequest))
[UpdateMyPasswordResponse](#updatemypasswordresponse)

Change the password of the authorized user




### GetMyProfile

> **rpc** GetMyProfile([GetMyProfileRequest](#getmyprofilerequest))
[GetMyProfileResponse](#getmyprofileresponse)

Returns the profile information of the authorized user




### UpdateMyProfile

> **rpc** UpdateMyProfile([UpdateMyProfileRequest](#updatemyprofilerequest))
[UpdateMyProfileResponse](#updatemyprofileresponse)

Changes the profile information of the authorized user




### GetMyEmail

> **rpc** GetMyEmail([GetMyEmailRequest](#getmyemailrequest))
[GetMyEmailResponse](#getmyemailresponse)

Returns the email address of the authorized user




### SetMyEmail

> **rpc** SetMyEmail([SetMyEmailRequest](#setmyemailrequest))
[SetMyEmailResponse](#setmyemailresponse)

Changes the email address of the authorized user
An email is sent to the given address, to verify it




### VerifyMyEmail

> **rpc** VerifyMyEmail([VerifyMyEmailRequest](#verifymyemailrequest))
[VerifyMyEmailResponse](#verifymyemailresponse)

Sets the email address to verified




### ResendMyEmailVerification

> **rpc** ResendMyEmailVerification([ResendMyEmailVerificationRequest](#resendmyemailverificationrequest))
[ResendMyEmailVerificationResponse](#resendmyemailverificationresponse)

Sends a new email to the last given address to verify it




### GetMyPhone

> **rpc** GetMyPhone([GetMyPhoneRequest](#getmyphonerequest))
[GetMyPhoneResponse](#getmyphoneresponse)

Returns the phone number of the authorized user




### SetMyPhone

> **rpc** SetMyPhone([SetMyPhoneRequest](#setmyphonerequest))
[SetMyPhoneResponse](#setmyphoneresponse)

Sets the phone number of the authorized user
An sms is sent to the number with a verification code




### VerifyMyPhone

> **rpc** VerifyMyPhone([VerifyMyPhoneRequest](#verifymyphonerequest))
[VerifyMyPhoneResponse](#verifymyphoneresponse)

Sets the phone number to verified




### ResendMyPhoneVerification

> **rpc** ResendMyPhoneVerification([ResendMyPhoneVerificationRequest](#resendmyphoneverificationrequest))
[ResendMyPhoneVerificationResponse](#resendmyphoneverificationresponse)

Resends a sms to the last given phone number, to verify it




### RemoveMyPhone

> **rpc** RemoveMyPhone([RemoveMyPhoneRequest](#removemyphonerequest))
[RemoveMyPhoneResponse](#removemyphoneresponse)

Removed the phone number of the authorized user




### RemoveMyAvatar

> **rpc** RemoveMyAvatar([RemoveMyAvatarRequest](#removemyavatarrequest))
[RemoveMyAvatarResponse](#removemyavatarresponse)

Remove my avatar




### ListMyLinkedIDPs

> **rpc** ListMyLinkedIDPs([ListMyLinkedIDPsRequest](#listmylinkedidpsrequest))
[ListMyLinkedIDPsResponse](#listmylinkedidpsresponse)

Returns a list of all linked identity providers (social logins, eg. Google, Microsoft, AD, etc.)




### RemoveMyLinkedIDP

> **rpc** RemoveMyLinkedIDP([RemoveMyLinkedIDPRequest](#removemylinkedidprequest))
[RemoveMyLinkedIDPResponse](#removemylinkedidpresponse)

Removes a linked identity provider (social logins, eg. Google, Microsoft, AD, etc.)




### ListMyAuthFactors

> **rpc** ListMyAuthFactors([ListMyAuthFactorsRequest](#listmyauthfactorsrequest))
[ListMyAuthFactorsResponse](#listmyauthfactorsresponse)

Returns all configured authentication factors (second and multi)




### AddMyAuthFactorOTP

> **rpc** AddMyAuthFactorOTP([AddMyAuthFactorOTPRequest](#addmyauthfactorotprequest))
[AddMyAuthFactorOTPResponse](#addmyauthfactorotpresponse)

Adds a new OTP (One Time Password) Second Factor to the authorized user
Only one OTP can be configured per user




### VerifyMyAuthFactorOTP

> **rpc** VerifyMyAuthFactorOTP([VerifyMyAuthFactorOTPRequest](#verifymyauthfactorotprequest))
[VerifyMyAuthFactorOTPResponse](#verifymyauthfactorotpresponse)

Verify the last added OTP (One Time Password)




### RemoveMyAuthFactorOTP

> **rpc** RemoveMyAuthFactorOTP([RemoveMyAuthFactorOTPRequest](#removemyauthfactorotprequest))
[RemoveMyAuthFactorOTPResponse](#removemyauthfactorotpresponse)

Removed the configured OTP (One Time Password) Factor




### AddMyAuthFactorU2F

> **rpc** AddMyAuthFactorU2F([AddMyAuthFactorU2FRequest](#addmyauthfactoru2frequest))
[AddMyAuthFactorU2FResponse](#addmyauthfactoru2fresponse)

Adds a new U2F (Universal Second Factor) to the authorized user
Multiple U2Fs can be configured




### VerifyMyAuthFactorU2F

> **rpc** VerifyMyAuthFactorU2F([VerifyMyAuthFactorU2FRequest](#verifymyauthfactoru2frequest))
[VerifyMyAuthFactorU2FResponse](#verifymyauthfactoru2fresponse)

Verifies the last added U2F (Universal Second Factor) of the authorized user




### RemoveMyAuthFactorU2F

> **rpc** RemoveMyAuthFactorU2F([RemoveMyAuthFactorU2FRequest](#removemyauthfactoru2frequest))
[RemoveMyAuthFactorU2FResponse](#removemyauthfactoru2fresponse)

Removes the U2F Authentication from the authorized user




### ListMyPasswordless

> **rpc** ListMyPasswordless([ListMyPasswordlessRequest](#listmypasswordlessrequest))
[ListMyPasswordlessResponse](#listmypasswordlessresponse)

Returns all configured passwordless authentications of the authorized user




### AddMyPasswordless

> **rpc** AddMyPasswordless([AddMyPasswordlessRequest](#addmypasswordlessrequest))
[AddMyPasswordlessResponse](#addmypasswordlessresponse)

Adds a new passwordless authentications to the authorized user
Multiple passwordless authentications can be configured




### VerifyMyPasswordless

> **rpc** VerifyMyPasswordless([VerifyMyPasswordlessRequest](#verifymypasswordlessrequest))
[VerifyMyPasswordlessResponse](#verifymypasswordlessresponse)

Verifies the last added passwordless configuration




### RemoveMyPasswordless

> **rpc** RemoveMyPasswordless([RemoveMyPasswordlessRequest](#removemypasswordlessrequest))
[RemoveMyPasswordlessResponse](#removemypasswordlessresponse)

Removes the passwordless configuration from the authorized user




### ListMyUserGrants

> **rpc** ListMyUserGrants([ListMyUserGrantsRequest](#listmyusergrantsrequest))
[ListMyUserGrantsResponse](#listmyusergrantsresponse)

Returns all user grants (authorizations) of the authorized user




### ListMyProjectOrgs

> **rpc** ListMyProjectOrgs([ListMyProjectOrgsRequest](#listmyprojectorgsrequest))
[ListMyProjectOrgsResponse](#listmyprojectorgsresponse)

Returns a list of organisations where the authorized user has a user grant (authorization) in the context of the requested project




### ListMyZitadelFeatures

> **rpc** ListMyZitadelFeatures([ListMyZitadelFeaturesRequest](#listmyzitadelfeaturesrequest))
[ListMyZitadelFeaturesResponse](#listmyzitadelfeaturesresponse)

Returns a list of features, which are allowed on these organisation based on the subscription of the organisation




### ListMyZitadelPermissions

> **rpc** ListMyZitadelPermissions([ListMyZitadelPermissionsRequest](#listmyzitadelpermissionsrequest))
[ListMyZitadelPermissionsResponse](#listmyzitadelpermissionsresponse)

Returns the permissions the authorized user has in ZITADEL based on his manager roles (e.g ORG_OWNER)




### ListMyProjectPermissions

> **rpc** ListMyProjectPermissions([ListMyProjectPermissionsRequest](#listmyprojectpermissionsrequest))
[ListMyProjectPermissionsResponse](#listmyprojectpermissionsresponse)

Returns a list of roles for the authorized user and project









## Messages


### AddMyAuthFactorOTPRequest
This is an empty request




### AddMyAuthFactorOTPResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| url |  string | - |  |
| secret |  string | - |  |
| details |  zitadel.v1.ObjectDetails | - |  |




### AddMyAuthFactorU2FRequest
This is an empty request




### AddMyAuthFactorU2FResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| key |  zitadel.user.v1.WebAuthNKey | - |  |
| details |  zitadel.v1.ObjectDetails | - |  |




### AddMyPasswordlessRequest
This is an empty request




### AddMyPasswordlessResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| key |  zitadel.user.v1.WebAuthNKey | - |  |
| details |  zitadel.v1.ObjectDetails | - |  |




### GetMyEmailRequest
This is an empty request




### GetMyEmailResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |
| email |  zitadel.user.v1.Email | - |  |




### GetMyPasswordComplexityPolicyRequest
This is an empty request




### GetMyPasswordComplexityPolicyResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| policy |  zitadel.policy.v1.PasswordComplexityPolicy | - |  |




### GetMyPhoneRequest
This is an empty request




### GetMyPhoneResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |
| phone |  zitadel.user.v1.Phone | - |  |




### GetMyProfileRequest
This is an empty request




### GetMyProfileResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |
| profile |  zitadel.user.v1.Profile | - |  |




### GetMyUserRequest
This is an empty request
the request parameters are read from the token-header




### GetMyUserResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| user |  zitadel.user.v1.User | - |  |
| last_login |  google.protobuf.Timestamp | - |  |




### HealthzRequest
This is an empty request




### HealthzResponse
This is an empty response




### ListMyAuthFactorsRequest
This is an empty request




### ListMyAuthFactorsResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| result | repeated zitadel.user.v1.AuthFactor | - |  |




### ListMyLinkedIDPsRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| query |  zitadel.v1.ListQuery | list limitations and ordering |  |




### ListMyLinkedIDPsResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ListDetails | - |  |
| result | repeated zitadel.idp.v1.IDPUserLink | - |  |




### ListMyPasswordlessRequest
This is an empty request




### ListMyPasswordlessResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| result | repeated zitadel.user.v1.WebAuthNToken | - |  |




### ListMyProjectOrgsRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| query |  zitadel.v1.ListQuery | list limitations and ordering |  |
| queries | repeated zitadel.org.v1.OrgQuery | criterias the client is looking for |  |




### ListMyProjectOrgsResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ListDetails | - |  |
| result | repeated zitadel.org.v1.Org | - |  |




### ListMyProjectPermissionsRequest
This is an empty request




### ListMyProjectPermissionsResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| result | repeated string | - |  |




### ListMyRefreshTokensRequest
This is an empty request




### ListMyRefreshTokensResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ListDetails | - |  |
| result | repeated zitadel.user.v1.RefreshToken | - |  |




### ListMyUserChangesRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| query |  zitadel.change.v1.ChangeQuery | - |  |




### ListMyUserChangesResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ListDetails | - |  |
| result | repeated zitadel.change.v1.Change | - |  |




### ListMyUserGrantsRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| query |  zitadel.v1.ListQuery | list limitations and ordering |  |




### ListMyUserGrantsResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ListDetails | - |  |
| result | repeated UserGrant | - |  |




### ListMyUserSessionsRequest
This is an empty request




### ListMyUserSessionsResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| result | repeated zitadel.user.v1.Session | - |  |




### ListMyZitadelFeaturesRequest
This is an empty request




### ListMyZitadelFeaturesResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| result | repeated string | - |  |




### ListMyZitadelPermissionsRequest
This is an empty request




### ListMyZitadelPermissionsResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| result | repeated string | - |  |




### RemoveMyAuthFactorOTPRequest
This is an empty request




### RemoveMyAuthFactorOTPResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |




### RemoveMyAuthFactorU2FRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| token_id |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |




### RemoveMyAuthFactorU2FResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |




### RemoveMyAvatarRequest
This is an empty request




### RemoveMyAvatarResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |




### RemoveMyLinkedIDPRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| idp_id |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |
| linked_user_id |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |




### RemoveMyLinkedIDPResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |




### RemoveMyPasswordlessRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| token_id |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |




### RemoveMyPasswordlessResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |




### RemoveMyPhoneRequest
This is an empty request




### RemoveMyPhoneResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |




### ResendMyEmailVerificationRequest
This is an empty request




### ResendMyEmailVerificationResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |




### ResendMyPhoneVerificationRequest
This is an empty request




### ResendMyPhoneVerificationResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |




### RevokeAllMyRefreshTokensRequest
This is an empty request




### RevokeAllMyRefreshTokensResponse
This is an empty response




### RevokeMyRefreshTokenRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| id |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |




### RevokeMyRefreshTokenResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |




### SetMyEmailRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| email |  string | TODO: check if no value is allowed | string.email: true<br />  |




### SetMyEmailResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |




### SetMyPhoneRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| phone |  string | - | string.min_len: 1<br /> string.max_len: 50<br /> string.prefix: +<br />  |




### SetMyPhoneResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |




### UpdateMyPasswordRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| old_password |  string | - | string.min_len: 1<br /> string.max_bytes: 70<br />  |
| new_password |  string | - | string.min_len: 1<br /> string.max_bytes: 70<br />  |




### UpdateMyPasswordResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |




### UpdateMyProfileRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| first_name |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |
| last_name |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |
| nick_name |  string | - | string.max_len: 200<br />  |
| display_name |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |
| preferred_language |  string | - | string.max_len: 10<br />  |
| gender |  zitadel.user.v1.Gender | - |  |




### UpdateMyProfileResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |




### UpdateMyUserNameRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| user_name |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |




### UpdateMyUserNameResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |




### UserGrant



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| org_id |  string | - |  |
| project_id |  string | - |  |
| user_id |  string | - |  |
| roles | repeated string | - |  |
| org_name |  string | - |  |
| grant_id |  string | - |  |




### VerifyMyAuthFactorOTPRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| code |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |




### VerifyMyAuthFactorOTPResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |




### VerifyMyAuthFactorU2FRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| verification |  zitadel.user.v1.WebAuthNVerification | - | message.required: true<br />  |




### VerifyMyAuthFactorU2FResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |




### VerifyMyEmailRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| code |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |




### VerifyMyEmailResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |




### VerifyMyPasswordlessRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| verification |  zitadel.user.v1.WebAuthNVerification | - | message.required: true<br />  |




### VerifyMyPasswordlessResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |




### VerifyMyPhoneRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| code |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |




### VerifyMyPhoneResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |






