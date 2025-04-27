FROM python:3.10-slim

WORKDIR /app

# Tambahkan ini: install gcc, libpq-dev, dan upgrade pip
RUN apt-get update && apt-get install -y gcc build-essential \
    && pip install --upgrade pip setuptools wheel

COPY . .

RUN pip install --no-cache-dir -r requirements.txt

CMD ["python", "app.py"]
