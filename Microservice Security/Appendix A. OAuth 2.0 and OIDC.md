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

##### Why can't legacy system handle normal OAuth2.0 flow

###### 1. No Web Browser Integration
**Browsers** have built-in mechanisms to **automatically follow redirects** by reading the `Location` header in the HTTP response. This is part of their default functionality.

 **Legacy CLI tools** (like `curl` or old terminal-based apps) or **desktop apps** that aren’t browsers generally **don’t automatically follow redirects**. Instead, they will simply receive the HTTP response with the redirect status code and the Location header and stop at that point. Without further instructions, they don't make the new request to the `Location`.

###### Example:
 If you use `curl` on the command line and it encounters a redirect (e.g., `301 Moved Permanently`), by default it will **not** automatically follow the redirect. You need to tell it explicitly to follow redirects using the `-L` flag:
```
curl -L http://example.com
```

###### 2. No Handling of State and Cookies
Browsers handle **state** (e.g., cookies, sessions) seamlessly, which is crucial for OAuth flows. When a browser is redirected, it can carry over session information, handle authentication cookies, and securely maintain the context between requests.

Legacy systems, especially those that are **not web-based**, **don’t have this mechanism**. For example, in a **CLI application**, there’s no native way to store and send cookies or manage sessions across requests automatically. This makes it difficult to **maintain the flow of authentication** after a redirect (as needed in OAuth flows).

###### **3. No User Interface for User Interaction**:
Redirection in OAuth involves redirecting the user to an **authorization server** (such as a login page) where the user enters their credentials and consents to access. The browser manages the UI, showing the login screen and handling redirects back to the app with the authentication token or code.
  
**CLI tools** or **desktop apps** typically have no **GUI** to show a login screen. In OAuth, the user’s involvement in granting permissions is required. For legacy systems, without the ability to open a web page or interact via a browser, the **redirection and user input handling** becomes impossible.


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

`(1)`: The first step of the client application is to initiate the authorization code request. The HTTP request to get the authorization code looks like the following:

```
GET https://localhost:8085/oauth/authorize? 
	response_type=code& 
	client_id=application_id&
	redirect_uri=https%3A%2F%2Fweb.application.domain%2Flogin
```

(+) The request caries the `client_id`, `response_type`, `redirect_uri`. The `response_type` indicates to the `AS` that an authorization code is expected as the response to this request. 
==> This authorization code is provided as a query parameter in an HTTP redirect on the provided `redirect_uri`. The `redirect_uri` is the location to which the `AS` should redirect the browser (user agent) upon successful authentication.

`(2)`: In `HTTP`, a redirect happens when the server sends a response code of 302. The response would contain an `HTTP header` named `Location`, and the value of the `Location` header would bear the URL to which the browser should redirect (similar to what the client app has registered with the `AS`).

```
Location: https://web.application.domain/login?code=hus83nn-8ujq6-7snuelq
```

`(3)`: Upon receiving this `authorization request`, the authorization server first validates the `client ID` and the `redirect_uri`; if these parameters are valid, it presents the user with the login page of the authorization server.

`(4)`: The user needs to enter their username and password on this login page. When the username and password are validated, the authorization server issues the `authorization code` and provides it to the `user agent` via an HTTP redirect.

```
https://web.application.domain/login?code=hus83nn-8ujq6-7snuelq
```

#### NOTE:
(+) Because the authorization code is` provided to the user agent via the redirect_uri`, it must be passed over `HTTPS`. Also, because this is a browser redirect, the value of the `authorization code is visible to the end user`, and also may be logged in server logs.

==> To reduce the risk that this data will be compromised, the authorization code usually has a` short lifetime` (no more than 30 seconds) and is a `one-time-use code`. If the code is used more than once, the authorization server revokes all the tokens previously issued against it.

`(5)`: Upon receiving the authorization code, the client application` issues a token request to the authorization server`, requesting an `access token` in exchange for the authorization code.

```
\> curl \
-u application1:application1secret \
-H "Content-Type: application/x-www-form-urlencoded" \
-d "grant_type=authorization_code& 
	code=hus83nn-8ujq6-7snuelq&	
	redirect_uri=https%3A%2F%2Fweb.application.domain%2Flogin" \   
https://localhost:8085/oauth/token
```

`(6)`: Upon validation of these details, the authorization server issues an access token to the client application in an HTTP response:

```
"access_token":"de09bec4-a821-40c8-863a-104dddb30204", "refresh_token":" heasdcu8-as3t-hdf67-vadt5-asdgahr7j3ty3", "token_type":"bearer", 
"expires_in":3599
```

#### Summary
(+) This grant type is suitable to provide user credentials for web, mobile, and desktop applications that you don’t fully trust, as RO's credentials are only provided to the `AS`.

(+) A client application that uses this grant type needs to have some prerequisites to use this protocol securely. Because the application needs to know and deal with sensitive information, such as the client secret, refresh token, and authorization code, it needs to be able to store and use these values with caution. It needs to have mechanisms for encrypting the client secret and refresh token when storing and to use HTTPS.

#### B5. Implicit grant type
(+) The `implicit grant type` is similar to the `authorization code grant type`, but it doesn’t involve the intermediary step of getting an authorization code before getting the access token. Instead, the authorization server `issues the access token directly in response` to the implicit grant request

![[Pasted image 20250405161941.png]]

`(1)`: With the implicit grant type, when the user attempts to log in to the application, the cli ent application initiates the login flow by creating an implicit grant request. This request should contain the client ID and the redirect_uri:

```
GET https://localhost:8085/oauth/authorize? 
	response_type=token& 
	client_id=application_id& 
	redirect_uri=https%3A%2F%2Fweb.application.domain%2Flogin
```

(+) As you can see in the HTTPS request, the difference between the authorization code grant’s initial request and the implicit grant’s initial request is the fact that the `response_type` parameter in this case is `token`, while authorization code is `code`.

`(2)`: When the user has consented to the required scopes, the authorization server issues an access token and provides it to the user agent on the redirect _uri itself as a URI fragment. The following is an example of such a redirect:

```
https://web.application.domain/login#access_token=jauej28slah2& expires_in=3599
```

`(3)`: When the user agent (web browser) receives this redirect, it makes an HTTPS request to the web.application.domain/login URL. Because the access_token field is provided as a URI fragment (denoted by the # character in the URL), that particular value doesn’t get submitted to the server on web.application.domain. Only the authorization server that issued the token and the user agent (web browser) get to know the value of the access token

#### Summary
The implicit grant type doesn’t require your client application to maintain any sen sitive information, such as a client secret or a refresh token. This fact makes it a good candidate for use in SPAs, where rendering the content happens on web browsers through JavaScript. These types of applications execute mostly on the client side (browser); therefore, these applications are incapable of handling sensitive informa tion such as client secrets.

#### C. Scopes bind capabilities to an OAuth2.0 access token
(+) Each access token that an authorization server issues is associated with one or more scopes. `A scope defines the purpose of a token`. A token can have more than one purpose; hence, it can be associated with multiple scopes. A scope defines what the client application can do at the resource server with the corresponding token.

(+) When a client application requests a token from the authorization server, along with the token request, it also specifies the scopes it expects from the token.

![[Pasted image 20250405163052.png]]

(+) That doesn’t necessarily mean the authorization server has to respect that request and issue the token with all requested scopes. An authorization server can decide on its own, also with the resource owner’s consent, which scopes to associate with the access token

##### C1. Types of access tokens

(+) `Reference Token`: is just a string ,and only the issuer of the tokens knows how to validate it (meaning the AS), so when the resource server gets a reference token, it has to talk to the authorization server all the time to validate the token.

(+) `Self-contained Token`: the resource server can validate the token itself; there’s no need to talk to the authorization server. A self contained token is a signed JWT or a JWS.

#### D. What is OpenID Connect?
(+) `OpenID Connect` is built on top of OAuth 2.0 as an additional identity layer. It uses the concept of an `ID token`. An ID token is a `JWT that contains authenticated user information`, including user claims and other relevant attributes.

(+) When an `authorization server` issues an ID token, it signs the contents of the JWT (`a signed JWT is called a JWS, or JSON Web Signature`), using its private key. Before any application accepts an ID token as valid, it should verify `its contents by validating the signature of the JWT`.

#### NOTE:
(+) `ID token` is consumed by an `application to get information`, such as a `user’s username, email address, phone number`, and so on. An `access token` is a `credential used by an application to access a secured API` on behalf of an end user or just by being itself. OAuth 2.0 provides only an access token, whereas OpenID Connect provides both an access token and an ID token

```
"iss":"http://server.example.com", 
"sub":"janedoe@example.xom", 
"aud":"8ajduw82swiw", 
"nonce":"82jd27djuw72jduw92ksury", 
"exp":1311281970, 
"iat":1311280970, 
"auth_time":1539339705, 
"acr":"urn:mace:incommon:iap:silver", 
"amr":"password", 
"azp":"8ajduw82swiw"
```

(+) `iss`: The identifier of the issuer of the ID token (usually, an identifier to rep resent the authorization server that issued the ID token). 
(+) `sub`: The subject of the token for which the token was issued (usually, the user who authenticated at the authorization server). 
(+) `aud`: The audience of the token; a collection of identifiers of entities that are supposed to use the token for a particular purpose. It must contain the OAuth 2.0 client_id of the client application, and zero or more other identifiers (an array). If a particular client application uses an ID token, it should validate whether it’s one of the intended audiences of the ID token; the client application’s client_id should be one of the values in the aud claim.
(+) `iat`: The time at which the ID token was issued. 
(+) `exp`: The time at which the ID token expires. An application must use an ID token only if its exp claim is later than the current timestamp

(+) An ID token usually is obtained as part of the access token response. OAuth 2.0 pro viders support various grant types for obtaining access tokens, as we discussed in sec tion A.5. An ID token usually is sent in the response to a request for an access token by using a grant type. You need to specify openid as a scope in the token request to inform the authorization server that you require an ID token in the response:

```
GET https://localhost:8085/oauth/authorize? 
	response_type=code& 
	scope=openid& 
	client_id=application1& 
	redirect_uri=https%3A%2F%2Fweb.application.domain%2Flogin
```

The ID token is sent in the response: 

```
"access_token": "sdfj82j7sjej27djwterh720fnwqudkdnw72itjswnrlvod92hvkwyfp", "expires_in": 3600, "token_type": "Bearer", 
"id_token": "sdu283ngk23rmas….."
```