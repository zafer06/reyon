<?php
namespace Solokod\Bot;

require "vendor/autoload.php";
require "src/helper.php";

use Solokod\Bot\Cache;

$route = null; 
if (isset($_GET["route"])) {
    $route = addslashes($_GET["route"]);
}

$code = 200;
$method = $_SERVER["REQUEST_METHOD"];
$ret = array();

//if ($method == "POST" && $route == "list")

if ($method == "POST") {
    try {
        $cache = new Cache();
        $data = $cache->cache_data();

        $ret = array(
            "code" => 200,
            "status" => http_status(200),
            "error" => "",
            "data" => $data
        );

    } catch(Exception $e) {
        $ret = array(
            "code" => 500,
            "status" => http_status(500),
            "error" => $e->getMessage(),
            "data" => null
        );
    }
    
} else {
    $ret = array(
        "code" => 405,
        "status" => http_status(405),
        "error" => "Method was not allowed",
        "data" => null
    );
}

set_header($code);
echo json_encode($ret);