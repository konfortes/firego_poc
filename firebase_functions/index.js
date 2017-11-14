// The Cloud Functions for Firebase SDK to create Cloud Functions and setup triggers.
const functions = require('firebase-functions');

// The Firebase Admin SDK to access the Firebase Realtime Database.
const admin = require('firebase-admin');
admin.initializeApp(functions.config().firebase);

exports.createSupplier = functions.firestore
  .document('suppliers/{ID}')
  .onCreate(event => {
    // Then return a promise of a set operation to update the count
    console.log(event);
    console.log(event.auth);
    console.log('supplier_id: ' + event.params.ID);
    return event.data.ref.set({
      created_by: 'RONEN'
    }, {merge: true});
  });

  exports.updateSupplier = functions.firestore
    .document('suppliers/{ID}')
    .onUpdate(event => {
      const data = event.data.data();
      const previousData = event.data.previous.data();

    // We'll only update if the name has changed.
    // This is crucial to prevent infinite loops.
    if (data.name == previousData.name) return;

      console.log(event)
      return event.data.ref.set({
        updated_by: 'RONEN'
      }, {merge: true});
    });
