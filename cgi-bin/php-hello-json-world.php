<?php

header("Cache-Control: no-cache");
header("Content-type: application/json");

$obj->message   = "Hello, World, from William and Chris!";
$obj->date      = date('h:m:s Y-m-d'). " UTC";
$obj->currentIP = getenv('REMOTE_ADDR');

echo json_encode($obj);

?>