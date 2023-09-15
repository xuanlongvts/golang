<?php
// $result = parseTags("a #bcd #lsdjf ddd");
$result = parseTags(" aa #sdf #dfd https://php-play.dev/?c=DwfgDgFmBQAkBOBTAzgVw#DYBcAEBebYAhvMo#gCqEDm");

echo "================ <br />";
print json_encode($result);

function parseTags($string = ''){
    $regexLink = "/(https?:\/\/(?:www\.|(?!www))[a-zA-Z0-9][a-zA-Z0-9-]+[a-zA-Z0-9]\.[^\s]*[-A-Z0-9+&@#\/%=~_|]{2,}|www\.[a-zA-Z0-9][a-zA-Z0-9-]+[a-zA-Z0-9]\.[^\s]*[-A-Z0-9+&@#\/%=~_|]{2,}|https?:\/\/(?:www\.|(?!www))[a-zA-Z0-9]+\.[^\s]*[-A-Z0-9+&@#\/%=~_|]{2,}|www\.[a-zA-Z0-9]+\.[^\s]*[-A-Z0-9+&@#\/%=~_|]{2,})/im";

    $matcheLinks = [];
    $matchCountLink = preg_match_all($regexLink, $string, $matcheLinks);
    $arrLinks = [];
    $i = 0;
    while ($i < $matchCountLink) {
        $arrLinks[] = $matcheLinks[1][$i];
        $i++;
    }

    if (count($arrLinks) > 0) {
        echo "1. <br /> ";
        print_r($arrLinks);
    }

    // parse tags
    $regex = "/(#[A-Za-z0-9_\-\]\[]+)/im";
    $matches = [];
    $matchCount = preg_match_all($regex, $string, $matches);
    $returnData = [];
    $i = 0;

    if (count($matches) > 0) {
        echo "<br /> 2.----------------------- <br />";
        print_r($matches[0]);
        echo "<br /><br />";
    }

    while ($i < $matchCount) {
        // Check exists tags in link
        $checkExists = false;
        foreach ($arrLinks as $link) {
            if (strpos($link, $matches[1][$i]) !== false) {
                $checkExists = true;
            }
        }

        if (!$checkExists) {
            $returnData[] = [
                'name' => str_replace(['[', ']', ' ', '#'], '', $matches[1][$i])
            ];
        }

        $i++;
    }
    return $returnData;
}