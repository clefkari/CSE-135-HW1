<html>
<?php

header("Cache-Control: no-cache");
header("Content-type: text/html");
header("Set-Cookie: destroyed");

?>

<body>

    <head><title>PHP Session Destroyed</title></head>
    <h1>PHP Session Destroyed</h1>

    <br/>
    <a href="/cgi-bin/php-sessions-1.php">Session 1 Page</a>
    <br/>
    <a href="/cgi-bin/php-sessions-2.php">Session 2 Page</a>
    <br/>
    <a href="/php-cgiform.html">CGI Page</a>
    <br/>
    
</body>
</html>