<html>

<?php
header("Cache-Control: no-cache");
header("Content-type: text/html");
?>

<head><title>Hello CGI World</title></head>

<body>

<h1 align=center>Hello HTML World</h1>
<hr/>

<p>Hello, World, from William and Chris!</p>
<p><?php echo "This program was generated at: ". date('h:m:s Y-m-d'). " UTC" ?></p>
<p><?php echo "Your current IP address is: ". getenv('REMOTE_ADDR') ?></p>


</body>
</html>
