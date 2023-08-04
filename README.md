# Meta-conversion-api
# Golang API for Sending Purchase Events to Facebook Pixel

## Introduction

This Golang API webserver is designed to send Purchase Events to a Facebook Pixel. The API accepts Purchase Events in the "mable" format, processes the data, and sends it to the Facebook Pixel for event tracking.

## Prerequisites

Before running the API, make sure you have the following installed:

- Go programming language (version 1.15 or later)

## Installation

1. Clone this repository to your local machine.
2. Navigate to the project directory.

## Usage

1. Replace the dummy data in the provided payload with your actual Facebook Pixel ID, access key, and test event code.

2. Start the Golang API webserver by running the following command in the terminal: go run "file path"

3. The API will be running at http://localhost:8080.

4. To send a Purchase Event to Facebook Pixel, make a POST request to the API endpoint '/post/facebook' with the payload containing the necessary data in the "mable" format.

5. The API will validate the payload, perform normalization and hashing as per Facebook API requirements, and send the event data to the Facebook Pixel.

6. The API will respond with a JSON object indicating the status of the event submission.
