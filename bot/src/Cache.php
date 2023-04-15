<?php
namespace Solokod\Bot;

use DateTime;

const CACHE_FILE = "cache.json";

class Cache {

    public function __construct() {
        // Cache object init
    }

    public function cache_data() {
     if (file_exists(CACHE_FILE)) {
            //$file_creation_date = filectime(CACHE_FILE);
            $hour_diff = cache_time(CACHE_FILE);

            bot_log("Cache dosya tarihi -> " . date('Y-m-d H:i:s', filemtime(CACHE_FILE)));
            bot_log("Cache dakika farki -> " . $hour_diff);

            if ($hour_diff > 59) {
                $this->remove_cache();
                $this->add_cache();
            }
        } else {
            $this->add_cache();
        }

        return $this->get_cache();
    }

    private function add_cache() {
        $bot = new Crawler();
        $data = $bot->run_bot();

        $json = json_encode($data);
        return file_put_contents(CACHE_FILE, $json);
    }

    private function get_cache() {
        $data = file_get_contents(CACHE_FILE);
	    return json_decode($data);
    }

    private function remove_cache(): bool {
        return unlink(CACHE_FILE);
    }
}