# import cycle error

> 아래 블로그의 내용을 참고함
> - https://dongyea.tistory.com/5

- Go에서는 패키지간 서로 참조 하게되면 순환 참조 에러가 발생한다.  
- 동일한 패키지에 다른 패키지에 속한  struct 의 method 들을 가지는 interface 만들어서 사용하므로서 이를 해결 하 수 있다.

# 테스트

- parent.go에서 `child_error` 를 사용하면 import cycle error 발생
