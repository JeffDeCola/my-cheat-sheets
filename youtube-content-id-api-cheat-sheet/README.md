# YOUTUBE CONTENT ID API CHEAT SHEET

`youtube-content-id` _is an api for YouTube's Rights Management System._

[GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/)

## WHAT IT CONTENT-ID

Content owners or Administartors will have access to their assets:

* Metadata
* Ownership Information
* Policy Information
* Create New Assets
* Mannage Assets
* Claim Content
* Upload Videos

By combining `Youtube Content ID` with `Youtube Data API` and `Youtube Player API`,
you can enable your app to updaload, manage and play videos.

## YOUTUBE CONTENT PARTNER - ALLOWS ACCESS

You need to be a Youtube Content Partner.

If you are a partner you can access the Content ID Online Interface
[here](https://www.youtube.com/content_id?o=U).

You will also see it listed as a google's API service manager
[here](https://console.developers.google.com/apis/dashboard)

## AUTHENTIFICATION - CONTENT-ID USES OAuth 2.0

Google uses many ways to authentificate, such as API Key,  Service Account or
OAuth 2.0 Client ID.

Content ID uses OAuth 2.0 Client ID for authentification.

To create Create a `OAuth 2.0 Client ID` goto credentials page
[here](https://console.developers.google.com/projectselector/apis/credentials)
and select create credentials.

The Sellect `other` and create.  You now have a client ID.

## OAuth 2.0 AVAILABLE FLOWS

* Server-Side Flow
* Client-Side Flow
* Installed App Flow
* Device Flow
* Service Account Flow (This is what we'll use)

## OAuth 2.0 SERVICE ACCOUNT FLOW - To AUTHORISE API CALLS

A Service Account is a Google Account that can be used to
authorize API calls without any user interaction. It's Server to Server.

For the OAuth 2.0 flow for Service Accounts, the steps for authenticating
and authorizing access to API resources do not take place in a web browser

Your application can use the OAuth 2.0 flow for Service Accounts
for the YouTube Content ID API if it is authenticating and authorizing
requests as a YouTube content owner.

### GET AN OAuth 2.0 ACCESS TOKEN for SERVICE ACCOUNTS

You can use the client libraries [here](https://github.com/google/google-api-go-client)
to implement OAuth 2.0 in your application.

* Step 1A. Create a JWT (JSON WEB TOKEN).

Create a JWT which is used to get an access token.

* Step 1B. Request Access Token

* Step 1C. Extract Access Token

This token is good for one hour.

### USE ACCESS TOKEN IN API REQUEST

* Step 2. Making API Request

After Getting access token for your user...?????

### REFRESH ACCESS TOKEN

?????










