# Go 채널 (Channels) 및 컨텍스트 (Context)

## 채널 (Channels)

### 1. 채널이란?
채널은 **고루틴 간 안전하게 데이터를 주고받을 수 있는 통신 파이프**(메시지 큐)입니다.
고루틴끼리 값을 전달할 때 사용되며, 동시성 문제를 피할 수 있도록 설계되었습니다.

---

### 2. 채널 생성
`make()` 함수를 사용하여 채널을 생성할 수 있습니다.
```go
messages := make(chan string) // 문자열 타입의 채널 생성
```
- **버퍼드 채널**: `make(chan T, capacity)`와 같이 용량을 지정하면 비동기적으로 메시지를 보낼 수 있습니다.
- **언버퍼드 채널**: 용량을 지정하지 않으면 데이터를 받을 때까지 블로킹됩니다.

---

### 3. 채널 데이터 송수신

#### 3.1 데이터 보내기
```go
messages <- "This is message"
```

#### 3.2 데이터 받기
```go
msg := <-messages
```

#### 3.3 데이터 반복 수신 (Range)
```go
for n := range messages {
    fmt.Println(n) // 채널에서 데이터 읽기
}
```
- `range`를 사용하면 채널이 닫힐 때까지 데이터를 읽을 수 있습니다.

---

### 4. 채널 닫기
채널을 더 이상 사용하지 않을 경우 `close()`를 호출합니다.
```go
close(messages)
```
- **닫는 주체는 데이터를 보내는 측**이어야 합니다.
- 닫힌 채널에 값을 보내면 **패닉이 발생**합니다.

---

### 5. 고루틴 릭 (Goroutine Leak)
채널이 닫히지 않으면, 채널에서 값을 기다리는 고루틴이 영원히 블로킹되어 **좀비 고루틴(Leak)**이 발생할 수 있습니다.

---

### 6. Select 문

#### 6.1 여러 채널을 동시에 기다릴 때 사용
```go
select {
case n1 := <-ch1:
    fmt.Println("Received from ch1:", n1)
case n2 := <-ch2:
    fmt.Println("Received from ch2:", n2)
default:
    fmt.Println("No data available")
}
```
- 여러 채널 중 **준비된 데이터가 있는 경우 하나를 랜덤하게 선택**합니다.
- `default` 케이스를 추가하면 **블로킹을 방지**할 수 있습니다.

---

### 7. 시간 관련 채널 (`time` 패키지)

#### 7.1 일정 간격 실행 (`time.Tick`)
```go
ticker := time.Tick(1 * time.Second)
```
- **일정 간격**(1초)마다 신호를 주는 채널을 반환합니다.
- 단, 중지할 수 없으므로 `time.NewTicker` 사용을 고려하세요.

#### 7.2 일정 시간 후 실행 (`time.After`)
```go
<-time.After(5 * time.Second)
```
- 지정된 시간(5초)이 지난 후 **한 번만 신호를 주는 채널**을 반환합니다.
- 주로 타임아웃 처리에 사용됩니다.

---

## 컨텍스트 (Context)

### 1. 컨텍스트란?
컨텍스트는 **작업의 범위, 취소, 데드라인(시간 제한) 및 요청 범위의 값을 전달**하는 구조체입니다.

### 2. 주요 기능
- **취소 신호 전파**: 작업 중단 또는 취소 가능
- **데드라인/타임아웃 지정**: 지정된 시간 이후 자동 취소
- **값 전달**: 요청 관련 정보를 저장 및 전파

---

### 3. 컨텍스트 생성 및 사용

#### 3.1 기본 컨텍스트
```go
ctx := context.Background() // 기본 컨텍스트
```

#### 3.2 취소 가능한 컨텍스트
```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel() // 작업 완료 후 반드시 호출하여 리소스 해제
```

#### 3.3 타임아웃 컨텍스트
```go
ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
defer cancel()
```

#### 3.4 컨텍스트 활용 예시
```go
func doSomething(ctx context.Context) {
    select {
    case <-ctx.Done():
        fmt.Println("작업 취소됨")
        return
    case result := <-someChannel:
        fmt.Println("작업 완료:", result)
    }
}
```
- `ctx.Done()`을 통해 컨텍스트 취소 여부를 감지할 수 있습니다.

---

## 추가 참고 사항

### 1. 버퍼드 vs 언버퍼드 채널
- **버퍼드 채널**: `make(chan int, 3)`와 같이 용량 지정 가능, 비동기 전송 가능
- **언버퍼드 채널**: 용량 미지정, 수신자가 준비될 때까지 블로킹됨

### 2. 컨텍스트 전파 패턴
- 부모 컨텍스트에서 생성된 컨텍스트는 **자식 컨텍스트로 전파**됩니다.
- 컨텍스트 내부에 너무 많은 데이터를 담지 않고 최소한의 정보(취소, 타임아웃, 로깅 등)만 유지하는 것이 좋습니다.

---
