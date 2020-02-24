# Let's encrypt

The objective of [Let's Encrypt](https://letsencrypt.org) and [ACME protocol](https://ietf-wg-acme.github.io/acme/) is to make it possible to set up an HTTPS server without human interaction.

Two steps are needed:

1. an agent runing on a server needs to prove that is in control of a domain
2. the agent can request, renew or revoke a certificate for the domain

## Domain validation

Let's Encrypt (LE) identifies the agent by a public key. So, the first time the agent interacts with LE it generates a new key pair and proves to LE CA that it is in control of the domain. The CA might give the agent a choice of:

1. provision a DNS record under example.com
2. provision am HTTP resource under a well-known URI on http://example.com

Along with the challenges, the LE CA also provides a *nonce* that the agent must sign with its private key to prove the control of the key pair.

If the signature over the once is valid, and the challenges check out, the agent is said to posses an "authorized key pair".

## Certificate issuance and revocation

Once the agent has an authorized key pair, it can request, renew and revoke the certificates.

# Sources

* [Web Application Security](https://learning.oreilly.com/library/view/web-application-security/9781492053101/) (2020)
* [The Web Application Hacker's Handbook](https://learning.oreilly.com/library/view/the-web-application/9781118026472/) (2011)
