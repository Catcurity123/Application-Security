(+) This part of the book takes you through securing a `microservice at the edge` (or at the entry point) in a typical microservices deployment. In most cases, `micro services are behind a set of APIs` that is exposed to the outside world via `an API gateway`.

(+) The `API Gateway pattern` is mostly about `edge security`, while the `Service Mesh pattern` deals with `service-to-service security`. Or, in other words, the API Gateway deals with `north/south traffic`, while the Service Mesh deals with `east/west traffic`.

(+) `Edge security` is about protecting a set of resources (for example, a set of microservices) at the entry point to the deployment, at the API gateway. The API gateway is the only entry point to our microservices deployment for requests originating from out side.

(+) In the `Service Mesh pattern`, the architecture is much more decentralized. Each microservice has its own policy enforcement point much closer to the service—mostly it is a proxy, running next to each microservice

(+) The API gateway is the `centralized policy enforcement point to the entire microservices deployment`, while in a service mesh, `a proxy running along with each microservice provides another level of policy enforcement at the service level`.

#### 1. The need for an API gateway in a microservice deployment
(+) Having an API gateway in a microservices deployment is important. The API gateway is a crucial piece of infrastructure in our architecture, since it plays a critical role that helps us clearly `separate the functional requirements from the nonfunctional ones`.

(+) The API gateway is the entry point to the microservices deployment, which screens all incoming messages for security and other QoS features.

(+) It acts as the front door to Netflix’s server infrastructure, handling traffic from Netflix users around the world. In figure 3.1, Zuul is used to expose the Order Processing microservice via an API. We do not `expose Inventory and Delivery microservices from the API gateway`, because external applications don’t need access to those.

![[Pasted image 20250406151130.png]]

##### 1.1 Decoupling security from the microservices
(+) One key aspect of microservices best practices is the single *responsibility principle.* 

(+) In the example in chapter 2, the secured `OrderProcessing` microservice was implemented in such a way that it had to talk the the AS and validate the access tokens it got from client applications, in addition to the core business functionality of processing orders. 

(+) Executing all these steps becomes a problem because the microservice loses its atomic characteristics by performing more operations than it’s supposed to. 

(+) The coupling of security and business logic introduces` unwanted complexity and maintenance overhead to the microservice`. For example, making changes in the security protocol would require changes in the microservice code, and scaling up the micro service would result in more connections to the authorization server.

(+) This helps in 2 main ways:

###### CHANGES IN THE SECURITY PROTOCOL REQUIRE CHANGES IN THE MICROSERVICE
(+) This unwanted overhead compromises your agility in designing, developing, and deploying your microservice. So it is best to decouple this from the microservice itself.

###### SCALING UP THE MICROSERVICE RESULTS IN MORE CONNECTIONS TO THE AUTHORIZATION SERVER
(+) In certain cases, you need to run more instances of the microservice to cater to rising demand. Because each microservice talks to the authorization server for token validation, scaling your microservice will also increase the number of connections to the authorization server.

![[Pasted image 20250406152903.png]]

##### 1.1.2 The inherent complexities of microservice deployments make them harder to consume
(+) A microservices deployment typically consists of many microservices and many inter actions among these microservices.

![[Pasted image 20250406153000.png]]

(+) An API gateway solution, which usually comes as part of API management soft ware, can bring consistency to the interfaces that are being exposed to the consuming applications.

##### 1.1.3 The rawness of microservices does not make them ideal for external exposure

(+) Microservices can be as granular as they need to be. Suppose that you have two opera tions in your Products microservice: one for retrieving your product catalog and another for adding items to the catalog, meaning `GET /products` and `POST /products`

(+) In practice there could be more operation, therefore, we can decide to implement GET and POST on 2 different microservices. So that they can scale independently and failure in 1 microservice does not affect the other.

==> From a consuming point of view, however, it would be odd for the consuming applications to have to talk to two endpoints (two APIs) for the add and retrieve operations. A strong REST advocate could argue that it makes more sense to have these two operations on the same API (same endpoint).

(+) The API Gateway architectural pattern is an ideal solution to this problem. It pro vides the consuming application a single API with two resources (GET and POST). Each resource can be backed by a microservice of its own, providing the scalability and robustness required by the microservices layer.

![[Pasted image 20250406153710.png]]

#### 1.2 Security at the edge
(+) In most cases, these API gateways use OAuth 2.0 as the security protocol to secure the APIs they expose at the edge.

##### 1.2.1 Delegating access
(+) A microservice is exposed to the outside world as an API. Similar to microservices, even for APIs the audience is a system that acts on behalf of itself, or on behalf of a human user or another system. It’s unlikely (but not impossible) for human users to interact directly with APIs.

![[Pasted image 20250406222346.png]]
(+) A user (resource owner) should be allowed to perform only the actions on microser vices that they are privileged to perform. The data that the user retrieves from the microservice, or updates via the microservice, should be only the data that they are entitled to receive or update

(+) As a result, the application has a responsibility to deal with the delegated rights appropriately.

##### 1.2.3 Why not basic authentication to secure APIs
(+) Basic authentication allows a user (or a system) with a valid username and password to access a microservice via an API. In fact, basic authentication (or basic auth) is a standard security protocol introduced with HTTP/1.0 in RFC 1945.

(+) It allows you to pass the `base64-encoded username and password`, in the HTTP Authorization header, along with a request to an API. This is considered insecure because:
	`1.` `The username and password are static, long-live credentials`. The client will have to retain this credentials to access the microservice, The longer this information is retained, the higher the chance of compromise. And because these credentials almost never change, a compromise of this information could have severe consequences.
	`1.` `The username and password are static, long-live credentials`. After an application gets access to the username and password of a user, it can do everything that user can do with the microservice.

##### 1.2.4 Why not mutual to secure APIs?
(+) Mutual Transport Layer Security is a mechanism by which a client application verifies a server and the server verifies the client application by exchanging respective certificates and proving that each one owns the corresponding private keys

(+) `mTLS` solves one of the problems with basic authentication by having a lifetime for its certificates. The certificates used in mTLS are time-bound, and whenever a certifi cate expires, it’s no longer considered valid. Therefore, the risk of applying `mTL is limited by its lifetime`.

==> Then again, unlike basic authentication (where you send your password over the wire), when you use mTLS, the corresponding private key never leaves its owner or is never passed over the wire, hence the `advantage of mTLS over basicAuth`.

(+) However, mTLS doesn’t provide a mechanism to `represent the end user who uses the corresponding application`. You can use mTLS to authenticate the client application that talks to the microservice, but it `does not represent the end user`. Therefore, mTLS is mostly used to secure communication between a client application and a microservice, or `communications among microservices`. 

##### 1.2.5 Why OAuth2.0
(+) `Who`: Ensure that only permitted entities are granted access to your resources 
(+) `What purpose`: Ensure that the permitted entities can perform only what they’re allowed to perform on your resources
(+) `How long`: Ensure that access is granted for only the desired period

(+) You may have a Netflix account, and to view the trending movies on Netflix on your smart TV, you need to `delegate access from your Netflix account to your smart TV.` Delegation is a key requirement in securing microservices—and out of all security protocols, `OAuth 2.0, which is designed for access delegation, fits best in securing microservices at the edge`.

#### 1.3 Setting up an API gateway with Zuul
(+) This is the command to send POST request to API Gateway, we are sending the request to port 9090, and the API Gateway will map this port to the OrderProcessing Microservice.

```
curl http://localhost:9090/retail/orders/f2389beb-bef3-409f-a2b0-cb657acda076 | jq
```

(+) The following line you find there instructs the Zuul proxy to route requests received on `/retail` to the server running on http://localhost:8080: 

```
zuul.routes.retail.url=http://localhost:8080
```

![[Pasted image 20250406231645.png]]

##### 1.3.1 Enforcing OAuth2.0-based security at the Zuul gateway
(+) In a typical production deployment architecture, the` authorization server is deployed inside the organization’s network`, and only the required endpoints are exposed externally. Usually, the API gateway is the only component that’s allowed access from outside; everything else is restricted within the local area network of the organization.

![[Pasted image 20250406232201.png]]

(+) The following is the endpoints specification for Zuul API Gateway:

```
zuul.routes.retail.url=http://localhost:8080
zuul.routes.token.url=http://localhost:8085
zuul.sensitiveHeaders=
ribbon.eureka.enabled=false
server.port=9090
authserver.introspection.endpoint=http://localhost:8085/oauth/check_token
```

(+) Once the client application gets an access token from the token endpoint of the authorization server, the client application accesses the Order Processing microservice via the Zuul gateway with this token.

(+) The purpose of exposing the Order Processing microservice via the Zuul gateway is to make the `gateway enforce all security-related policies while the Order Processing microservice focuses only on the business logic` it executes.

###### Enforcing token validation at the Zuul Gateway
(+) In the previous example, the Zuul gateway talked to the authorization server (token issuer) to validate the token it got from the curl client application, this is called `token introspection`.

![[Pasted image 20250407001315.png]]

(+) To do this in Zuul, you use a request filter, which intercepts requests and performs various operations on them. A filter can be one of four types:
`1. Prerequest filter`: Executes before the request is routed to the target service 
`2. Route filter`: Can handle the routing of a message 
`3. Post-request filter` : Executes after the request has been routed to the target service 
`4. Error filter` : Executes if an error occurs in the routing of a request

(+) The run method of this class contains the logic related to introspecting the token through the authorization server. If you look at that method, you’ll notice that a few validations are performed on the `Zuul gateway itself to check whether the client is sending the access token in an HTTP header named Authorization` and whether the header is received `in the correct format`.

```
//If the authorization server doesn't respond with a 200.
if (responseCode != 200) {
  log.error("Response code from authz server is " + responseCode);
  handleError(requestContext);
}
```

(+) If the server doesn’t respond with 200, the authentication has failed. The authentica tion could have failed for many reasons. The token may have been incorrect, the token may have expired, the authorization server may have been unavailable, and so on.

###### OAuth2.0 Token Introspection Profile
(+) After receiving the `access token` from the client, the `APIGW` will ask the `AS` if the token is correct or not.

```
POST /oauth/check_token 
Content-Type: application/x-www-form-urlencoded 
Authorization: Basic YXBwbGljYXRpb24xOmFwcGxpY2F0aW9uMXNlY3JldA== 

token=626e34e6-002d-4d53-9656-9e06a5e7e0dc& 
token_type_hint=access_token&
```

(+) The `token_type` field indicates to the authorization server whether this token is an access_token or refresh_token. When the authorization server completes the introspection, it responds with a payload similar to this:

```
HTTP/1.1 200 OK 
Content-Type: application/json 
Cache-Control: no-store 

{ 
"active": true, 
"client_id": "application1", 
"scope": "read write", 
"sub": "application1", "aud": "http://orders.ecomm.com"
} 
```

(+) Using the information in the introspection response, the gateway can allow or refuse access to the resources. It can also perform fine-grained access control (authorization) based on the scopes as well as get to know which client application.

###### Self-Validation of Tokens without integrating with an AS
(+) The current method does not scale well as the APIGW is relying on AS to check if the token is valid or not. Look at what an authorization server does when someone asks it to validate a token through an introspection call:

`(1)` It checks to see whether that particular token exists in its token store (data base). This step verifies that the token was issued by itself and that the server knows details about that token. 
`(2)` It checks whether the token is still valid (token state, expiry, and so on). 
`(3)` Based on the outcome of these checks, it responds to the requester with the information discussed under the “OAuth 2.0 token introspection profile” section.

(+) The fundamental problem is not that the `APIGW can not verify the token, but is whether the current token is issued by a trusted AS or not.`

==> `JWTs are designed to solve this problem`

(+) `A JSON Web Signature (JWS) is a JWT signed by the authorization server`. By verifying the sig nature of the JWS, the gateway knows that this token was issued by a trusted party and that it can trust the information contained in the body

###### Pitfalls of self-validating tokens and how to avoid them
(+) The self-validating token mechanism discussed in the preceding section comes at a cost, with pitfalls that you have to be mindful of. If one of these tokens is `prematurely revoked, the API gateway won’t know that the token has been revoked`, because the revocation happens at the `authorization server end, and the gateway no longer communicates with the authorization server for the validation` of tokens.

==> One way to solve this problem is to issue `short-lived JWTs (tokens) to client applications to minimize the period during which a revoked token` will be considered valid on the API gateway. In practice, however, `applications with longer user sessions have to keep refreshing their access tokens` when the tokens carry a shorter expiration.

==> Another solution is for the `authorization server to inform the API gateway` when ever a token has been revoked. The gateway and authorization server can maintain this communication channel via a pub/sub mechanism. This way, `whenever a token is revoked, the gateway receives a message from the authorization server through a message broker`. Then the gateway can maintain a list of revoked tokens until their expiry and before validating a given token check if it exists in the “revoked tokens” list.

![[Pasted image 20250407005214.png]]

(+) Another problem with the self-contained access token is that the `certificate used to verify a token signature might have expired`. When this happens, the gateway can no longer verify the signature of an incoming JWT (as an access token)

==> To solve this problem, you need to make sure that whenever a certificate is renewed, you deploy the new certificate on the gateway.

==> Sometimes, provisioning new certificates can be a little harder when you have a large-scale deployment with certificates having a short lifetime.

==> In that case, you do `not need to provision the token issuer’s certificate to the gateway`, but the `certificate of the certificate authority (CA) corresponding to the token issuer’s certificate`. Then the gateway can fetch the token issuer’s certificate dynamically from an endpoint exposed by the authorization server to do the JWT signature validation, and check whether that certificate is signed by a certificate authority it trusts.

#### 1.4 Securing communication between Zuul and the microservice
(+) We must also have to consider what happens if someone `accesses the microservice directly, bypassing the API gateway layer.` In this section, we discuss how to secure access to your microservice in such a case.

##### 1.4.1 Preventing access through the firewall
(+) First and foremost, you need to make sure that your microservice isn’t directly exposed to external clients, so you need to make sure that it sits behind your organization’s firewall.

(+) Although the API Gateway pattern ensures that no one outside the organization gets access to the microservice, a risk still exists that unintended internal clients may gain access to the microservice.

![[Pasted image 20250407010421.png]]

##### 1.4.2 Securing the communication between the API gateway and microservices by using mTLS.
(+) To make sure that your microservice is secure from internal clients and accessible only via the API gateway, you need to build a mechanism in which the microservice rejects any requests coming from clients other than the API gateway

(+) mTLS verification happens at the Transport layer of the microservice and doesn’t propagate up to the Application layer. Microser vices developers don’t have to write any application logic to handle the client verifica tion, which is done by the underlying Transport-layer implementation.

![[Pasted image 20250407010620.png]]