# OAuth 2.0 AUTHORIZATON CHEAT SHEET

`OAuth 2.0 Authorization` _is ???????????????????????????????._

`OAuth 2.0 Web Server Side Authorization Flow` _allows users to authetificate._

[GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/)

## OAuth 2.0 AVAILABLE FLOWS

1. Server-Side Flow (Scope of this cheat sheet)
1. Client-Side Flow
1. Installed App Flow
1. Device Flow
1. Service Account Flow

The following diagram illustrates how everything fits for
Web Server-Side Flow. To accesss a users info they must be
in your g-suite domain.

![IMAGE - OAuth 2.0 Web Server Authorization Flow - IMAGE](OAuth-2.0-web-server-authorization-flow.jpg)

## STEP 1 - CREATE OAuth 2.0 CLIENT ID & SECRET

To create Create a `OAuth 2.0 Client ID` goto credentials page
[here](https://console.developers.google.com/projectselector/apis/credentials)
and select create credentials.

Create a OAuth 2.0 client IDs for a web application.

You will now have a Client ID and a Secret.

## STEP 2 -

## STEP 3 -

## STEP 4 - USE ACCESS TOKEN TO CALL API

After Getting access token for your user...?????

You can use the `google/google-api-go-client` client libraries
[here](https://github.com/google/google-api-go-client)
to implement OAuth 2.0 in your application.

### REFRESH ACCESS TOKEN

?????