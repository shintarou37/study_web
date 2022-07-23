import React, { useState } from 'react'
import useSWR from 'swr'
const fetcher = (...args: any[]) => fetch(...args).then(res => res.json())

export default function List(props: any) {
  const { data } = useSWR('/test_data/data.json', fetcher)
  return (
    <div>
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