A small tool to make API requests against a Keycloak backend. You'll receive a JWT and the corresponding roles to the user you are logging in as in JSON format.

Keep in mind, that you have to edit the api.go file.
Change the 'Host, User, clientSecret, Realm' 'keycloakClient:' under the 'NewAPIServer' to the appropriate values.
