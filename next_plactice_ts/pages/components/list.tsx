import React, { useState } from 'react'
import useSWR from 'swr'
const fetcher = (...args: any[]) => fetch(...args).then(res => res.json())

export default function List(props: any) {
  const { data } = useSWR('/test_data/data.json', fetcher)
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
        <table className="table table-dark">
          <tbody>
            {data != undefined ? data.data.map((value, key)=> (
              <tr key={key}>
                <th>{value.title}</th>
                <td>{value.content}</td>
              </tr>
            )) : <tr><th></th><td>no data.</td><td></td></tr>}
          </tbody>
        </table>

    </div>
  )
}