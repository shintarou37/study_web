import React, { useState } from 'react'
// import useSWR from 'swr'
import axios from 'axios'
import { type } from 'os';
// const fetcher = (...args: any[]) => fetch(...args).then(res => res.json())

export default function List(props: any) {
  // 関数の外で定義するとエラーになる
// const [results, b] = useState({"title": "", "content": ""})
console.log("results")
// const clickFunc = (ret: any) => {

//   b(ret)
// }
//   // let results : any = {"title": "", "content": ""};

  const url = axios.get("http://localhost:8080/")
//   // thenで成功した場合の処理
//   .then((ret) => {
//     // console.log(ret.data)
//       // results = ret.data;
//       // clickFunc(ret.data)
//   })
//   // catchでエラー時の挙動を定義
//   .catch(err => {
//       console.log("err:", err);
//   });
  return (
    <div>
      <form>
      <h3>フォーム</h3><br></br>
        <label>タイトル</label><br></br>
          <input name="title"></input><br></br>
          <label>内容</label><br></br>
          <textarea name="content"></textarea><br></br>
          <input type="submit" value="送信"></input>
      </form><br></br><br></br>
      <h3>一覧</h3>
      <ul>
        {/* <li>{results.title}</li> */}
      </ul>
    </div>
  )
}