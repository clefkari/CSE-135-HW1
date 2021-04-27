<?php

header("Cache-Control: no-cache");
header("Content-type: text/html");

?>

<head><title>Environment Variables</title></head>

<body>

<h1 align=center>Environment Variables</h1>
<hr/>

<ul>
<?php

foreach (getenv() as $key => $val){
  echo "<li>{$key} : {$val}</li>";
}

foreach ($_SERVER as $key => $val){
  echo "<li>{$key} : {$val}</li>";
}

?>
</ul>

</body>
</html>