const express = require('express');
const fs = require('fs');
const csv = require('csv-parser');
const app = express();
const PORT = 4000;

app.get('/test-bigdata', (req, res) => {
    const start = Date.now();
    const cpuStart = process.cpuUsage();
    let totalRows = 0;
    let totalAge = 0;
    let totalSalary = 0;

    fs.createReadStream('data/dataset.csv')
        .pipe(csv())
        .on('data', (row) => {
            totalRows++;
            totalAge += parseInt(row.age);
            totalSalary += parseInt(row.salary);
        })
        .on('end', () => {
            const end = Date.now();
            const memoryUsed = process.memoryUsage().rss / (1024 * 1024);
            const cpuUsage = process.cpuUsage(cpuStart);
            const cpuPercent = ((cpuUsage.user + cpuUsage.system) / 1000) / (end - start) * 100;

            res.json({
                total_rows: totalRows,
                avg_age: totalAge / totalRows,
                total_salary: totalSalary,
                execution_time_ms: end - start,
                memory_used_mb: parseFloat(memoryUsed.toFixed(2)),
                cpu_percent: parseFloat(cpuPercent.toFixed(2))
            });
        });
});

app.listen(PORT, () => {
    console.log(`Node.js service running on port ${PORT}`);
});
