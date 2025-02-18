# **단일 책임 원칙(SRP: Single Responsibility Principle)**  

## **정의**  
> **모든 객체(클래스, 모듈, 함수)는 단 하나의 책임만 가져야 한다.**  

## **이점**  
- 유지보수가 용이해진다.  
- 코드의 재사용성이 증가한다.  
- 수정이 필요한 경우, 최소한의 변경만으로 대응할 수 있다.  

---

## **SRP 위반 사례**  

```go
type FinanceReport struct { // 회계 보고서
    report string
}

func (f *FinanceReport) SendReport(email string) { // 보고서 전송 기능
    ...
}
```

### **왜 SRP를 위반했을까?**  
- `FinanceReport`는 **두 가지 책임(역할)**을 가지고 있다.  
  1. 보고서를 관리하는 역할  
  2. 보고서를 전송하는 역할  

- 만약 **마케팅 보고서**를 추가하려면?  

```go
type MarketingReport struct { // 마케팅 보고서
    report string
}

func (m *MarketingReport) SendReport(email string) { // 보고서 전송 기능
    ...
}
```

- 새로운 보고서 유형이 생길 때마다 `SendReport` 메서드를 추가해야 한다.  
- **각 보고서 클래스가 중복된 전송 로직을 포함하게 되어 유지보수가 어렵다.**  

---

## **SRP를 준수하는 방법: 책임 분리**  
- **보고서 관리**와 **보고서 전송**을 각각의 클래스로 분리해야 한다.  

### **1. Report 인터페이스 도입**  
```go
type Report interface {  // Report() 메서드를 포함하는 인터페이스
    Report() string
}
```

### **2. 보고서별 구현**  
```go
type FinanceReport struct { // 회계 보고서
    report string
}

func (r *FinanceReport) Report() string { // Report 인터페이스 구현
    return r.report
}

type MarketingReport struct { // 마케팅 보고서
    report string
}

func (r *MarketingReport) Report() string { // Report 인터페이스 구현
    return r.report
}
```

### **3. ReportSender 추가 (전송 책임 분리)**  
```go
type ReportSender struct {} // 보고서 전송을 담당

func (s *ReportSender) SendReport(report Report) {
    // Report 인터페이스 객체를 받아 처리
    ...
}
```

---

## **정리**  
✅ `FinanceReport`, `MarketingReport` 등은 **오직 보고서 데이터를 관리**하는 역할만 수행한다.  
✅ `ReportSender`는 **보고서를 전송하는 역할만 담당**한다.  
✅ 새로운 보고서 타입이 추가되더라도 `ReportSender`는 수정할 필요가 없다.  

**➡ SRP를 준수하면 코드의 유지보수성과 확장성이 증가한다!**
