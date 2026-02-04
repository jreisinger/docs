Amazon EventBridge (AI generated notes, 4.2.2026)

1. Co to je
- ako CloudWatch events ale rozsirena funkcionalita
- event bus na smerovanie eventov
- event je JSON
2. Preco existuje
- Chcel si reagovať na zmenu stavu EC2? Musel si pollovat API
- Chcel si spustiť Lambda každú hodinu? Musel si riešiť cron externe
- Chcel si prepojiť služby? Musel si písať glue code
3. Ako to funguje
```
┌─────────────────┐     ┌─────────────────┐     ┌─────────────────┐
│                 │event│                 │event│                 │  
│  EVENT SOURCE   │────▶│   EVENT BUS     │────▶│    TARGET       │
│                 │     │                 │     │                 │
│ • AWS Services  │     │  ┌───────────┐  │     │ • Lambda        │
│ • Custom Apps   │     │  │   RULES   │  │     │ • SNS/SQS       │
│ • Scheduled     │     │  │ (filter)  │  │     │ • Step Functions│
│                 │     │  └───────────┘  │     │ • EC2/ECS       │
└─────────────────┘     └─────────────────┘     └─────────────────┘
```

Priklad: Scheduled Lambda (cron trigger)

Lambda funkcia spustená CloudWatch Events pravidlom každých 5 minút:

```go
package main

import (
      "context"
      "fmt"
      "time"

      "github.com/aws/aws-lambda-go/events"
      "github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, event events.CloudWatchEvent) error {
      fmt.Printf("Event source: %s\n", event.Source)
      fmt.Printf("Detail type: %s\n", event.DetailType)
      fmt.Printf("Time: %s\n", event.Time.Format(time.RFC3339))
      fmt.Printf("Detail: %s\n", string(event.Detail))

      // tvoja logika tu — cleanup, report, sync...
      return nil
}

func main() {
      lambda.Start(handler)
}
```

Terraform pravidlo, ktoré to spúšťa:

```tf
resource "aws_cloudwatch_event_rule" "every_5_min" {
  name                = "every-5-minutes"
  schedule_expression = "rate(5 minutes)"
}

resource "aws_cloudwatch_event_target" "lambda" {
  rule = aws_cloudwatch_event_rule.every_5_min.name
  arn  = aws_lambda_function.my_func.arn
}
```
