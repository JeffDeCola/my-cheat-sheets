# SINGLE-VENDOR SINGLE-STORE CHEAT SHEET

_A Single-Vendor (Single-Store) E-Commerce Platform.
More simply, **an online store**._

Other E-Commerce platform configurations,

* **MULTI-VENDOR**
  * [multi-vendor marketplace](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/e-commerce/multi-vendor/multi-vendor-marketplace-cheat-sheet)
  * [multi-vendor multi-store](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/e-commerce/multi-vendor/multi-vendor-multi-store-cheat-sheet)
* **SINGLE-VENDOR**
  * [single-vendor single-store](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/e-commerce/single-vendor/single-vendor-single-store-cheat-sheet)
    **(You are Here)**

Table of Contents,

* [OVERVIEW](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/e-commerce/single-vendor/single-vendor-single-store-cheat-sheet#overview)
* [USERS (ADMIN, VENDORS & CUSTOMERS)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/e-commerce/single-vendor/single-vendor-single-store-cheat-sheet#users-admin-vendors--customers)
* [STEPS TO CREATE AN ONLINE STORE USING WORDPRESS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/e-commerce/single-vendor/single-vendor-single-store-cheat-sheet#steps-to-create-an-online-store-using-wordpress)
  * [SOFTWARE STACK](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/e-commerce/single-vendor/single-vendor-single-store-cheat-sheet#software-stack)
  * [INSTALATION STEPS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/e-commerce/single-vendor/single-vendor-single-store-cheat-sheet#instalation-steps)
  * [CONFIGURATION STEPS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/e-commerce/single-vendor/single-vendor-single-store-cheat-sheet#configuration-steps)

Documentation and reference,

* [See my offsite demo](https://single-vendor-single-store.jeffdecola.com)
* [WordPress](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-architectures/software-as-a-service/wordpress-cheat-sheet)
  content management system

[GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/)

## OVERVIEW

The following illustration is a high-level view of an
Single-Vendor (Single-Store) E-Commerce platform (SaaS)
on a private server using,

* WordPress (Content Management System)
* WooCommerce PlugIn (E-Commerce Platform)

Or more simply, an **online store** hosted on your private server.

![IMAGE - single-vendor-single-store.jpg - IMAGE](../../../../docs/pics/single-vendor-single-store.jpg)

## USERS (ADMIN, VENDORS & CUSTOMERS)

There are 3 main users of an E-Commerce site that have
the following capabilities, 

* **ADMIN**
  * Manages Stores
  * Manages Vendors and Customers
  * Integrate CRM (user.com)
  * Payment systems (Stripe/Paypal)
  * Shipping systems (USPS/FedEx)
* **VENDORS**
  * Manage E-Commerce Website
  * Editing inventory and Pricing
  * CMS Page Editor
  * Payments
  * Managing Orders
  * Sales Analytics
* **CUSTOMERS**
  * Browsing Products/Categories
  * Live Search
  * Coupons/Discount Codes
  * SEO Optimised
  * Purchasing
  * Shipping/Tracking

## STEPS TO CREATE AN ONLINE STORE USING WORDPRESS

These are the high-level steps I did to create my
[online store demo](https://single-vendor-single-store.jeffdecola.com).

### SOFTWARE STACK

* WordPress
  * WooCommerce Plugin
    * WooCommerce Shiping & Tax Plugin
    * Creative Mail by Constant Contact Plugin
    * Facebook for WooCommerce Plugin
    * Google Ads & Marketing by Kliken Plugin
    * Mailchimp for WooCommerce Plugin
    * WC Vendors Marketplace Plugin
  * Loginizer Plugin

### INSTALLATION STEPS

* GET WORDPRES SETUP ON YOUR SERVER
  * Get a WebServer like Bluehost
  * Buy a domain name
  * Setup SSL (i.e. https)
  * Install WordPress
* INSTALL WOOCOMMERCE PLUGIN
  * Install the WooCommerce Plugin
  * The setup wizard will help you with basic plugin setup
  * Chose the Storefront theme
  * This will add four new Pages: **Cart, Checkout, My account and Shop**
  * This will add 2 new Users: **Shop manager & Customer**
* INSTALL OTHER PLUGINS
  * Install the Loginizer Plugin
  * Install Jetpack by wordpress.com

### USERS

What you should have,

* Vanilla Users
  * Subscriber
  * Contributor
  * Author
  * Editor
  * Administrator
* WooCommerce Users
  * Shop manager
  * Customer

### PAGES

What you should have,

* Vanilla Pages
  * Privacy Policy (Privacy Policy Page)
  * Sample Page
* WooCommerce Pages
  * Cart (Cart Page)
  * Checkout (Checkout Page)
  * My account (My Accounts Page)
  * Shop (Shop Page)

## BASIC WOOCOMMERCE CONFIGURATION

* Go throught the steps to get the store up means having a product.
  * Add a product
  * Setup payments
  * Setup tax
  * Setup shipping

### PRETTY IT UP

How to make it look better using the WooCommerce Storefront Theme,

* Add a logo
* 