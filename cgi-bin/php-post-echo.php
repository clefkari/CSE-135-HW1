<html>
<?php

header("Cache-Control: no-cache");
header("Content-type: text/html");

?>

<head><title>POST Message Body</title></head>

<body>

<h1 align=center>POST Message Body</h1>
<hr/>

<p>Message Body: 
<?php
  $_POST = file_get_contents("php://input");
  echo $_POST;
?>
</p>

</body>
</html>