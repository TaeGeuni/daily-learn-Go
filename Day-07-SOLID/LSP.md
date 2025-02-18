**리스코프 치환 원칙 (LSP, Liskov Substitution Principle)**

## 정의
리스코프 치환 원칙(LSP)은 객체 지향 프로그래밍의 중요한 원칙 중 하나로, 다음과 같이 정의됩니다.

> q(x)를 타입 T의 객체 x에 대해 증명할 수 있는 속성이라 하자. 그렇다면 S가 T의 하위 타입이라면 q(y)는 타입 S의 객체 y에 대해 증명할 수 있어야 한다.

즉, 자식 클래스(S)는 언제나 부모 클래스(T)를 대체할 수 있어야 하며, 프로그램의 기능에 영향을 주지 않아야 합니다.

## 이점
- **예상치 못한 작동을 예방**할 수 있습니다.
- 코드의 **유지보수성을 향상**시킵니다.
- 객체 지향 프로그래밍에서 **다형성을 올바르게 활용**할 수 있습니다.

## LSP 위반 사례
### 예제: 사각형과 정사각형 문제
```java
class Rectangle {
    protected int width;
    protected int height;

    public void setWidth(int width) {
        this.width = width;
    }

    public void setHeight(int height) {
        this.height = height;
    }

    public int getArea() {
        return width * height;
    }
}

class Square extends Rectangle {
    @Override
    public void setWidth(int width) {
        this.width = width;
        this.height = width; // 정사각형이므로 높이도 동일하게 설정
    }

    @Override
    public void setHeight(int height) {
        this.height = height;
        this.width = height; // 정사각형이므로 너비도 동일하게 설정
    }
}
```
이 코드에서 `Square` 클래스는 `Rectangle`을 상속하지만, `setWidth`와 `setHeight`의 동작이 부모 클래스의 기대와 다르게 동작합니다. 이는 LSP를 위반하는 대표적인 사례입니다.

## 해결 방법
### 1. 상속 대신 조합(Composition) 사용
```java
class Rectangle {
    protected int width;
    protected int height;

    public void setWidth(int width) {
        this.width = width;
    }

    public void setHeight(int height) {
        this.height = height;
    }

    public int getArea() {
        return width * height;
    }
}

class Square {
    private int side;

    public void setSide(int side) {
        this.side = side;
    }

    public int getArea() {
        return side * side;
    }
}
```
이제 `Square` 클래스는 `Rectangle`을 상속하지 않으며, 각각 독립적인 개념으로 동작합니다. 따라서 LSP 위반이 발생하지 않습니다.

### 2. 인터페이스 활용
```java
interface Shape {
    int getArea();
}

class Rectangle implements Shape {
    private int width;
    private int height;
    
    public void setWidth(int width) {
        this.width = width;
    }
    
    public void setHeight(int height) {
        this.height = height;
    }
    
    public int getArea() {
        return width * height;
    }
}

class Square implements Shape {
    private int side;
    
    public void setSide(int side) {
        this.side = side;
    }
    
    public int getArea() {
        return side * side;
    }
}
```
이렇게 하면 `Rectangle`과 `Square`가 각각 `Shape` 인터페이스를 구현하여 독립적으로 동작하므로, LSP를 위반하지 않습니다.

## 결론
LSP는 객체 지향 설계에서 매우 중요한 원칙이며, 이를 준수하지 않으면 예상치 못한 버그가 발생할 수 있습니다. 올바른 설계를 위해 **상속보다 조합을 고려**하고, **인터페이스를 활용**하는 방법을 적용하면 LSP를 준수하는 더 나은 설계를 할 수 있습니다.

---
Go에서는 상속이 없지만, 잘못된 인터페이스 설계로 인해 LSP 위반이 발생할 수 있습니다.

이를 방지하려면:
- ✔ 인터페이스를 너무 포괄적으로 설계하지 말 것
- ✔ 불필요한 메서드 요구 사항을 줄이기 위해 인터페이스를 분리할 것

Go에서도 LSP를 신경 쓰지 않으면 예상치 못한 버그가 발생할 수 있으니, 올바른 인터페이스 설계를 해야 합니다! 🚀