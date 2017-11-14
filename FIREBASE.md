npm install -g firebase-tools

firebase login
firebase init functions

// The Cloud Functions for Firebase SDK to create Cloud Functions and setup triggers.
const functions = require('firebase-functions');

// The Firebase Admin SDK to access the Firebase Realtime Database.
const admin = require('firebase-admin');
admin.initializeApp(functions.config().firebase);

// Listen for any change on document `marie` in collection `users`
exports.myFunctionName = functions.firestore
  .document('users/marie').onWrite((event) => {
    // ... Your code here
  });

firebase deploy --only functions
