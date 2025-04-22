<?php
$start = microtime(true);

$file = fopen("data/dataset.csv", "r");
if ($file === false) {
    http_response_code(500);
    echo json_encode(["error" => "Cannot open file"]);
    exit;
}

$header = fgetcsv($file);
$totalRows = 0;
$totalAge = 0;
$totalSalary = 0;

while (($row = fgetcsv($file)) !== false) {
    $totalRows++;
    $totalAge += (int)$row[2];
    $totalSalary += (int)$row[3];
}
fclose($file);

$avgAge = $totalAge / $totalRows;
$executionTime = round((microtime(true) - $start) * 1000, 2);
$memoryUsed = round(memory_get_usage(true) / 1024 / 1024, 2);

header("Content-Type: application/json");
echo json_encode([
    "total_rows" => $totalRows,
    "avg_age" => round($avgAge, 2),
    "total_salary" => $totalSalary,
    "execution_time_ms" => $executionTime,
    "memory_used_mb" => $memoryUsed
]);
?>
