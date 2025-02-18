**의존 관계 역전 원칙 (DIP, Dependency Inversion Principle)**

## 정의
의존 관계 역전 원칙(DIP)은 전통적인 상위 계층이 하위 계층에 직접 의존하는 구조를 반전(역전)시켜, 상위 계층이 하위 계층의 **구현으로부터 독립적**이도록 하는 원칙입니다.

## 원칙
1. **상위 모듈은 하위 모듈에 의존해서는 안 된다.** → 둘 다 추상 모듈(인터페이스 또는 추상 클래스)에 의존해야 한다.
2. **추상 모듈은 구체화된 모듈에 의존해서는 안 된다.** → 오히려, 구체화된 모듈이 추상 모듈에 의존해야 한다.

## DIP 위반 예제 (Go)

```go
type EmailService struct {}

func (e EmailService) SendEmail(to string, message string) {
    fmt.Println("Sending email to", to, ":", message)
}

type Notification struct {
    emailService EmailService
}

func (n Notification) Notify(to string, message string) {
    n.emailService.SendEmail(to, message)
}
```

### 문제점
- `Notification`은 `EmailService`의 **구체적인 구현에 직접 의존**하고 있습니다.
- 새로운 메시지 전송 방식(SMS, Push 알림 등)을 추가하려면 `Notification`을 수정해야 합니다.
- DIP를 위반하므로 확장성과 유지보수성이 낮습니다.

## DIP 준수 예제 (Go)

```go
type MessageSender interface {
    Send(to string, message string)
}

type EmailService struct {}

func (e EmailService) Send(to string, message string) {
    fmt.Println("Sending email to", to, ":", message)
}

type Notification struct {
    sender MessageSender
}

func (n Notification) Notify(to string, message string) {
    n.sender.Send(to, message)
}
```

### 개선점
- `MessageSender` 인터페이스를 도입하여 `Notification`이 **구체적인 구현이 아닌 인터페이스에 의존**하도록 변경했습니다.
- 이제 `EmailService` 외에도 **SMSService, PushNotificationService 등 다양한 구현을 추가할 수 있습니다.**
- `Notification`의 코드 변경 없이 새로운 전송 방법을 쉽게 추가할 수 있어 **OCP(개방-폐쇄 원칙)도 준수**합니다.

## 결론
DIP를 준수하면 **유연하고 확장 가능한 코드**를 작성할 수 있습니다. 특히, 의존성을 인터페이스로 추상화함으로써 변경에 강한 아키텍처를 만들 수 있습니다. 이를 통해 유지보수성과 재사용성이 향상됩니다.

