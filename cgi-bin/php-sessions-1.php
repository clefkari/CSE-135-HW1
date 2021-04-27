<html>
<?php

header("Cache-Control: no-cache");
header("Content-type: text/html");

$username = file_get_contents("php://input");
$name     = "";

if ($username[0] == 'u'){
    $name = substr($username, 9);
}

if (strlen($name) > 0){
    header("Set-Cookie: {$name}");
}

?>

<body>

    <head><title>PHP Sessions</title></head>
    <h1>PHP Sessions Page 1</h1>
    <table>
        <?php
        if(strlen($name) > 0){
            echo "<tr><td>Cookie:</td><td>". $name. "</td></tr>";
        }else if (getenv('HTTP_COOKIE') != NULL && strcmp(getenv('HTTP_COOKIE'), "destroyed")){
            echo "<tr><td>Cookie:</td><td>". getenv('HTTP_COOKIE'). "</td></tr>\n";
        }else{
            echo "<tr><td>Cookie:</td><td>None</td></tr>";
        }
        ?>
    </table>

    <br/>
    <a href="/cgi-bin/php-sessions-2.php">Session 2 Page</a>
    <br/>
    <a href="/php-cgiform.html">CGI Page</a>
    <br/><br/>

    <form action="/cgi-bin/php-destroy-session.php" method="get">
    <button type="submit">Destroy Session</button>
    
</body>
</html>