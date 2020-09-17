---
author: Sacha Richters
date: 11-Sep-2020
---

# Renaming files application

## Purpose

This application is used for customers of MasterControl to receive their files from EFP with the naming convention [infocard_id][revision][vault].[existing extension]

The application would receive an input from the user in the form of the keys.txt file. This file contains the information from the database about the infocardkey, infocard_id, revision and vault.

## Language choice

Go languages can be used for many different applications, just like Python. One of these usage is scripting; automation.
This is what we need to make the application.

Laurens has previously made a scripting application with Go language for a customer to receive specific files from the EFP. I have done some research into Go because I wanted to know why he used Go for that tool instead of Python (we build our gxp cloud application with Python, so he knew that language, so there had to be a reason why he chose another language).
I found that Python is very slow when executing, and has typing issues. Go language has solved both these issues, so it is faster (which is good as we would need to change a lot of files) and there would be less issues with types.
I would also need just as much time to learn scripting with Python as I would need to learn scripting with Go.

This is why I choose Go for this application.
