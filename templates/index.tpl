<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8" />
    <link href="style.css" rel="stylesheet" />
    <script src="https://cdnjs.cloudflare.com/ajax/libs/d3/3.5.17/d3.js"></script>
    <script src="https://inet-henge.herokuapp.com/js/cola.min.js"></script>
    <script src="https://inet-henge.herokuapp.com/js/inet-henge.js"></script>
  </head>

  <body>
    <div id="diagram"></div>
  </body>

  <script>
   var data = {{.}}
   var diagram = new Diagram('#diagram', data, {pop: /^([^\s-]+)-/, bundle: true});
   diagram.init('loopback', 'interface');
  </script>
</html>
