# idy
rest service for two factor authentication

Rest service for Two factor based authentication. 
It provide :-

a) bar code generation - this allows you to display a qr code for user to scan using google authenticator.

  This will generate an secret key, which you can store in your data store.

b) validation - validates the key generated with user's secret. Typeically, key in something like 201011 and we will validate if 
this is legit. 

c) key storage. 

## Generate QR 
/generateQr

## Validation REST 
/validate?key=208088



