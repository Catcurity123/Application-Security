#### C.1 What is JSON Web Token
(+) A JWT (jot) is a container that carries different types of assertions or claims from one place to another in a cryptographically safe manner. 

(+) `An assertion` is a strong statement about someone or something issued by an entity. 

==> Imagine the state can create a JWT with your personal info, which include your name, address, eye color, hair color, gender,... All these items are attributes, or `claims` about you and are also known as `attribute assertions.`

(+) Anyone who gets this JWT can decide whether to accepts what's in it as true, based on the level of trust they have in the issuer of the token. 

(+) But before accepting it, how do you know who issued it? The issuer of a JWT signs it by using the issuer's private key.

![[Pasted image 20260116141828.png]]

(+) In addition to attribute assertions, a JWT can carry AuthN and AuthZ assertions.

#### C.2 What does a JWT look like
![[Pasted image 20260116142001.png]]

(+) What we see above is a `JSON Web Signature (JWS)`, the JWS, which is the most commonly used format of a JWT, has three parts with `a dot (.) `separating them.
	1. The first part is known as the JSON Object Signing and Encryption (JOSE) header.
	2. The second part is the claim set, or body (or payload)
	3. The third part is the signature.

###### The JOSE header
(+) The JOSE header is a base64url-encoded JSON object, which expresses the metadata related to the JWT, such as the algorithm used to sign the message. The decoded version is:

```
{
	"alg": "RS256",
}
```

###### The claim set
(+) The JWT claim set is a base64url-encoded JSON object, which carries the assertions. The decoded version is:

```
{
	"sub": "peter", 
	"aud": "*.ecomm.com", 
	"nbf": 1533270794, 
	"iss": "sts.ecomm.com", 
	"exp": 1533271394, 
	"iat": 1533270794, 
	"jti": "5c4d1fa1-7412-4fb1-b888-9bc77e67ff2a"
}
```

(-) `iss`: this is the attribute in the JWT claim set carries an identifier corresponding to the `issuer`, or asserting party, of the JWT.

(-) `sub`: the subject is the owner of the JWT - or in other words, the JWT carries the claims about the subject.

(-) `aud`: the audience claim is about the intended recipient of the token. The value of the `aud` attribute can be any string of a URI that's known to the microservice or the recipient of the JWT.
==> Each microservices must check the `aud` parameter before accepting any JWT as valid. If a `foo` microservice received an `aud` as `bar.ecomm.com` it should be rejected.

(-) `exp`: this is an attribute that expresses the time of expiration in seconds, which is calculated from 1970-01-01T0:0:0Z as measured in Coordinated Universal Time (UTC). Any recipient of a JWT must make sure that the time represented by the `exp` attribute is not in the past when accepting a JWT. Or in other words, the token has not expired.

(-) `iat`: is an attribute in the JWT claim set expresses the time when the JWT was issued. 

(-) `nbf`: is an attribute in the JW claim set expresses that the token shouldnt be processed or accept as valid `not before` the time.
==> The lifetime of a JWT is normally the difference between the `exp` and the `iat` claim, but if there is a `nbf` claim, the lifetime of JWT is calculated as the difference between `exp` and `nbf`

(-) `jti`: is the JWT identifier for the token. It is a unique identifier for each token. The token issuer should not issue two JWTs with the same `jti`. If the recipient of the JWT accepts tokens from multiple issuers, a given `jti` will be unique only along with the corresponding issuer identifier.

###### The Signature
(+) The signature is created by concatenating the header and payload and the dot, then hash it using the algorithm in the header.

#### C3. JSON Web Signature
(+) A JWS can be serialized in two formats `Compact Serialization` or `JSON Serialization`. Not all JWS is a JWT, `A JWS becomes a JWT only when it follows compact serialization and carries a JSON object`.

(+) Under JWT terminology, we call this a claim set. The below figure shows a `compact-serialized` JWS or a JWT.

![[Pasted image 20260116145644.png]]

(+) With `JSON serialization`, the JWS is represented as a JSON payload, it is not called a JWT. The payload parameter in the `JSON-serialized JWS` can carry any value. The message below is a JSON message with all its related metadata.

![[Pasted image 20260116150111.png]]

(+) Unlike in a JWT, a JSON serialized JWS can carry multiple signature corresponding to the same payload.

#### C4. JSON Web Encryption
(+) So we know established that JWT is a `compact-serialized JWS`. It can also be a `compact-serialized JSON Web Encryption (JWE)`. Like JWS, a JWE represents an encrypted message using compact serialization. or JSON serialization. 
==> A JWE is called a JWT only when compact serialization is used. 

==> JWS addresses the `integrity` and `non-repudiation` aspects of the data contained in it, through the hash signature based on the payload).

(+) A compact-serialized JWE has five parts; each part is base64url-encoded and separated by a dot. 
![[Pasted image 20260116150929.png]]

(+) With JSON serialization, the JWE is represented as a JSON payload, it isn't called a JWT. The `ciphertext` attribute in the `JSON-serialized JWE` carries the encrypted value of any payload, which can be JSON, MXL or even binary.

![[Pasted image 20260116151050.png]]