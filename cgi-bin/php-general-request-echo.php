<html>
<?php

header("Cache-Control: no-cache");
header("Content-type: text/html");

?>

<head><title>General Request Echo</title></head>

<body>

<h1 align=center>General Request Echo</h1>
<hr/>

<table>
<?php
  echo "<tr><td>Protocol:</td><td>". getenv("SERVER_PROTOCOL"). "</td></tr>"; 
  echo "<tr><td>Method:</td><td>". getenv("REQUEST_METHOD"). "</td></tr>"; 
  echo "<tr><td>Query String:</td><td>". getenv("QUERY_STRING"). "</td></tr>"; 
  echo "<tr><td>Message Body:</td><td>". file_get_contents("php://input"). "</td></tr>"; 
?>
</table>

</body>
</html>