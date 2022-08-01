import type { NextPage } from 'next'
import Head from 'next/head'
import Image from 'next/image'
import List from './components/list'
import styles from '../styles/Home.module.css'
import useSWR from 'swr'
import {useState} from 'react'

const fetcher = (...args: any) => fetch(...args).then(res => res.json())

const Home: NextPage = () => {
  {console.log("index.js")}
  const [ address, setAddress ] = useState('http://localhost:8080/')
  const { data, error } = useSWR(address, fetcher)
  let items 
  // {console.log(data)}
  {if(data){
    // console.log("------------")
    // {console.log(typeof data)}
    // console.log(data.length)
    items = data.map((value: any,key: any)=>(
      <li key={key} value={key}>
        {value.title}
        {value.content}
      </li>      
    ))
  }}
  // console.log("data----" + JSON.stringify(data))
  const numbers = [1, 2, 3, 4, 5];
    console.log("------------")
  console.log(typeof numbers)

  return (
    <div className={styles.container}>
      <Head>
        <title>Create Next App</title>
        <meta name="description" content="Generated by create next app" />
      </Head>
      <main className={styles.main}>
      {data ? 
      <div>
        {items}
        {/* {numbers.map((number) =>
          <li>{number}</li>
        )} */}
        <p>table 外</p>
          {/* { data.map((value: any, key: any) => {
              {console.log(value.title)}
              {console.log(key)}
              {value.title}
              {key}
          })
          } */}
        <p>table 外</p>
      {/* <p>{data[0].title}</p>
      <p>{data[0].content}</p> */}
      </div>

  :
      <p>"データがありません"</p>
  }
        {/* <List /> */}
      </main>
    </div>
  )
}

export default Home
