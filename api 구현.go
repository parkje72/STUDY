package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// 구조체 - 데이터
type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// 슬라이스 초기화 - 데이터 저장공간
var users = []User{}

func main() {
	//데이터 전송
	//최적경로 설정(라우터)
	rt := http.NewServeMux()
	//라우터의 핸들러함수로 최적경로 설정
	rt.HandleFunc("/user/", func(w http.ResponseWriter, r *http.Request) {

		//메소드에 따른 기능 실행
		switch r.Method {
		case "POST": //데이터 등록(Create)
			post(w, r)
		case "GET": //데이터 조회(Read)
			get(w, r)
		case "PUT": //데이터 수정(Update)
			put(w, r)
		case "DELETE": //데이터 삭제(Delete)
			delete(w, r)
		}

	})

	//서버 구동 - ListenAndServe()
	err := http.ListenAndServe(":8081", rt)
	if err != nil {
		fmt.Println(err)
	}
}

// 기능 구현 함수
// 새로운 유저 등록
func post(w http.ResponseWriter, r *http.Request) {
	//바디 파라미터에 등록한 문자열 데이터를 디코딩해서 슬라이스에 저장
	var user User

	err := json.NewDecoder(r.Body).Decode(&users) //바디 파라미터 -> User 구조체
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("에러 - 디코딩 실패"))
	}

	//users 슬라이스에 추가
	users = append(users, user)

	//데이터 보여주기
	data, _ := json.Marshal(user)
	w.Header().Add("Content-Type", "application/json") //콘텐츠 타입 명시
	fmt.Fprint(w, string(data))

}

// 해당 아이디 유저 조회
func get(w http.ResponseWriter, r *http.Request) {
	//URL 에서 사용자 id 가져오기
	path := strings.Split(r.URL.Path, "/")
	id := path[len(path)-1]

	//users 슬라이스 순회하며 id 비교해서 조회할 사용자 찾기
	var index = -1
	for i, user := range users {
		if user.ID == id {
			index = i
			break
		}
	}

	//index 존재하면 업데이트 그외에는 에러처리
	if index != -1 {
		data, err := json.Marshal(users[index])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("유저 마샬링 에러"))
			return
		}
		w.Header().Add("Content-Type", "application/json") //콘텐츠 타입 명시
		fmt.Fprint(w, string(data))

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("에러 - 디코딩 실패"))
		}

	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("해당되는 유저 없음"))
	}

	// json.NewEncoder(w).Encode(users[index]) //바디 파라미터 -> User 구조체

}

// 해당 아이디 유저 정보 수정
func put(w http.ResponseWriter, r *http.Request) {

	//URL 에서 사용자 id 가져오기
	path := strings.Split(r.URL.Path, "/")
	id := path[len(path)-1]

	//users 슬라이스 순회하며 id 비교해서 조회할 사용자 찾기
	var index = -1
	for i, user := range users {
		if user.ID == id {
			index = i
			break
		}
	}

	//바디 파라미터에 등록한 문자열 데이터를 디코딩
	var user User
	err := json.NewDecoder(r.Body).Decode(&user) //바디 파라미터 -> User 구조체
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("에러 - 디코딩 실패"))
	}

	//index 존재하면 업데이트 그외에는 에러처리
	if index != -1 {
		users[index] = user
		w.Write([]byte("업데이트 성공"))
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("해당되는 유저 없음"))
	}

	//업데이트 된 데이터 보여주기
	data, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("유저 마샬링 실패"))
	}
	w.Header().Add("Content-Type", "application/json") //콘텐츠 타입 명시
	fmt.Fprint(w, string(data))
}

// 해당 아이디 데이터 삭제
func delete(w http.ResponseWriter, r *http.Request) {
	//URL 에서 사용자 id 가져오기
	// var user User
	path := strings.Split(r.URL.Path, "/")
	id := path[len(path)-1]

	//users 슬라이스 순회하며 id 비교해서 조회할 사용자 찾기
	var index = -1
	for i, user := range users {
		if user.ID == id {
			index = i
			break
		}
	}

	//index 바뀌었다면 삭제 그외에는 에러처리
	if index != -1 {
		users = append(users[:index], users[index+1:]...)
		w.Write([]byte("삭제 성공"))
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("해당되는 유저 없음"))
	}

}
