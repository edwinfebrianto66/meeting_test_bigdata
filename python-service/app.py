from flask import Flask, jsonify
import pandas as pd
import time
import os
import psutil

app = Flask(__name__)

@app.route('/test-bigdata', methods=['GET'])
def test_bigdata():
    process = psutil.Process(os.getpid())
    cpu_start = process.cpu_times()
    start_time = time.time()

    df = pd.read_csv('data/dataset.csv')
    total_rows = len(df)
    avg_age = df['age'].mean()
    total_salary = df['salary'].sum()

    end_time = time.time()
    cpu_end = process.cpu_times()
    execution_time_ms = int((end_time - start_time) * 1000)
    memory_used_mb = round(process.memory_info().rss / (1024 * 1024), 2)
    cpu_percent = process.cpu_percent(interval=0.1)

    return jsonify({
        'total_rows': total_rows,
        'avg_age': round(avg_age, 2),
        'total_salary': int(total_salary),
        'execution_time_ms': execution_time_ms,
        'memory_used_mb': memory_used_mb,
        'cpu_percent': round(cpu_percent, 2)
    })

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=4002)
