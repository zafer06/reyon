<?php
namespace Solokod\Bot;

use DOMDocument;
use DateInterval;
use DateTime;

class Crawler {

    public function __construct() {
        // "bot constructor...";
    }

    public function run_bot() {
        //$url = "http://localhost:8080/reyon-data.html";
        
        $dom = new DOMDocument('1.0');
        @$dom->loadHTMLFile("reyon-data.html");
        $dom->preserveWhiteSpace = false;
        
        $form = $dom->getElementById("inlinemodform");
        $threads = $form->getElementsByTagName("ol");

        $job_list = array();
    
        foreach ($threads as $li) {
            $title = $li->getElementsByTagName("a")->item(1)->nodeValue;
            $link = $li->getElementsByTagName("a")->item(1)->getAttribute("href");
            $user = $li->getElementsByTagName("a")->item(2)->nodeValue;
            $date = $this->format_date($li->getElementsByTagName("div")->item(6)->nodeValue);
            $reply_user = $li->getElementsByTagName("div")->item(7)->nodeValue;
            
            //var_dump($date);
    
            $job = array(
                "title" => $title,
                "link"  => $link,
                "user"  => $user,
                "date"  => $date,
                "reply" => $reply_user
            );
    
            array_push($job_list, $job);
        }
    
        return $job_list;
    }
    
    private function format_date($date) {
        if ($date == "") return $date;
    
        $date_arr = explode(" ", $date);
        $d = $date_arr[0];
        $t = $date_arr[1];
    
        $today = new DateTime();
    
        if ($d === "Bugün") {
            return $today->format('Y-m-d') ." ". $t;
    
        } else if ($d == "Dün") {
            $interval = new DateInterval('P1D');
            $today->sub($interval); 
            return $today->format('Y-m-d') ." ". $t;
    
        } else {
            $arr = explode("-", $d);
            return $arr[2] ."-". $arr[1] ."-". $arr[0] ." ". $t;
        }
    
        return $date;
    }
}