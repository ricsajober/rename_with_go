---
author: Sacha Richters
date: 11-Sep-2020
---

# Renaming files application

## Purpose

This application is used for customers of MasterControl to receive their files from EFP with the naming convention [infocard_id][revision][vault].[existing extension]

The application would receive an input from the user in the form of the keys.txt file. This file contains the information from the database about the infocardkey, infocard_id, revision and vault.

## Language choice

This issue is to complicated to simply solve with a batch-file. Therefor we would need to create an application.  
The languages we can use for this application are: c++, java, Python and Go.

Applications written with c++ are usually slow to compile and the language is complex. Java is very dificult to set-up and use (and also not the easiest language). Pyhton is a lot easier in use, but does not have type safety, it is also relatively slow compared to c++ and java.

Go is supposed to solve the issues that the other languages have. Also there was an application made previously with the Go Language (by Laurens) which was used to extract certain files from the EFP based on user-input. For this application we are able to re-use some code from that application.

Therefor we would choose the Go language. Personally I have seen all four languages before, and worked with them. But I would still need to learn a bit to actually create this application
