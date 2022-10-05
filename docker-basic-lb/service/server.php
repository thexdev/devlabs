<?php

header('Content-Type: application/json');

$message = sprintf(
    'Hello from service %s',
    getenv('SERVICE_ID')
);

echo json_encode(compact('message'));
