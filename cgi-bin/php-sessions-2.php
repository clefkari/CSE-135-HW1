<html>
<?php

header("Cache-Control: no-cache");
header("Content-type: text/html");

?>

<body>

    <head><title>PHP Sessions</title></head>
    <h1>PHP Sessions Page 2</h1>
    <table>
        <?php
        if (getenv('HTTP_COOKIE') != NULL && strcmp(getenv('HTTP_COOKIE'), "destroyed")){
            echo "<tr><td>Cookie:</td><td>". getenv('HTTP_COOKIE'). "</td></tr>\n";
        }else{
            echo "<tr><td>Cookie:</td><td>None</td></tr>";
        }
        ?>
    </table>

    <br/>
    <a href="/cgi-bin/php-sessions-1.php">Session 1 Page</a>
    <br/>
    <a href="/php-cgiform.html">CGI Page</a>
    <br/><br/>

    <form action="/cgi-bin/php-destroy-session.php" method="get">
    <button type="submit">Destroy Session</button>
    
</body>
</html>