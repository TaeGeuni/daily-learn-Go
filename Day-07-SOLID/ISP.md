**인터페이스 분리 원칙 (ISP, Interface Segregation Principle)**

## 정의
인터페이스 분리 원칙(ISP)은 객체 지향 설계 원칙 중 하나로, 다음과 같이 정의됩니다.

> 클라이언트는 자신이 이용하지 않는 메서드에 의존하지 않아야 한다.

즉, 인터페이스는 작고 명확하게 분리되어야 하며, 불필요한 메서드가 포함되지 않도록 해야 합니다.

## 이점
- **불필요한 의존성을 제거**하여 유지보수성을 향상시킵니다.
- **인터페이스가 작아지므로** 클라이언트가 불필요한 메서드를 구현하지 않아도 됩니다.
- **코드의 가독성과 재사용성이 증가**합니다.

## ISP 위반 사례

### 잘못된 예제
```go
type Report interface {
    Report() string
    Pages() int
    Author() string
    WrittenDate() time.Time
}

func SendReport(r Report) {
    send(r.Report())
}
```
**문제점:**
- `SendReport` 함수는 `Report()` 메서드만 필요하지만, `Report` 인터페이스는 `Pages()`, `Author()`, `WrittenDate()`도 포함하고 있습니다.
- 이는 **불필요한 의존성**을 초래하며, ISP를 위반합니다.

## 해결 방법
### 올바른 인터페이스 분리
```go
type Report interface {
    Report() string
}

type WrittenInfo interface {
    Pages() int
    Author() string
    WrittenDate() time.Time
}

func SendReport(r Report) {
    send(r.Report())
}
```
**개선점:**
- `Report` 인터페이스는 `Report()` 메서드만 포함하도록 분리했습니다.
- `WrittenInfo` 인터페이스는 문서의 메타데이터 관련 메서드를 포함하도록 설계했습니다.
- 이제 `SendReport` 함수는 **필요한 기능만 의존**하게 되어 ISP를 준수합니다.

## 결론
ISP를 준수하면 인터페이스가 단순하고 명확해지며, 불필요한 의존성을 줄일 수 있습니다. 올바른 인터페이스 설계를 통해 유지보수성과 확장성을 높이는 것이 중요합니다.

