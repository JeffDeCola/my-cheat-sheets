# YOUTUBE CONTENT ID API CHEAT SHEET

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

```text
*** THIS CHEAT SHEET IS UNDER CONSTRUCTION - CHECK BACK SOON ***
```

_An api for YouTube's Rights Management System._

Table of Contents

* [WHAT IS YOUTUBE'S CONTENT-ID SYSTEM](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/api/youtube-content-id-api-cheat-sheet#what-is-youtubes-content-id-system)
* [YOUTUBE CONTENT PARTNER - ALLOWS ACCESS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/api/youtube-content-id-api-cheat-sheet#youtube-content-partner---allows-access)
* [AUTHENTIFICATION - CONTENT-ID USES OAuth 2.0](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/api/youtube-content-id-api-cheat-sheet#authentification---content-id-uses-oauth-20)

## WHAT IS YOUTUBE'S CONTENT-ID SYSTEM

Content owners or Administrators will have access to their assets:

* Metadata
* Ownership Information
* Policy Information
* Create New Assets
* Manage Assets
* Claim Content
* Upload Videos

Furthermore, by combining `Youtube Content ID` with `Youtube Data API`
and `Youtube Player API`, you can enable your app to upload, manage and play videos.

## YOUTUBE CONTENT PARTNER - ALLOWS ACCESS

To use the content-id api, the user must be a Youtube Content Partner.

If you are a Youtube Content Partner you can access the
Content ID Online Interface [here](https://www.youtube.com/content_id?o=U).

You will also see this api listed in google's API Service Manager
[here](https://console.developers.google.com/apis/dashboard)

## AUTHENTIFICATION - CONTENT-ID USES OAuth 2.0

Google uses many ways to authenticate, such as `API Key`,  `Service Account` or
`OAuth 2.0 Client ID`.

YouTube's Content ID uses `OAuth 2.0 Client ID` for authentication.

Refer to [OAuth 2.0 Authorization Cheat Sheet](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/authorization/OAuth-2.0-authorization-cheat-sheet)
for how to authenticate.
