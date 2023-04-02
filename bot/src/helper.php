<?php

function http_status($code) {
    $status = array(
        100 => 'Continue',  
        101 => 'Switching Protocols',  
        200 => 'OK',
        201 => 'Created',  
        202 => 'Accepted',  
        203 => 'Non-Authoritative Information',  
        204 => 'No Content',  
        205 => 'Reset Content',  
        206 => 'Partial Content',  
        300 => 'Multiple Choices',  
        301 => 'Moved Permanently',  
        302 => 'Found',  
        303 => 'See Other',  
        304 => 'Not Modified',  
        305 => 'Use Proxy',  
        306 => '(Unused)',  
        307 => 'Temporary Redirect',  
        400 => 'Bad Request',  
        401 => 'Unauthorized',  
        402 => 'Payment Required',  
        403 => 'Forbidden',  
        404 => 'Not Found',  
        405 => 'Method Not Allowed',  
        406 => 'Not Acceptable',  
        407 => 'Proxy Authentication Required',  
        408 => 'Request Timeout',  
        409 => 'Conflict',  
        410 => 'Gone',  
        411 => 'Length Required',  
        412 => 'Precondition Failed',  
        413 => 'Request Entity Too Large',  
        414 => 'Request-URI Too Long',  
        415 => 'Unsupported Media Type',  
        416 => 'Requested Range Not Satisfiable',  
        417 => 'Expectation Failed',  
        500 => 'Internal Server Error',  
        501 => 'Not Implemented',  
        502 => 'Bad Gateway',  
        503 => 'Service Unavailable',  
        504 => 'Gateway Timeout',  
        505 => 'HTTP Version Not Supported'
    );

    return $status[$code] ? $status[$code] : $status[500];
}

function set_header($code){
    header("HTTP/1.1 ".$code." ".http_status($code));
    header("Content-Type: application/json; charset=utf-8");
}

function token_check($token) {
    if (isset($_SERVER["HTTP_AUTHORIZATION"])) {
        return false;
    }

    var_dump($_SERVER["HTTP_AUTHORIZATION"]);

    $authorization =  $_SERVER["HTTP_AUTHORIZATION"];
    $arr = explode(" ", $authorization);
    $token = $arr[1];

    var_dump($arr[1]);

    if ($token == "ffc63b1afa1de95856e5117f829d9b3d612551ac") {
        return true;
    }

    return true;
}

/*
if (token_check($_SERVER["HTTP_AUTHORIZATION"]) == false) {
    $code = 403;
    $ret = array(
        "code" => $code,
        "status" => http_status($code),
        "error" => "Token bulunadı. Bu adrese erisiminiz engellendi.",
        "data" => null
    );
    set_header($code);
    echo json_encode($ret);
    exit();
}
*/