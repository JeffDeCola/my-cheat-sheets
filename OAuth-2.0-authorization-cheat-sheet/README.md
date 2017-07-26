# OAuth 2.0 AUTHORIZATON CHEAT SHEET

`OAuth 2.0 Authorization` _is ???????????????????????????????._

[GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/)

## CREATE OAuth 2.0 CLIENT ID

To create Create a `OAuth 2.0 Client ID` goto credentials page
[here](https://console.developers.google.com/projectselector/apis/credentials)
and select create credentials.

The Sellect `other` and create.  You now have a client ID.

## OAuth 2.0 AVAILABLE FLOWS

1. Server-Side Flow
2. Client-Side Flow
3. Installed App Flow
4. Device Flow
5. Service Account Flow (This is the scope of this cheat sheet)

## OAuth 2.0 SERVICE ACCOUNT FLOW - To AUTHORISE API CALLS

A Service Account is a Google Account that can be used to
authorize API calls without any user interaction. It's Server to Server.

For the OAuth 2.0 flow for Service Accounts, the steps for authenticating
and authorizing access to API resources do not take place in a web browser

Your application can use the OAuth 2.0 flow for Service Accounts
for the YouTube Content ID API if it is authenticating and authorizing
requests as a YouTube content owner.

## STEP 1 - CREATE JWT (JSON WEB TOKEN)

You can use the `golang/oauth2` client libraries
[here](https://github.com/golang/oauth2/)
to implement OAuth 2.0 in your application.

Create a JWT which is used to get an access token.

## STEP 2 - REQUEST ACCESS TOKEN

## STEP 3 - RECEIVE AND EXTRACT ACCESS TOKEN

This token is good for one hour.

## STEP 4 - USE ACCESS TOKEN TO CALL API

After Getting access token for your user...?????

You can use the `google/google-api-go-client` client libraries
[here](https://github.com/google/google-api-go-client)
to implement OAuth 2.0 in your application.

### REFRESH ACCESS TOKEN

?????