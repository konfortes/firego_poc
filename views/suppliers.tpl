<!DOCTYPE html>
<html>

<head>
  <meta charset="utf-8">
  <title></title>
  <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css" integrity="sha384-BVYiiSIFeK1dGmJRAkycuHAHRg32OmUcww7on3RYdg4Va+PmSTsz/K68vbdEjh4u" crossorigin="anonymous">
</head>

<body>
  <div id="app">
    <actions-table :actions="actions"></actions-table>
  </div>
</body>

<script type="text/javascript" src="https://unpkg.com/vue"></script>
<!-- <script type="text/javascript" src="https://unpkg.com/axios/dist/axios.min.js"></script> -->
<script src="https://www.gstatic.com/firebasejs/4.5.1/firebase.js"></script>
<script src="https://www.gstatic.com/firebasejs/4.5.1/firebase-firestore.js"></script>


<script>
  var config = {
    apiKey: "AIzaSyBAhUhctGLvzyidJbCFkwWYqugDZAdpISQ",
    authDomain: "actions-f7aae.firebaseapp.com",
    databaseURL: "https://actions-f7aae.firebaseio.com",
    projectId: "actions-f7aae",
    storageBucket: "actions-f7aae.appspot.com",
    messagingSenderId: "506602779408"
  };
  firebase.initializeApp(config);

  var db = firebase.firestore();

  db.collection("suppliers").add({
    supplier_id: 666,
    actions: [{
      type: 'js'
    }]
  })
</script>

</html>
