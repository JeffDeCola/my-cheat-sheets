# SINGLE-VENDOR SINGLE-STORE CHEAT SHEET

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_A Single-Vendor (Single-Store) E-Commerce Platform.
More simply, **an online store**._

Other Configurations

* [multi-vendor marketplace](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-architectures/software-as-a-service/content-management/multi-vendor-marketplace-cheat-sheet)
* [multi-vendor multi-store](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-architectures/software-as-a-service/content-management/multi-vendor-multi-store-cheat-sheet)
* [single-vendor single-store](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-architectures/software-as-a-service/content-management/single-vendor-single-store-cheat-sheet)
  **(You are Here)**

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-architectures/software-as-a-service/content-management/single-vendor-single-store-cheat-sheet#overview)
* [USERS (ADMIN, VENDORS & CUSTOMERS)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-architectures/software-as-a-service/content-management/single-vendor-single-store-cheat-sheet#users-admin-vendors--customers)
* [CREATE AN ONLINE STORE USING WORDPRESS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-architectures/software-as-a-service/content-management/single-vendor-single-store-cheat-sheet#create-an-online-store-using-wordpress)
  * [SOFTWARE STACK](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-architectures/software-as-a-service/content-management/single-vendor-single-store-cheat-sheet#software-stack)
  * [INSTALL & BASIC CONFIGURATION](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-architectures/software-as-a-service/content-management/single-vendor-single-store-cheat-sheet#install--basic-configuration)
  * [USERS & PAGES](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-architectures/software-as-a-service/content-management/single-vendor-single-store-cheat-sheet#users--pages)
  * [BASIC WOOCOMMERCE CONFIGURATION](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-architectures/software-as-a-service/content-management/single-vendor-single-store-cheat-sheet#basic-woocommerce-configuration)
  * [PRETTY IT UP](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-architectures/software-as-a-service/content-management/single-vendor-single-store-cheat-sheet#pretty-it-up)

Documentation and Reference

* my offsite
  [demo of an online store](https://single-vendor-single-store.jeffdecola.com)
* [wordPress](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-architectures/software-as-a-service/content-management/wordpress-cheat-sheet)
  content management system
* [WooCommerce](https://woocommerce.com/)
  e-commerce plugin

## OVERVIEW

The following illustration is a high-level view of an
Single-Vendor (Single-Store) E-Commerce platform (SaaS)
on a private server using,

* WordPress (Content Management System)
* WooCommerce PlugIn (E-Commerce Platform)

Or more simply, this illustration shows an **online store**
hosted on your private server,

![IMAGE - single-vendor-single-store.jpg - IMAGE](../../../../../docs/pics/software/service-architectures/single-vendor-single-store.svg)

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
  * SEO Optimized
  * Purchasing
  * Shipping/Tracking

## CREATE AN ONLINE STORE USING WORDPRESS

These are the high-level steps I did to create my
[online store demo](https://single-vendor-single-store.jeffdecola.com).

### SOFTWARE STACK

* WORDPRESS
  * WooCommerce Plugin
    * WooCommerce Shipping & Tax Plugin
    * Creative Mail by Constant Contact Plugin
    * Facebook for WooCommerce Plugin
    * Google Ads & Marketing by Kliken Plugin
    * Mailchimp for WooCommerce Plugin
    * WC Vendors Marketplace Plugin
  * Loginizer Plugin

### INSTALL & BASIC CONFIGURATION

* GET WORDPRESS SETUP ON YOUR SERVER
  * Get a WebServer like Bluehost
  * Buy a domain name
  * Setup SSL (i.e. https)
  * Install WordPress
* INSTALL WOOCOMMERCE PLUGIN
  * Install the WooCommerce Plugin
  * The setup wizard will help you with basic plugin setup
  * Chose the Storefront theme
  * This will add four new Pages: **Cart, Checkout, My account & Shop**
  * This will add 2 new Users: **Shop manager & Customer**
* INSTALL OTHER PLUGINS
  * Install the Loginizer Plugin
  * Install Jetpack by wordpress.com

### USERS & PAGES

What you should have for users,

* WORDPRESS USERS
  * Subscriber
  * Contributor
  * Author
  * Editor
  * Administrator
* WOOCOMMERCE USERS
  * Shop manager
  * Customer

What you should have for pages,

* WORDPRESS PAGES
  * Privacy Policy (Privacy Policy Page)
  * Sample Page
* WOOCOMMERCE PAGES
  * Cart (Cart Page)
  * Checkout (Checkout Page)
  * My account (My Accounts Page)
  * Shop (Shop Page)

### BASIC WOOCOMMERCE CONFIGURATION

* Go through the steps to get the store up means having a product.
  * Add a product
  * Setup payments
  * Setup tax
  * Setup shipping

### PRETTY IT UP

How to make it look better using the WooCommerce Storefront Theme,

* Add a logo
