# ProjectGoLive (Inventory Tracker)

## Capstone project of GoSchool

Picture this:
A friend gives you a box of muffin, you chuck it in the fridge and accidentally forgets about it as you are busy.
A week later and suddenly you have a midnight craving, only to reach into your fridge and find a box of expired muffins.
Either you risk having stomachache or it's another box of muffin thrown and wasted in the bin.

## Project Brief
- Inventory tracker that actively tracks expiry dates & usage.
- Sends notifications to the user as per user's settings (eg. near expiry).
- Intended mostly towards perishables goods like food, but also tracks other household items like tools left in storage.
- Main purpose of the app is to reduce wastage through active monitoring of the inventory.

## Requirements
- Go (1.15 & above)

## How to run
- Create a PGL_db schema in your MYSQL database, Gorm package will auto port the tables over
- Populate your env file with the required variables
- Run the precompiled .exe

## Features
- Tracks expiry dates (& how long till it's expired)
- Sends notification to the user by Twilio
- Simple UI

## CreatedBy
- [monkeygoessnap](https://github.com/monkeygoessnap)