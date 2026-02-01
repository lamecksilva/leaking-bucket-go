Leaking Bucket Rate Limiter (Go)

This repository contains a simple and practical implementation of the Leaking Bucket rate limiting algorithm written in Go.
The goal is to demonstrate how the algorithm works, how it behaves under concurrent load, and how it can be tested using a stress script.

What is the Leaking Bucket algorithm?

Leaking Bucket is a rate limiting algorithm that processes requests at a constant and predictable rate, regardless of how bursty incoming traffic is.

You can think of it as a bucket with a small hole at the bottom. Incoming requests fill the bucket, while requests are processed by “leaking” out at a fixed speed. If requests arrive faster than they can be processed and the bucket becomes full, new requests are rejected.

This approach is useful for smoothing traffic spikes, protecting services from overload, and enforcing stable throughput.

How this implementation works

This project implements the Leaking Bucket algorithm using Go concurrency primitives.

Requests are added to a bucket with a fixed capacity. A worker processes requests at a constant interval, simulating the leak rate. When the bucket is full, incoming requests are denied immediately, which makes rate limiting behavior easy to observe during stress tests.

The implementation is intentionally simple and educational, focusing on clarity rather than production-level optimizations.

Requirements

Go 1.20 or newer

curl (for stress testing)

Running the project

Clone the repository and navigate into it:

git clone https://github.com/lamecksilva/leaking-bucket-go.git
cd leaking-bucket-go


Run the application:

go run main.go


By default, the server will start locally and expose an endpoint protected by the Leaking Bucket rate limiter.

Testing with the stress script

The repository includes a stress script designed to send multiple concurrent requests using curl, making it easy to observe rate limiting in action.

Make the script executable:

chmod +x stress.sh


Run the stress test:

./stress.sh


The script sends multiple requests at the same time. As the request rate exceeds the leak rate, you should start seeing rejected requests (for example, HTTP 429 responses), confirming that the rate limiter is working as expected.

Why Leaking Bucket?

Unlike algorithms that allow bursts (such as Token Bucket), Leaking Bucket enforces a strict, steady output rate. This makes it ideal for systems that require predictable processing and controlled resource usage, even when traffic arrives in sudden spikes.

Notes

This project is meant for learning and experimentation. It provides a clear view of how Leaking Bucket behaves under load and how concurrency in Go can be used to implement rate limiting mechanisms.
