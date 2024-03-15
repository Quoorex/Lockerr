FROM python:3.13.0a4-slim

WORKDIR /lockerr/

# Install Poetry
RUN pip install --no-cache poetry && \
    poetry config virtualenvs.create false

COPY pyproject.toml poetry.lock ./

RUN poetry install --no-root

COPY . .

ENV BOT_TOKEN=""

CMD ["poetry", "run", "python", "-m", "lockerr"]