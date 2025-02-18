# **개방-폐쇄 원칙 (OCP: Open-Closed Principle)**  

## **정의**  
> **확장에는 열려 있고, 변경에는 닫혀 있어야 한다.**  

즉, 기존 코드를 수정하지 않고도 새로운 기능을 추가할 수 있어야 한다.  

## **이점**  
- 기존 코드를 수정하지 않고 기능을 확장할 수 있다.  
- **상호 결합도를 줄여** 유지보수성을 높이고, 변경에 따른 버그 발생 가능성을 줄인다.  

---

## **OCP 위반 사례**  

```go
func SendReport(r *Report, method SendType, receiver string) {
    switch method {
    case Email:
        // 이메일 전송
    case Fax:
        // 팩스 전송
    case PDF:
        // PDF 파일 생성
    case Printer:
        // 프린팅
    ...
    }
}
```

### **왜 문제인가?**  
- 새로운 전송 방식(e.g., **SMS**)을 추가하려면 `SendReport` 함수를 수정해야 한다.  
- 시간이 지날수록 `SendReport` 함수는 점점 **비대해지고 유지보수가 어려워진다.**  
- **OCP를 위반**하여 변경이 어렵고 확장성이 낮아진다.  

---

## **OCP를 준수하는 방법: 인터페이스 활용**  

### **1. ReportSender 인터페이스 정의**  
```go
type ReportSender interface {
    Send(r *Report)
}
```

### **2. 전송 방식을 개별 구조체로 분리**  
```go
type EmailSender struct{}

func (e *EmailSender) Send(r *Report) {
    // 이메일 전송 로직
}

type FaxSender struct{}

func (f *FaxSender) Send(r *Report) {
    // 팩스 전송 로직
}
```

### **3. 새로운 전송 방식 추가 (OCP 준수)**  
```go
type SMSSender struct{}

func (s *SMSSender) Send(r *Report) {
    // SMS 전송 로직
}
```

**➡ 기존 `SendReport` 함수를 수정하지 않고도 새로운 전송 방식을 쉽게 추가할 수 있다!**  

---

## **정리**  
✅ `SendReport` 같은 단일 함수에서 모든 전송 방식을 관리하면 **확장성이 떨어진다.**  
✅ `ReportSender` 인터페이스를 활용하여 **새로운 기능 추가 시 기존 코드 수정 없이 확장 가능**하게 만든다.  
✅ OCP를 준수하면 **유지보수가 쉬워지고, 코드 변경 없이 기능을 확장할 수 있다.**  

**➡ OCP를 적용하면 변경 없이 확장 가능하여 코드 품질이 향상된다!** 🚀
