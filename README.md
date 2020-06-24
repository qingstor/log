# log

A log designed for critical mission.

## Goals

- Safe
  - Type Safe: No reflection
  - Runtime Safe: No panic, no error check
- Performance
  - Low Latency: Log entry with 10 fields in 1ms
  - Low Footprint: No extra objects allocated

## Design

### Where is `Fatal` and `Panic`?