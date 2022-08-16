import Link from 'next/link'
import useSWR from 'swr'
import { useRouter } from 'next/router'
import { useState } from 'react'
import axios from 'axios'
import Head from 'next/head'
import styles from '../styles/Home.module.css'

const fetcher = (url: string) => fetch(url).then(res => res.json())

export default function Detail() {
  const router = useRouter()
  // クエリパラメーターを取得する
  let { id } = router.query;
  // クエリパラメーターを取得できなかった場合は"false"という文字列を格納する
  if(!id){
    id = "false"
  }
  let { data, error } = useSWR(`http://localhost:8080/detail?id=${id}`, fetcher)
  const [title, setTitle] = useState('');
  const [content, setContent] = useState('');
  const sendUpdate = async () => {
    axios.post(`http://localhost:8080/edit?id=${id}&title=${title}&content=${content}`)
    .then((response)=> {
      setTitle("")
      setContent("")
      // 変数を全て変更すると反映されないため必要なキーを指定して変更している
      data.title = response.data.title
      data.content = response.data.content
      data.UpdatedAt = response.data.UpdatedAt
    })
  };
  const sendDelete = async () => {
    axios.delete(`http://localhost:8080/delete?id=${id}`)
    .then(()=> {
      router.push("/")
    })
  };
  // dataがない場合に戻り値を渡すと一瞬レイアウトが崩れる
  if(data){
    return (
      <div className={styles.container}>
        <Head>
          <title>Create Next App</title>
          <meta name="description" content="Generated by create next app" />
        </Head>
        <h1>詳細画面</h1>
        {data ? 
            <div>
              <ul>
                ID
                <li>{data.ID}</li>
                タイトル
                <li>{data.title}</li>
                内容
                <li>{data.content}</li>
                更新日
                <li>{data.UpdatedAt}</li>
                作成日
                <li>{data.CreatedAt}</li>
              </ul>
            </div>
          :
            <p></p>
          }
        <h1>編集フォーム</h1>
        <label>タイトル:</label><br></br>
        <input type="text" name="title" value={title} onChange={(e) => setTitle(e.target.value)}/><br></br>
        <label>内容</label><br></br>
        <input type="text" name="content" value={content} onChange={(e) => setContent(e.target.value)}/><br></br>
        <button type="submit" onClick={sendUpdate}>更新</button><br></br>
        <button type="submit" onClick={sendDelete}>削除</button><br></br>
        <Link href="/">
            <a>トップ画面へ戻る</a>
        </Link>
      </div>
    )
  }
}