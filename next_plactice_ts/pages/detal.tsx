import Link from 'next/link'
import useSWR, { useSWRConfig } from 'swr'
import { useRouter } from 'next/router'
import { useState } from 'react'
import axios from 'axios'
import Head from 'next/head'
import styles from '../styles/Home.module.css'
import { apiURL } from '../unify/const'

export default function Detail() {
  const { mutate } = useSWRConfig();
  const router = useRouter();

  // クエリパラメーターを取得する
  let { id } = router.query;
  // クエリパラメーターを取得できなかった場合は"false"という文字列を格納する
  if(!id){
    id = "false";
  }
  // fetcher関数の第一引数にはuseSWRの第一引数が入る
  const fetcher = async (address: string) => {
    const res = await fetch(address);
    // もしステータスコードが 200-299 の範囲内では無い場合はエラーページに遷移する
    if (!res.ok) {
      router.push("/_error");
    }
  
    return res.json();
  }
  const { data, error } = useSWR(`${apiURL}/detail?id=${id}`, fetcher);

  const [title, setTitle] = useState('');
  const [content, setContent] = useState('');
  const sendUpdate = async () => {
    axios.post(`${apiURL}/edit?id=${id}&title=${title}&content=${content}`)
    .then(()=> {
      setTitle("");
      setContent("");
      // SWRがrefetchを行う
      mutate(`${apiURL}/detail?id=${id}`);
    })
    // Go側でエラーがあった場合
    .catch(()=> {
      router.push("/_error");
    });
  };
  const sendDelete = async () => {
    axios.delete(`${apiURL}/delete?id=${id}`)
    .then(()=> {
      router.push("/");
    })
    // Go側でエラーがあった場合
    .catch(()=> {
      router.push("/_error")
    });
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