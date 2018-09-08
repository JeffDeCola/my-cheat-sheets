# OAuth 2.0 SERVICE ACCOUNT AUTHERIZATION CHEAT SHEET

`OAuth 2.0 Service Account Authorization Flow` _is using your
service account to access your data or one of your teams data via g-suite._

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## OAuth 2.0 SERVICE ACCOUNT FLOW - TO AUTHORISE API CALLS

A Service Account is a Google Account that can be used to
authorize API calls without ANY user interaction. It's Server-to-Server.

For the OAuth 2.0 flow for Service Accounts, the steps for authenticating
and authorizing access to API resources do not take place in a web browser

Your application can use the OAuth 2.0 flow for Service Accounts
such as the YouTube Content ID API or google Cloud Storage API.

## CREATE OAuth 2.0 CLIENT ID & SECRET

To create an `OAuth 2.0 Client ID` goto credentials page
[here](https://console.developers.google.com/projectselector/apis/credentials)
and select create credentials.

Create a OAuth 2.0 client IDs for a `web application`.

You will now have a Client ID and a Secret.  Obviously, keep this in a safe place.

## HIGH-LEVEL VIEW

The following diagram illustrates how everything fits together. To
access a users info they must be in your g-suite domain.

![IMAGE - OAuth 2.0 Service Account Authorization Flow - IMAGE](../../../docs/pics/OAuth-2.0-service-account-authorization-flow.jpg)
