<?php

header("Cache-Control: no-cache");
header("Content-type: text/html");

?>

<head><title>GET query string</title></head>

<body>

<h1 align=center>GET query string</h1>
<hr/>

<p>Query String: 
<?php
$query = getenv('QUERY_STRING');
echo $query;
?>
</p>

<table> Formatted Query String:
  <?php
    $strings = explode("&",$query);
    foreach ($strings as $string){
      echo "<tr><td>". str_replace("="," : ",$string). "</tr></td>";
    }
  ?>
</table>

</body>
</html>