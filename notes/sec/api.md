# Authorization header

- to pass a token or credentials to the API host (provider):

```
Authorization: <type> <token/credentials>
```

Types

- Basic - uses base64-encoded credentials
- Bearer - uses an API token
- AWS-HMAC-SHA256 - uses access key and secret key

# Authentication

- process of proving and verifying an identity

## Basic authentication

- completely depends on other security controls
- credentials passed in plaintext or base64-encoded (not encrypted!) to save space
```
$ echo "username:password"|base64
dXNlcm5hbWU6cGFzc3dvcmQK
$ echo "dXNlcm5hbWU6cGFzc3dvcmQK"|base64 -d
username:password
```
- credentials may be required for every request or just once in exchange for API key or token

## API keys

- unique strings generated by providers for approved consumers
- passed in query string:
```
/api/v1/users?apikey=ju574n3x4mp134p1k3y #gitleaks:allow
```
- header:
```
"API-Secret": "17813fg8-46a7-5006-e235-45be7e9f2345" #gitleaks:allow
```
- or as a cookie:
```
Cookie: API-Key=4n07h3r4p1k3y
```
- process of acquiring API key depends on the provider, e.g. you need to register with email
- more secure than basic-auth because:
    - usually longer, more complex and random -> very difficult to brute force
    - can have expiration date to limit the time of validity

## JSON Web Tokens

1. Consumer authenticates with username/password
2. Provider generates and sends back JWT
3. Consumer adds JTW to the `Authorization` header of all requests

- three parts separated by `.`:

1. Header - algorithm used to sign the payload
2. Payload - username, timestamp, issuer
3. Signature - encrypted message used to validate the token

- generally secure but can be implemented insecurely

## HMAC

- hash-based message authentication code
- primary authn method used by AWS
- security depends on the consumer and provider keeping the secret key private

1. provider creates secret key and shares it with consumer
2. consumer applies HMAC hash function to each API request data and secret key
3. provider calculates the hash and see if it matches

- algorithms from weaker to stronger:
```
HMAC-MD5
HMAC-SHA1
HMAC-SHA256
HMAC-SHA512
```
- optimize for performance vs security

## OAuth 2.0

- standard that allows different services to access each other's data (often using APIs)
- E.g. you want share tweets on LinkedIn:
1. Authorize LinkedIn (via web GUI) to access your tweets.
2. Twitter generates a limited, time-based access token for LinkedIn.
3. LinkedIn provides the token to Twitter to post on your behalf.
- one of the most trusted forms of API authorization
- but also expands attack surface
- poor implementation can expose API providers to token injection, authorization code reuse, cross-site request forgery, invalid redirection, and phishing attacks

## No authentication

- plenty of instances to have no authn at all
- if API does not handle sensitive data and provides public info only

# Common vulnerabilities

## Information disclosure

- useful to start gaining access to an API
- API responses
- public sources like: code repos, search results, news, social media, target's website, and public API directories
- unknowingly sharing all users' info
- verbose messaging like: "user does not exist", "incorrect password"
- tech stack components vendors and versions

## Broken object level authorization (BOLA)

- consumer can access resources that they shouldn't, e.g. 
```
GET /api/resouce/**1**
GET /users/account/find?user_id=**15**
POST /company/account/**Apple**/balance
POST /admin/pwreset/account/**90**
```

## Broken user authentication

Missing or incorrect implementation of
- user registration and password reset
- token generation (low randomness)
- token handling (hardcoded tokens)
- access to resources w/o authn

## Excessive data exposure

- API endpoint reports more info than needed
- this vulnerability bypasses every security control :-)

## Lack of resources and rate limiting

- lead to DoS
- many API providers monetize their APIs via rate limiting

Test
1. rate limiting works (HTTP 429)
2. is enforced

## Broken function level authorization (BFLA)

- user of one role or group can access API functionality of another role or group
- similar to BOLA but about actions instead of resources

Test
1. find administrative API docs (or discover endpoints)
2. send requests for admin functionality as non-admin

## Mass assignment

- occurs when more parameters are allowed than needed
```
{
"User": "John",
"Password": "secr3t",
"isAdmin": true <---
}
```

Test
1. find interesting parameters is docs
2. add them to a request (you can do fuzzing)

## Security misconfigurations

- lack of input sanitization -> upload of malicious payloads to the server
- API provider sends headers with instructions that might be misused, e.g.:
```
X-Powered-By: VulnService 1.11  <-- search for exploits
X-XSS-Protection: 0             <-- mount XSS attack
X-Response-Time: 566.43         <-- time requests
```
- no HTTPS used <- MITM attack
- unnecessary methods used <- bigger attack surface

Tools: Nessus, Qualys, OWASP ZAP, Nikto

## Injections

- no input sanitization
- if you get response directly from DB -> [No]SQL injection vulnerability (`"' OR 1=0--"`)

## Improper assets management

- exposing APIs that are retired or still in development -> bigger attack surface

Tests
- check changelogs, repo version history
- search for /v1/, /v2/, /v3/, ...
- search for /alpha/, /beta/, /test/, /uat/, /demo/, ...
- guessing, fuzzing, or bruteforcing requests

## Business logic vulnerabilities

- intended features that attackers can use maliciously
- come from an assumption that 
    - consumers will follow directions 
    - will use API only in certain way
    - wiil only use browser (not Burp or Postman :-) to interact with the wep app

Check API docs for
- "Only use feature X to perform function Y."
- "Do not do X with endpoint Y."
- "Only admins should perform request X."

---

Source: Hacking APIs (2022)