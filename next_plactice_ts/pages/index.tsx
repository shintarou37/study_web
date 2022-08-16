import type { NextPage } from 'next'
import Head from 'next/head'
import Image from 'next/image'
import styles from '../styles/Home.module.css'
import useSWR, { useSWRConfig } from 'swr'
import { useState, useEffect } from 'react'
import { title } from 'process'
import axios from 'axios'
import Link from 'next/link'
import { useRouter } from 'next/router'

const fetcher = (url: string) => fetch(url).then(res => res.json())

const Home: NextPage = () => {
  const { mutate } = useSWRConfig()
  const [ address, setAddress ] = useState('http://localhost:8080/')
  const { data, error } = useSWR(address, fetcher)
  const [title, setTitle] = useState('');
  const [content, setContent] = useState('');
  const router = useRouter()

  const sendRegister = async () => {
    axios.post(`http://localhost:8080/register?title=${title}&content=${content}`)
    .then(()=> {
      setTitle("")
      setContent("")
      // SWRがrefetchを行う
      mutate("http://localhost:8080/")
    })
    .catch(()=> {
      router.push("/_error")
    });
  };
  
  let datas;
  if(data){
    datas = data.map((value: any,key: any)=>{
      let detal_address = `/detal/?id=${value.ID}`;
      return <ul>
        <h1>{value.ID}番</h1>
        <p>タイトル</p>
        <li>{value.title}</li>
        <p>内容</p>
        <li>{value.content}</li>
        <Link href={detal_address}>
          <a>詳細</a>
        </Link>
      </ul>
    })
  }

  return (
    <div className={styles.container}>
      <Head>
        <title>Create Next App</title>
        <meta name="description" content="Generated by create next app" />
      </Head>
      <main className={styles.main}>
      <h1>投稿フォーム</h1>
        <label>タイトル:</label><br></br>
        <input type="text" name="title" value={title} onChange={(e) => setTitle(e.target.value)}/><br></br>
        <label>内容</label><br></br>
        <input type="text" name="content" value={content} onChange={(e) => setContent(e.target.value)}/><br></br>
        <button type="submit" onClick={sendRegister}>送信</button><br></br>
        {data ? 
          <div>
            {datas}
          </div>
        :
          <p></p>
        }
      </main>
    </div>
  )
}

export default Home