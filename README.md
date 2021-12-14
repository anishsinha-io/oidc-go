# oidc-go

This is an implementation of an OpenID Connect Provider (OP) as defined by [OpenID](https://openid.net/specs/openid-connect-core-1_0)
that is meant to be a full, production ready OP. 

Features:

- Easy to deploy on any platform. Just containerize it and run it in docker or kubernetes on any cloud platform (I personally use Google Cloud)
- Asymmetrically signed tokens. The RS256 algorithm is used to asymmetrically sign all tokens, and the keys are regularly and programatically rotated
- Customize the theme easily by modifying the files in the templates directory


I'll add more to this as I get closer to finishing it.