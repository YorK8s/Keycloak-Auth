package keycloak

import (
	"context"

	"github.com/Nerzal/gocloak/v13"
	"github.com/golang-jwt/jwt/v5"
)

type Keycloak struct {
	gocloak      gocloak.GoCloak // keycloak client
	clientId     string          // clientId specified in Keycloak
	clientSecret string          // client secret specified in Keycloak
	realm        string          // realm specified in Keycloak
}
type KeycloakLoginResponse struct {
	AccessToken, RefreshToken string
	ExpiresIn                 int
}

type MappingsRepresentation struct {
	Roles interface{}
}

func InitKeycloak(basePath, clientID, clientSecret, realm string) *Keycloak {
	return &Keycloak{
		gocloak:      *gocloak.NewClient(basePath),
		clientId:     clientID,
		clientSecret: clientSecret,
		realm:        realm,
	}
}

func (keycloak Keycloak) Login(ctx context.Context, username, password string) (*KeycloakLoginResponse, error) {
	result, err := keycloak.gocloak.Login(ctx, keycloak.clientId, keycloak.clientSecret, keycloak.realm, username, password)
	if err != nil {
		return nil, err
	}

	return &KeycloakLoginResponse{
		AccessToken:  result.AccessToken,
		RefreshToken: result.RefreshToken,
		ExpiresIn:    result.ExpiresIn,
	}, nil
}

func (keycloak Keycloak) DecodeJWT(ctx context.Context, AccessToken string) (*MappingsRepresentation, error) {
	var tokenString = AccessToken

	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	claims := token.Claims.(jwt.MapClaims)
	if err != nil {
		return nil, err
	}

	return &MappingsRepresentation{
		Roles: claims["realm_access"],
	}, nil
}

/*
func DecodeJWTT(ctx context.Context, AccessToken string) (*MappingsRepresentation, error) {
	claims := jwt.MapClaims{}
	_, _, err := jwt.NewParser().ParseUnverified(AccessToken, claims)
	if err != nil {
		return nil, err
	}
	for key, val := range claims {
		fmt.Printf("Key: %v, value: %v\n", key, val)
	}
	fmt.Println(AccessToken)
	return Roles, nil
}
*/

/*
func (keycloak Keycloak) GetRoleMappingByUserID(ctx context.Context, accessToken string) (*MappingsRepresentation, error) {
	subject, err := keycloak.GetSubjectFromToken(ctx, accessToken)
	if err != nil {
		return nil, err
	}

	result, err := keycloak.gocloak.GetRoleMappingByUserID(ctx, accessToken, keycloak.realm, *subject)
	if err != nil {
		return nil, err
	}
	fmt.Println(result.RealmMappings)
	fmt.Println(result.ClientMappings)
	return nil, nil
}
*/

func (keycloak Keycloak) ValidateToken(ctx context.Context, token string) (bool, error) {
	result, err := keycloak.gocloak.RetrospectToken(ctx, token, keycloak.clientId, keycloak.clientSecret, keycloak.realm)
	if err != nil {
		return false, err
	}

	return *result.Active, nil
}

type gocloakJWTResponse struct {
	Sub string `json:"sub"`
}

/*
func (keycloak Keycloak) GetSubjectFromToken(ctx context.Context, token string) (*string, error) {
	result, _, err := keycloak.gocloak.DecodeAccessToken(ctx, token, keycloak.realm)
	if err != nil {
		return nil, err
	}

	var parsedResult gocloakJWTResponse

	resultBytes, err := json.Marshal(result.Claims)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(resultBytes, &parsedResult); err != nil {
		return nil, err
	}

	return &parsedResult.Sub, nil
}
*/
