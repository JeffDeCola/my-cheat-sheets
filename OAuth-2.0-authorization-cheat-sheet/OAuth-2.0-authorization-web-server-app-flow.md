# OAuth 2.0 WEB SERVER APP AUTHERIZATION CHEAT SHEET

`OAuth 2.0 Web Server Side Authorization Flow` _allows users to authenticate themselfs
in order for an APP to gain access to their information via an API._

[GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/)

## HIGH-LEVEL VIEW

The following diagram illustrates how everything fits for
Web Server-Side Flow. To accesss a users info they must be
in your g-suite domain.

![IMAGE - OAuth 2.0 Web Server APP Authorization Flow - IMAGE](OAuth-2.0-web-server-app-authorization-flow.jpg)

## STEP 1 - CREATE OAuth 2.0 CLIENT ID & SECRET

To create Create a `OAuth 2.0 Client ID` goto credentials page
[here](https://console.developers.google.com/projectselector/apis/credentials)
and select create credentials.

Create a OAuth 2.0 client IDs for a Web Application.

Origin URI,

[http://127.0.0.1:3000](http://127.0.0.1:3000)

Redirect URI,

[http://127.0.0.1:3000/GoogleCallback](http://127.0.0.1:3000/GoogleCallback)

You will now have a Client ID and a Secret.

## STEP 2 - ???

You can use the `golang/oauth2` client libraries
[here](https://github.com/golang/oauth2)
to implement OAuth 2.0 in your application.

## STEP 3 - ????

## STEP 4 - USE ACCESS TOKEN TO CALL API

After Getting access token for your user...?????

You can use the `google/google-api-go-client` client libraries
[here](https://github.com/google/google-api-go-client)
to use APIs in your application.

### REFRESH ACCESS TOKEN

?????