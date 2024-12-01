# Huffy

### 사용 라이브러리

- air
- fiber
- sqlite3
- colly

### API 설계

- 학식 정보 받아오기
```
/api/menu/today?name={식당이름}
/api/menu/tomorrow?name={식당이름}
```

```
식당이름

<서울>
inmungwan 인문관식당
gyosuhaegwan 교수회관식당

<글로벌>
husaeng_professor 후생관 교직원식당
husaeng_student 후생관 학생식당
hufsdorm 기숙사식당
```

- 공지사항 받아오기
```
/api/notice
```