(+) The fundamental of `OAuth 2.0 is to fix the access delegation problem`, OIDC (OpenID Connect) is an` identity layer built on top of OAuth 2.0`.

#### A. The access delegation problem
(+) If you want someone else to access a resource (a microservice, an API, and so on) on your behalf and do something with the resource, you need to `delegate the corresponding access rights to that person (or thing)`. 
==> If you want a 3rd party to read FB status message, corresponding rights need to be given to that 3rd party.

(+) There are 2 models of access delegation, `credential sharing, and no credential sharing`. The first model is quite risky, as with the credentials, the 3rd party will have authorization to do anything that wee can do with our account.

###### How does OAuth 20 fix the access delegation problem
![[Pasted image 20250315213113.png]]

(+) With OAuth 2.0, the third-party web application` first redirects the user to Facebook` (where the user belongs). Facebook authenticates and `gets the user’s consent to share a temporary token` with a third-party web application, which is `only good enough to read the user’s Facebook status messages` for a limited time. Once the web application `gets the token from Facebook, it uses the token along with the API calls to Facebook`.

==> The temporary token Facebook issues has a limited lifetime and is bound to the Facebook user, the third-party web application, and the purpose. The purpose of the token here is to read the user’s Facebook status messages, and the token should be only good enough to do just that and no more

###### Terminology of OAuth 2.0
(+) `Resource owner`: decides who should have which level of access to the resources owned.  In this case the Facebook User
(+) ``