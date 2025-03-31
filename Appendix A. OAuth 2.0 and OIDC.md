(+) The fundamental of `OAuth 2.0 is to fix the access delegation problem`, OIDC (OpenID Connect) is an `identity layer built on top of OAuth 2.0`.

#### A. The access delegation problem
(+) If you want someone else to access a resource (a microservice, an API, and so on) on your behalf and do something with the resource, you need to `delegate the corresponding access rights to that person (or thing)`. 
==> If you want a 3rd party to read FB status message, corresponding rights need to be given to that 3rd party.

(+) There are 2 models of access delegation, `credential sharing, and no credential sharing`. The first model is quite risky, as with the credentials, the 3rd party will have authorization to do anything that wee can do with our account.

###### How does OAuth 20 fix the access delegation problem
![[Pasted image 20250315213113.png]]

(+) With OAuth 2.0, the third-party web application` first redirects the user to Facebook` (where the user belongs). Facebook authenticates and `gets the user’s consent to share a temporary token` with a third-party web application, which is `only good enough to read the user’s Facebook status messages` for a limited time. Once the web application `gets the token from Facebook, it uses the token along with the API calls to Facebook`.

==> The temporary token Facebook issues has a limited lifetime and is bound to the Facebook user, the third-party web application, and the purpose. The purpose of the token here is to read the user’s Facebook status messages, and the token should be only good enough to do just that and no more

###### Terminology of OAuth 2.0
(+) `Resource owner (RO)`: decides who should have which level of access to the resources owned.  In this case the Facebook User
(+) `Authorization service (AS)`: knows how to authenticate (or identify) the `resource owner` and grant access to 3rd party applications to access resources owned by the `resource owner`, with their consent.
(+) `Resource server (RS)`: guards the resources owned by the `resource owner`, and lets someone access a resource only if the access request comes along with a valid token issued by the `authorization service`.
(+) `Client`: the client consumes a resource on behalf of the `resource owner`.
(+) `Access token (AT)`: is the token that `AS` provides `client`. To validate the `AT` the `RS` may talk to the `AS`. 
(+) `Scope`: is the purpose of the `token`, the `RS` makes sure a given token can be used only for the scope attached to it. if the 3rd party application tries to write to the user's Facebook wall with the access token it got to read the status messages, the request will fail.
(+) `Grant flow`: is the flow of events that happens while the 3rd party gets the token, which is defined by `grant type`.

#### B. Grant types
(+) Different types of applications bearing different characteristics can consume your microservices. The `client application's` different ways to get an `access token` from the `authorization server`, is called grant type in OAuth 2.0.

(+) The standard OAuth 2.0 specification identifies five main grant types. Each grant type outlines the steps for obtaining an access token:
	(-) `Client credentials`: suitable for authentication between 2 systems with no end user.
	(-) `Resource owner password`: Suitable for trusted application 
	(-) `Authorization code`: Suitable for almost all the application with an end user
	(-) `Implicit`:
	(-) `Refresh token`: Used for renewing expired access token

##### B1. Client credentials grant type
(+) This grant type has only `two participants`: the `Client application` and the `AS`, there is no separate `Resource Owner` the `Client` itself is the `RO`. 

(+) Each clients carries its own credentials, composing of a `Client ID` and `Client secret`. The credentials should be stored securely (should not be in cleartext, instead encrypt it and store in a persistent storage).

(+) In this grant type, the `client` will send the combination of `id and secret` to the `AS` and `AS` will validate them and if valid, issues an `access token`.

![[Pasted image 20250331224815.png]]

```
\> curl \
-u application_id:application_secret \
-H "Content-Type: application/x-www-form-urlencoded" \
-d "grant_type=client_credentials" https://localhost:8085/oauth/token
```

  (-) `-u`: stands for `user` send `HTTP Basic Authentication`. The format is `-u username:password`
  (-) `-h`: stands for `header` add `HTTP request header`. The format is `-H "Header-Name:Value"`
  (-) `-d`: stands for `data` send `request body data`. The format is `-d "key1=value1&key2=value2"`

(+) The `AS` will validate this request and issues an access token in the following HTTP response:

```
{
	"access_token":"de09bec4-a821-40c8-863a-104dddb30204",
	"token_type":"bearer", 
	"expires_in":3599
}
```

(+) Even though we use a `client secret (application_secret)` in the `curl` command to authenticate the `client` application to the `token endpoin`t of the `AS` the `client` application can use `mTLS` instead if stronger authentication is required. In that case, we need to have a `public/private key pair at the client application end`, and the `AS` must trust the issuer of the public key or the certificate.

##### Summary
(+) The `client credentials grant type` is suitable for applications that access APIs and that don’t need to worry about an end user. Because of this, the client credentials grant type is mostly used for system-to-system authentication when an application, a periodic task, or any kind of a system directly wants to access your microservice over OAuth 2.0.

(+) Let’s take a weather microservice, for example. It provides weather predictions for the next five days. If you build a web application to access the weather microservice, you can simply use the `client credentials grant type` because the weather microservice `isn’t interested in knowing who uses your application`. It is concerned with only the application that accesses it, not the end user.


#### B2. RO password grant type
(+) The `RO password grant type` is an extension of the client credentials grant type, but it adds support for `RO authentication` with the `user’s username and password` . This grant type involves all four parties in the OAuth 2.0 flow—`RO`, `Client`, `RS`, `AS`.

(+) The way this work is the `RO` will provide their credentials to the `Client application`, the `Client` will use this information along with its own credentials to make a token request to the `AS`.

![[Pasted image 20250331230905.png]]

```
\> curl \
-u application_id:application_secret \
-H "Content-Type: application/x-www-form-urlencoded" \
-d "grant_type=password&username=user&password=pass" \ https://localhost:8085/oauth/token
```

  (-) `-u`: for `user`, in this context the client application.
  (-) `-H`: for `HTTP Request Header`.
  (-) `-d`: for `data`, in this case there are 3 keys, `grant type`, `user`, and `password`.

(+) In this case, the `AS` validate not only the `client` credentials, but the `user's` credentials also. Upon the successful validation, the `AS` will responds with a valid access token as follows:

```
"access_token":"de09bec4-a821-40c8-863a-104dddb30204", "refresh_token":" heasdcu8-as3t-hdf67-vadt5-asdgahr7j3ty3", "token_type":"bearer", 
"expires_in":3599
```

(+) With the `RO password grant type`, the `RO` (user of the application) needs to provide their username and password to the client application. Therefore, this grant type should be used only with client applications that are trusted by the `AS`.

(+) This model of access delegation is called access delegation with credential sharing. It is, in fact, what OAuth 2.0 wanted to avoid using. The only reason the `RO password grant type` was introduced in the OAuth2.0 specification was to help legacy applications using `HTTP Basic Authentication` migrate to OAuth2.0. Otherwise, avoid using the password grant type where possible.

##### Summary
(+) In this grant type, the `RO` needs to provide the client application with their credentials, the client application will make token request using its own credential, and user's credentials to the `AS`.

(+) It’s also critically important to deal with the user credentials responsibly. Ideally, the client application must not store the end user’s password locally, using it only to get an access token from the authorization server and then forgetting it. 

(+) The access token the client application gets at the end of the password grant flow has a limited lifetime. Before this token expires, the client application can get a new token by using the refresh_token received in the token response from the authorization server. This way, the client application doesn’t have to prompt for the user’s username and password every time the token on the application expires.

##### B3. Refresh token grant type
(+) Not every grant type issues a refresh token along with its access token, including the client credentials grant and the implicit grant. Therefore, the refresh token grant type is a special grant type that can be used only with applications that use other grant types to obtain the access token.

![[Pasted image 20250331235553.png]]

```
\> curl \
-u application_id:application_secret \
-H "Content-Type: application/x-www-form-urlencoded" \
-d "grant_type=refresh_token& refresh_token=heasdcu8-as3t-hdf67-vadt5-asdgahr7j3ty3" \ 
https://localhost:8085/oauth/token
```

##### Summary
(+) As in the earlier cases, the application’s client ID and client secret (application_id and application_secret) must be sent in base64-encoded format as the HTTP Authorization header. You also need to send the value of the refresh token in the request payload (body). Therefore, the refresh token grant should be used only with the applications that can store the client secret and refresh token values secure with out any risk of compromise.

(+) The refresh token usually has a limited lifetime, but it’s generally much longer than the access token lifetime, so an application can renew its token even after a significant duration of idleness.

#### B4. Authorization code grant type
(+) The authorization code grant is used with desktop applications and in web applications (accessed via a web browser) or native mobile applications that are capable of handling HTTP redirects.

(+) In the authorization code grant flow, the `client application` first initiates an authorization code request to the `AS`. This request provides the client ID of the application and a redirect URL to redirect the user when authentication is successful

![[Pasted image 20250401000811.png]] 